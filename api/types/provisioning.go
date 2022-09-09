/*
Copyright 2020-2022 Gravitational, Inc.

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

package types

import (
	"fmt"
	"time"

	"github.com/gravitational/teleport/api/defaults"

	"github.com/gravitational/trace"
)

// JoinMethod is the method used for new nodes to join the cluster.
type JoinMethod string

const (
	JoinMethodUnspecified JoinMethod = ""
	// JoinMethodToken is the default join method, nodes join the cluster by
	// presenting a secret token.
	JoinMethodToken JoinMethod = "token"
	// JoinMethodEC2 indicates that the node will join with the EC2 join method.
	JoinMethodEC2 JoinMethod = "ec2"
	// JoinMethodIAM indicates that the node will join with the IAM join method.
	JoinMethodIAM JoinMethod = "iam"
	// JoinMethodOIDCGCP indicates that the node will join with the GCP OIDC
	// join method.
	JoinMethodOIDCGCP JoinMethod = "oidc-gcp"
)

// ProvisionToken is a provisioning token
type ProvisionToken interface {
	Resource
	// SetMetadata sets resource metatada
	SetMetadata(meta Metadata)
	// GetRoles returns a list of teleport roles
	// that will be granted to the user of the token
	// in the crendentials
	GetRoles() SystemRoles
	// SetRoles sets teleport roles
	SetRoles(SystemRoles)
	// GetBotName returns the BotName field which must be set for joining bots.
	GetBotName() string
	// GetSuggestedLabels returns the set of labels that the resource should add when adding itself to the cluster
	GetSuggestedLabels() Labels

	// String returns user friendly representation of the resource
	String() string
}

// NewProvisionToken returns a new provision token with the given roles.
func NewProvisionToken(token string, roles SystemRoles, expires time.Time) (ProvisionToken, error) {
	return NewProvisionTokenFromSpec(token, expires, ProvisionTokenSpecV3{
		Roles: roles,
	})
}

// NewProvisionTokenFromSpec returns a new provision token with the given spec.
func NewProvisionTokenFromSpec(token string, expires time.Time, spec ProvisionTokenSpecV3) (ProvisionToken, error) {
	t := &ProvisionTokenV3{
		Metadata: Metadata{
			Name:    token,
			Expires: &expires,
		},
		Spec: spec,
	}
	if err := t.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}
	return t, nil
}

// MustCreateProvisionToken returns a new valid provision token
// or panics, used in tests
func MustCreateProvisionToken(token string, roles SystemRoles, expires time.Time) ProvisionToken {
	t, err := NewProvisionToken(token, roles, expires)
	if err != nil {
		panic(err)
	}
	return t
}

// setStaticFields sets static resource header and metadata fields.
func (p *ProvisionTokenV2) setStaticFields() {
	p.Kind = KindToken
	p.Version = V2
}

// CheckAndSetDefaults checks and set default values for any missing fields.
func (p *ProvisionTokenV2) CheckAndSetDefaults() error {
	p.setStaticFields()
	if err := p.Metadata.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err)
	}

	if len(p.Spec.Roles) == 0 {
		return trace.BadParameter("provisioning token is missing roles")
	}
	if err := SystemRoles(p.Spec.Roles).Check(); err != nil {
		return trace.Wrap(err)
	}

	if SystemRoles(p.Spec.Roles).Include(RoleBot) && p.Spec.BotName == "" {
		return trace.BadParameter("token with role %q must set bot_name", RoleBot)
	}

	if p.Spec.BotName != "" && !SystemRoles(p.Spec.Roles).Include(RoleBot) {
		return trace.BadParameter("can only set bot_name on token with role %q", RoleBot)
	}

	hasAllowRules := len(p.Spec.Allow) > 0
	if p.Spec.JoinMethod == JoinMethodUnspecified {
		// Default to the ec2 join method if any allow rules were specified,
		// else default to the token method. These defaults are necessary for
		// backwards compatibility.
		if hasAllowRules {
			p.Spec.JoinMethod = JoinMethodEC2
		} else {
			p.Spec.JoinMethod = JoinMethodToken
		}
	}
	switch p.Spec.JoinMethod {
	case JoinMethodToken:
		if hasAllowRules {
			return trace.BadParameter("allow rules are not compatible with the %q join method", JoinMethodToken)
		}
	case JoinMethodEC2:
		if !hasAllowRules {
			return trace.BadParameter("the %q join method requires defined token allow rules", JoinMethodEC2)
		}
		for _, allowRule := range p.Spec.Allow {
			if allowRule.AWSARN != "" {
				return trace.BadParameter(`the %q join method does not support the "aws_arn" parameter`, JoinMethodEC2)
			}
			if allowRule.AWSAccount == "" && allowRule.AWSRole == "" {
				return trace.BadParameter(`allow rule for %q join method must set "aws_account" or "aws_role"`, JoinMethodEC2)
			}
		}
		if p.Spec.AWSIIDTTL == 0 {
			// default to 5 minute ttl if unspecified
			p.Spec.AWSIIDTTL = Duration(5 * time.Minute)
		}
	case JoinMethodIAM:
		if !hasAllowRules {
			return trace.BadParameter("the %q join method requires defined token allow rules", JoinMethodIAM)
		}
		for _, allowRule := range p.Spec.Allow {
			if allowRule.AWSRole != "" {
				return trace.BadParameter(`the %q join method does not support the "aws_role" parameter`, JoinMethodIAM)
			}
			if len(allowRule.AWSRegions) != 0 {
				return trace.BadParameter(`the %q join method does not support the "aws_regions" parameter`, JoinMethodIAM)
			}
			if allowRule.AWSAccount == "" && allowRule.AWSARN == "" {
				return trace.BadParameter(`allow rule for %q join method must set "aws_account" or "aws_arn"`, JoinMethodEC2)
			}
		}
	case JoinMethodOIDCGCP:
		return trace.BadParameter("%q join method is compatible with token resource version 3 and higher")
	default:
		return trace.BadParameter("unknown join method %q", p.Spec.JoinMethod)
	}

	return nil
}

// GetVersion returns resource version
func (p *ProvisionTokenV2) GetVersion() string {
	return p.Version
}

// GetRoles returns a list of teleport roles
// that will be granted to the user of the token
// in the crendentials
func (p *ProvisionTokenV2) GetRoles() SystemRoles {
	return p.Spec.Roles
}

// SetRoles sets teleport roles
func (p *ProvisionTokenV2) SetRoles(r SystemRoles) {
	p.Spec.Roles = r
}

// GetAllowRules returns the list of allow rules
func (p *ProvisionTokenV2) GetAllowRules() []*TokenRule {
	return p.Spec.Allow
}

// GetAWSIIDTTL returns the TTL of EC2 IIDs
func (p *ProvisionTokenV2) GetAWSIIDTTL() Duration {
	return p.Spec.AWSIIDTTL
}

// GetJoinMethod returns joining method that must be used with this token.
func (p *ProvisionTokenV2) GetJoinMethod() JoinMethod {
	return p.Spec.JoinMethod
}

// GetBotName returns the BotName field which must be set for joining bots.
func (p *ProvisionTokenV2) GetBotName() string {
	return p.Spec.BotName
}

// GetKind returns resource kind
func (p *ProvisionTokenV2) GetKind() string {
	return p.Kind
}

// GetSubKind returns resource sub kind
func (p *ProvisionTokenV2) GetSubKind() string {
	return p.SubKind
}

// SetSubKind sets resource subkind
func (p *ProvisionTokenV2) SetSubKind(s string) {
	p.SubKind = s
}

// GetResourceID returns resource ID
func (p *ProvisionTokenV2) GetResourceID() int64 {
	return p.Metadata.ID
}

// SetResourceID sets resource ID
func (p *ProvisionTokenV2) SetResourceID(id int64) {
	p.Metadata.ID = id
}

// GetMetadata returns metadata
func (p *ProvisionTokenV2) GetMetadata() Metadata {
	return p.Metadata
}

// SetMetadata sets resource metatada
func (p *ProvisionTokenV2) SetMetadata(meta Metadata) {
	p.Metadata = meta
}

// GetSuggestedLabels returns the labels the resource should set when using this token
func (p *ProvisionTokenV2) GetSuggestedLabels() Labels {
	return p.Spec.SuggestedLabels
}

// SetExpiry sets expiry time for the object
func (p *ProvisionTokenV2) SetExpiry(expires time.Time) {
	p.Metadata.SetExpiry(expires)
}

// Expiry returns object expiry setting
func (p *ProvisionTokenV2) Expiry() time.Time {
	return p.Metadata.Expiry()
}

// GetName returns server name
func (p *ProvisionTokenV2) GetName() string {
	return p.Metadata.Name
}

// SetName sets the name of the ProvisionTokenV2.
func (p *ProvisionTokenV2) SetName(e string) {
	p.Metadata.Name = e
}

// String returns the human readable representation of a provisioning token.
func (p ProvisionTokenV2) String() string {
	expires := "never"
	if !p.Expiry().IsZero() {
		expires = p.Expiry().String()
	}
	return fmt.Sprintf("ProvisionToken(Roles=%v, Expires=%v)", p.Spec.Roles, expires)
}

// ProvisionTokensFromV1 converts V1 provision tokens to resource list.
// This is used to convert ProvisionTokenV1s within the StaticTokens resource
// to valid ProvisionTokens to be used for the purpose of validating a
// joining client.
func ProvisionTokensFromV1(in []ProvisionTokenV1) []ProvisionToken {
	if in == nil {
		return nil
	}
	out := make([]ProvisionToken, len(in))
	for i := range in {
		out[i] = in[i].V3()
	}
	return out
}

// V3 returns V3 version of the v1 token resource
// This allows the embedded V1 tokens within a StaticToken to be used as a ProvisionToken
func (p *ProvisionTokenV1) V3() *ProvisionTokenV3 {
	t := &ProvisionTokenV3{
		Metadata: Metadata{
			Name:      p.Token,
			Namespace: defaults.Namespace,
		},
		Spec: ProvisionTokenSpecV3{
			Roles:      p.Roles,
			JoinMethod: JoinMethodToken,
		},
	}
	if !p.Expires.IsZero() {
		t.SetExpiry(p.Expires)
	}
	t.CheckAndSetDefaults()
	return t
}

// String returns the human readable representation of a provisioning token.
func (p ProvisionTokenV1) String() string {
	expires := "never"
	if p.Expires.Unix() != 0 {
		expires = p.Expires.String()
	}
	return fmt.Sprintf("ProvisionToken(Roles=%v, Expires=%v)",
		p.Roles, expires)
}

// ProvisionTokenV3 methods

// setStaticFields sets static resource header and metadata fields.
func (p *ProvisionTokenV3) setStaticFields() {
	p.Kind = KindToken
	p.Version = V3
}

func (p *ProvisionTokenV3) CheckAndSetDefaults() error {
	p.setStaticFields()
	if err := p.Metadata.CheckAndSetDefaults(); err != nil {
		return trace.Wrap(err)
	}

	if len(p.Spec.Roles) == 0 {
		return trace.BadParameter("provisioning token is missing roles")
	}
	if err := SystemRoles(p.Spec.Roles).Check(); err != nil {
		return trace.Wrap(err)
	}

	if p.Spec.BotName == "" && SystemRoles(p.Spec.Roles).Include(RoleBot) {
		return trace.BadParameter("token with role %q must set bot_name", RoleBot)
	} else if p.Spec.BotName != "" && !SystemRoles(p.Spec.Roles).Include(RoleBot) {
		return trace.BadParameter("can only set bot_name on token with role %q", RoleBot)
	}

	switch p.Spec.JoinMethod {
	case JoinMethodIAM:
		providerCfg := p.Spec.GetIAM()
		if providerCfg == nil {
			return trace.BadParameter(
				`"aws_iam" configuration must be provided for join method %q`,
				JoinMethodIAM,
			)
		}
		if err := providerCfg.CheckAndSetDefaults(); err != nil {
			return trace.Wrap(err)
		}
	case JoinMethodEC2:
		providerCfg := p.Spec.GetEC2()
		if providerCfg == nil {
			return trace.BadParameter(
				`"aws_ec2" configuration must be provided for join method %q`,
				JoinMethodIAM,
			)
		}
		if err := providerCfg.CheckAndSetDefaults(); err != nil {
			return trace.Wrap(err)
		}
	case JoinMethodOIDCGCP:
		providerCfg := p.Spec.GetOIDCGCP()
		if providerCfg == nil {
			return trace.BadParameter(
				`"oidc_gcp" configuration must be provided for join method %q`,
				JoinMethodIAM,
			)
		}
		if err := providerCfg.CheckAndSetDefaults(); err != nil {
			return trace.Wrap(err)
		}
	case JoinMethodToken:
	default:
		return trace.BadParameter(`"join_method" must be specified`)

	}
	return fmt.Errorf("unimplemented")
}

// GetRoles returns a list of teleport roles
// that will be granted to the user of the token
// in the crendentials
func (p *ProvisionTokenV3) GetRoles() SystemRoles {
	return p.Spec.Roles
}

// SetRoles sets teleport roles
func (p *ProvisionTokenV3) SetRoles(r SystemRoles) {
	p.Spec.Roles = r
}

// SetExpiry sets expiry time for the object
func (p *ProvisionTokenV3) SetExpiry(expires time.Time) {
	p.Metadata.SetExpiry(expires)
}

// Expiry returns object expiry setting
func (p *ProvisionTokenV3) Expiry() time.Time {
	return p.Metadata.Expiry()
}

// GetName returns server name
func (p *ProvisionTokenV3) GetName() string {
	return p.Metadata.Name
}

// SetName sets the name of the ProvisionTokenV3
func (p *ProvisionTokenV3) SetName(e string) {
	p.Metadata.Name = e
}

// GetBotName returns the BotName field which must be set for joining bots.
func (p *ProvisionTokenV3) GetBotName() string {
	return p.Spec.BotName
}

// GetKind returns resource kind
func (p *ProvisionTokenV3) GetKind() string {
	return p.Kind
}

// GetSubKind returns resource sub kind
func (p *ProvisionTokenV3) GetSubKind() string {
	return p.SubKind
}

// SetSubKind sets resource subkind
func (p *ProvisionTokenV3) SetSubKind(s string) {
	p.SubKind = s
}

// GetResourceID returns resource ID
func (p *ProvisionTokenV3) GetResourceID() int64 {
	return p.Metadata.ID
}

// SetResourceID sets resource ID
func (p *ProvisionTokenV3) SetResourceID(id int64) {
	p.Metadata.ID = id
}

// GetVersion returns resource version
func (p *ProvisionTokenV3) GetVersion() string {
	return p.Version
}

// GetMetadata returns metadata
func (p *ProvisionTokenV3) GetMetadata() Metadata {
	return p.Metadata
}

// SetMetadata sets resource metadata
func (p *ProvisionTokenV3) SetMetadata(meta Metadata) {
	p.Metadata = meta
}

// GetJoinMethod returns joining method that must be used with this token.
func (p *ProvisionTokenV3) GetJoinMethod() JoinMethod {
	return p.Spec.JoinMethod
}

// GetSuggestedLabels returns the labels the resource should set when using this token
func (p *ProvisionTokenV3) GetSuggestedLabels() Labels {
	return p.Spec.SuggestedLabels
}

// String returns the human readable representation of a provisioning token.
func (p ProvisionTokenV3) String() string {
	expires := "never"
	if !p.Expiry().IsZero() {
		expires = p.Expiry().String()
	}
	return fmt.Sprintf("ProvisionToken(Roles=%v, Expires=%v)", p.Spec.Roles, expires)
}

// Validation for provider specific config

func (a *ProvisionTokenSpecV3AWSEC2) CheckAndSetDefaults() error {
	if len(a.Allow) == 0 {
		return trace.BadParameter("the %q join method requires defined token allow rules", JoinMethodEC2)
	}
	for _, allowRule := range a.Allow {
		if allowRule.AWSAccount == "" && allowRule.AWSRole == "" {
			return trace.BadParameter(
				`allow rule for %q join method must set "aws_account" or "aws_role"`,
				JoinMethodEC2,
			)
		}
	}
	if a.AWSIIDTTL == 0 {
		// default to 5 minute ttl if unspecified
		a.AWSIIDTTL = Duration(5 * time.Minute)
	}
	return nil
}

func (a *ProvisionTokenSpecV3AWSIAM) CheckAndSetDefaults() error {
	if len(a.Allow) == 0 {
		return trace.BadParameter("the %q join method requires defined token allow rules", JoinMethodIAM)
	}
	for _, allowRule := range a.Allow {
		if allowRule.AWSAccount == "" && allowRule.AWSARN == "" {
			return trace.BadParameter(
				`allow rule for %q join method must set "aws_account" or "aws_arn"`,
				JoinMethodEC2,
			)
		}
	}
	return nil
}

func (a *ProvisionTokenSpecV3OIDCGoogle) CheckAndSetDefaults() error {
	if len(a.Allow) == 0 {
		return trace.BadParameter("the %q join method requires defined token allow rules", JoinMethodOIDCGCP)
	}
	for _, allowRule := range a.Allow {
		projectIDPresent := allowRule.ProjectID != ""
		projectNumberPresent := allowRule.ProjectNumber != 0
		subjectPresent := allowRule.Sub != ""

		// Ensure at least one of the three unique fields is present to avoid users creating insecure configurations
		if !projectIDPresent && !projectNumberPresent && !subjectPresent {
			return trace.BadParameter(
				`allow rule for %q join method must set "project_id", "project_number" or "sub""`,
				JoinMethodOIDCGCP,
			)
		}

	}
	return nil
}
