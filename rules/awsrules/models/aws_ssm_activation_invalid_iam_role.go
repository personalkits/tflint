// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmActivationInvalidIAMRoleRule checks the pattern is valid
type AwsSsmActivationInvalidIAMRoleRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSsmActivationInvalidIAMRoleRule returns new rule with default attributes
func NewAwsSsmActivationInvalidIAMRoleRule() *AwsSsmActivationInvalidIAMRoleRule {
	return &AwsSsmActivationInvalidIAMRoleRule{
		resourceType:  "aws_ssm_activation",
		attributeName: "iam_role",
		max:           64,
	}
}

// Name returns the rule name
func (r *AwsSsmActivationInvalidIAMRoleRule) Name() string {
	return "aws_ssm_activation_invalid_iam_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmActivationInvalidIAMRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmActivationInvalidIAMRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmActivationInvalidIAMRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmActivationInvalidIAMRoleRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"iam_role must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
