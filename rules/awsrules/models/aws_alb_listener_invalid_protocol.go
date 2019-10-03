// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsALBListenerInvalidProtocolRule checks the pattern is valid
type AwsALBListenerInvalidProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsALBListenerInvalidProtocolRule returns new rule with default attributes
func NewAwsALBListenerInvalidProtocolRule() *AwsALBListenerInvalidProtocolRule {
	return &AwsALBListenerInvalidProtocolRule{
		resourceType:  "aws_alb_listener",
		attributeName: "protocol",
		enum: []string{
			"HTTP",
			"HTTPS",
			"TCP",
			"TLS",
			"UDP",
			"TCP_UDP",
		},
	}
}

// Name returns the rule name
func (r *AwsALBListenerInvalidProtocolRule) Name() string {
	return "aws_alb_listener_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsALBListenerInvalidProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsALBListenerInvalidProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsALBListenerInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsALBListenerInvalidProtocolRule) Check(runner *tflint.Runner) error {
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
					`protocol is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
