/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/coreos/go-semver/semver"
	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/client/webclient"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/auth"
	"github.com/gravitational/teleport/lib/config"
	"github.com/gravitational/teleport/lib/fixtures"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/tbot/botfs"
	"github.com/gravitational/teleport/lib/tbot/identity"
	"github.com/gravitational/teleport/lib/utils/golden"
	"github.com/gravitational/trace"
	"github.com/stretchr/testify/require"
)

type templateSSHClientAuthMock struct {
	auth.ClientI
	clusterName string
	t           *testing.T
}

func (m *templateSSHClientAuthMock) GetClusterName(opts ...services.MarshalOption) (types.ClusterName, error) {
	cn, err := types.NewClusterName(types.ClusterNameSpecV2{
		ClusterName: m.clusterName,
		ClusterID:   "aa-bb-cc",
	})
	require.NoError(m.t, err)
	return cn, nil
}

func (m *templateSSHClientAuthMock) Ping(ctx context.Context) (proto.PingResponse, error) {
	require.NotNil(m.t, ctx)
	return proto.PingResponse{
		ProxyPublicAddr: "tele.blackmesa.gov:443",
	}, nil
}

func (m *templateSSHClientAuthMock) GetCertAuthority(ctx context.Context, id types.CertAuthID, loadKeys bool, opts ...services.MarshalOption) (types.CertAuthority, error) {
	require.NotNil(m.t, ctx)
	require.Equal(m.t, types.CertAuthID{
		Type:       types.HostCA,
		DomainName: m.clusterName,
	}, id)
	require.False(m.t, loadKeys)

	ca, err := types.NewCertAuthority(types.CertAuthoritySpecV2{
		Type:        types.HostCA,
		ClusterName: m.clusterName,
		ActiveKeys: types.CAKeySet{
			SSH: []*types.SSHKeyPair{
				// Two of these to ensure that both are written to known hosts
				{
					PrivateKey: []byte(fixtures.SSHCAPrivateKey),
					PublicKey:  []byte(fixtures.SSHCAPublicKey),
				},
				{
					PrivateKey: []byte(fixtures.SSHCAPrivateKey),
					PublicKey:  []byte(fixtures.SSHCAPublicKey),
				},
			},
		},
	})
	require.NoError(m.t, err)
	return ca, nil
}

type templateSSHClientMockBot struct {
	mockAuth *templateSSHClientAuthMock
}

func (t *templateSSHClientMockBot) Client() auth.ClientI {
	return t.mockAuth
}

func (t *templateSSHClientMockBot) GetCertAuthorities(ctx context.Context, caType types.CertAuthType) ([]types.CertAuthority, error) {
	return t.mockAuth.GetCertAuthorities(ctx, caType, false)
}

func (t *templateSSHClientMockBot) AuthPing(ctx context.Context) (*proto.PingResponse, error) {
	ping, err := t.mockAuth.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &ping, err
}

func (t *templateSSHClientMockBot) ProxyPing(_ context.Context) (*webclient.PingResponse, error) {
	return nil, trace.NotImplemented("not implemented")
}

func TestTemplateSSHClient_Render(t *testing.T) {
	tests := []struct {
		name       string
		sshVersion *semver.Version
		goldenName string
	}{
		{
			name:       "all enabled",
			sshVersion: semver.New("8.5.0"),
			goldenName: "ssh_config",
		},
		{
			name:       "no accepted algorithms",
			sshVersion: semver.New("8.0.0"),
			goldenName: "ssh_config_no_accepted_algos",
		},
		{
			name:       "no accepted and host keys algorithms",
			sshVersion: semver.New("7.3.0"),
			goldenName: "ssh_config_no_host_keys_algos",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dir := t.TempDir()
			mockAuth := &templateSSHClientAuthMock{
				t:           t,
				clusterName: "black-mesa",
			}
			mockBot := &templateSSHClientMockBot{
				mockAuth: mockAuth,
			}
			getSSHVersion := func() (*semver.Version, error) {
				return tt.sshVersion, nil
			}
			getExecutablePath := func() (string, error) {
				return "/path/to/tbot", nil
			}

			template := TemplateSSHClient{
				ProxyPort: 1337,
				generator: config.NewCustomSSHConfigGenerator(getSSHVersion, getExecutablePath),
			}
			// ident is passed in, but not used.
			var ident *identity.Identity
			dest := &DestinationConfig{
				DestinationMixin: DestinationMixin{
					Directory: &DestinationDirectory{
						Path:     dir,
						Symlinks: botfs.SymlinksInsecure,
						ACLs:     botfs.ACLOff,
					},
				},
			}

			err := template.Render(context.Background(), mockBot, ident, dest)
			require.NoError(t, err)

			replaceTestDir := func(b []byte) []byte {
				return bytes.ReplaceAll(b, []byte(dir), []byte("/test/dir"))
			}

			knownHostBytes, err := os.ReadFile(filepath.Join(dir, knownHostsName))
			require.NoError(t, err)
			knownHostBytes = replaceTestDir(knownHostBytes)
			sshConfigBytes, err := os.ReadFile(filepath.Join(dir, sshConfigName))
			require.NoError(t, err)
			sshConfigBytes = replaceTestDir(sshConfigBytes)
			if golden.ShouldSet() {
				golden.SetNamed(t, "known_hosts", knownHostBytes)
				golden.SetNamed(t, "ssh_config", sshConfigBytes)
			}
			require.Equal(
				t, string(golden.GetNamed(t, "known_hosts")), string(knownHostBytes),
			)
			require.Equal(
				t, string(golden.GetNamed(t, tt.goldenName)), string(sshConfigBytes),
			)
		})
	}
}
