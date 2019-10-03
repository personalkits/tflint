// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsEc2ClientVpnEndpointInvalidTransportProtocolRule checks the pattern is valid
type AwsEc2ClientVpnEndpointInvalidTransportProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2ClientVpnEndpointInvalidTransportProtocolRule returns new rule with default attributes
func NewAwsEc2ClientVpnEndpointInvalidTransportProtocolRule() *AwsEc2ClientVpnEndpointInvalidTransportProtocolRule {
	return &AwsEc2ClientVpnEndpointInvalidTransportProtocolRule{
		resourceType:  "aws_ec2_client_vpn_endpoint",
		attributeName: "transport_protocol",
		enum: []string{
			"tcp",
			"udp",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2ClientVpnEndpointInvalidTransportProtocolRule) Name() string {
	return "aws_ec2_client_vpn_endpoint_invalid_transport_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2ClientVpnEndpointInvalidTransportProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2ClientVpnEndpointInvalidTransportProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2ClientVpnEndpointInvalidTransportProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2ClientVpnEndpointInvalidTransportProtocolRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`transport_protocol is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
