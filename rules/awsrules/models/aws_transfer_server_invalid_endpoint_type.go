// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsTransferServerInvalidEndpointTypeRule checks the pattern is valid
type AwsTransferServerInvalidEndpointTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsTransferServerInvalidEndpointTypeRule returns new rule with default attributes
func NewAwsTransferServerInvalidEndpointTypeRule() *AwsTransferServerInvalidEndpointTypeRule {
	return &AwsTransferServerInvalidEndpointTypeRule{
		resourceType:  "aws_transfer_server",
		attributeName: "endpoint_type",
		enum: []string{
			"PUBLIC",
			"VPC_ENDPOINT",
		},
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidEndpointTypeRule) Name() string {
	return "aws_transfer_server_invalid_endpoint_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidEndpointTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidEndpointTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidEndpointTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidEndpointTypeRule) Check(runner *tflint.Runner) error {
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
					`endpoint_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
