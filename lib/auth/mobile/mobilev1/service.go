// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mobilev1

import (
	"context"
	mobilev1pb "github.com/gravitational/teleport/api/gen/proto/go/teleport/mobile/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/auth/keystore"
	"github.com/gravitational/teleport/lib/authz"
	"github.com/gravitational/teleport/lib/jwt"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"time"
)

type jwtSigner interface {
	// GetDomainName returns local auth domain of the current auth server
	GetDomainName() (string, error)

	// GetCertAuthority returns certificate authority by given id. Parameter loadSigningKeys
	// controls if signing keys are loaded
	GetCertAuthority(ctx context.Context, id types.CertAuthID, loadSigningKeys bool) (types.CertAuthority, error)

	// GetKeyStore returns the KeyStore used by the auth server
	GetKeyStore() *keystore.Manager
}

// ServiceConfig holds configuration options for
// the mobile gRPC service.
type ServiceConfig struct {
	Authorizer authz.Authorizer
	Logger     logrus.FieldLogger
	JWTSigner  jwtSigner
	Clock      clockwork.Clock
}

// Service implements the teleport.mobile.v1.MobileService RPC service.
type Service struct {
	mobilev1pb.UnimplementedMobileServiceServer

	authorizer authz.Authorizer
	logger     logrus.FieldLogger
	clock      clockwork.Clock
	jwtSigner  jwtSigner
}

// NewService returns a new users gRPC service.
func NewService(cfg ServiceConfig) (*Service, error) {
	switch {
	case cfg.Authorizer == nil:
		return nil, trace.BadParameter("authorizer is required")
	case cfg.JWTSigner == nil:
		return nil, trace.BadParameter("jwt signer is required")
	}

	if cfg.Logger == nil {
		cfg.Logger = logrus.WithField(trace.Component, "mobile.service")
	}
	if cfg.Clock == nil {
		cfg.Clock = clockwork.NewRealClock()
	}

	return &Service{
		logger:     cfg.Logger,
		authorizer: cfg.Authorizer,
		clock:      cfg.Clock,
		jwtSigner:  cfg.JWTSigner,
	}, nil
}

func (s *Service) CreateAuthToken(ctx context.Context, req *mobilev1pb.CreateAuthTokenRequest) (*mobilev1pb.CreateAuthTokenResponse, error) {
	authCtx, err := s.authorizer.Authorize(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if req.Username == "" {
		req.Username = authCtx.User.GetName()
	}

	isUser := authz.IsLocalUser(*authCtx) && req.Username == authCtx.User.GetName()
	isAdmin := authz.HasBuiltinRole(*authCtx, string(types.RoleAdmin))
	if !isUser && !isAdmin {
		return nil, trace.AccessDenied("not user or admin requesting")
	}

	clusterName, err := s.jwtSigner.GetDomainName()
	if err != nil {
		return nil, trace.Wrap(err, "getting cluster name")
	}
	ca, err := s.jwtSigner.GetCertAuthority(ctx, types.CertAuthID{
		// TODO(noah): Is this the right CA to use?
		Type:       types.JWTSigner,
		DomainName: clusterName,
	}, true)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	signingKey, err := s.jwtSigner.GetKeyStore().GetJWTSigner(ctx, ca)
	if err != nil {
		return nil, trace.Wrap(err, "getting signing key")
	}
	signer, err := services.GetJWTSigner(signingKey, clusterName, s.clock)
	if err != nil {
		return nil, trace.Wrap(err, "creating signer")
	}
	token, err := signer.Sign(jwt.SignParams{
		Username: req.Username,
		Audience: clusterName,
		URI:      clusterName,
		Expires:  s.clock.Now().Add(time.Minute * 5),
	})
	if err != nil {
		return nil, trace.Wrap(err, "signing token")
	}

	return &mobilev1pb.CreateAuthTokenResponse{Token: token}, nil
}
