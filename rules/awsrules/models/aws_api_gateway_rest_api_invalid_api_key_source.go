// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsAPIGatewayRestAPIInvalidAPIKeySourceRule checks the pattern is valid
type AwsAPIGatewayRestAPIInvalidAPIKeySourceRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayRestAPIInvalidAPIKeySourceRule returns new rule with default attributes
func NewAwsAPIGatewayRestAPIInvalidAPIKeySourceRule() *AwsAPIGatewayRestAPIInvalidAPIKeySourceRule {
	return &AwsAPIGatewayRestAPIInvalidAPIKeySourceRule{
		resourceType:  "aws_api_gateway_rest_api",
		attributeName: "api_key_source",
		enum: []string{
			"HEADER",
			"AUTHORIZER",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayRestAPIInvalidAPIKeySourceRule) Name() string {
	return "aws_api_gateway_rest_api_invalid_api_key_source"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayRestAPIInvalidAPIKeySourceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayRestAPIInvalidAPIKeySourceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayRestAPIInvalidAPIKeySourceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayRestAPIInvalidAPIKeySourceRule) Check(runner *tflint.Runner) error {
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
					`api_key_source is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
