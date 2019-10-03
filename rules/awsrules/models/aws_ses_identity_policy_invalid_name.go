// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSesIdentityPolicyInvalidNameRule checks the pattern is valid
type AwsSesIdentityPolicyInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSesIdentityPolicyInvalidNameRule returns new rule with default attributes
func NewAwsSesIdentityPolicyInvalidNameRule() *AwsSesIdentityPolicyInvalidNameRule {
	return &AwsSesIdentityPolicyInvalidNameRule{
		resourceType:  "aws_ses_identity_policy",
		attributeName: "name",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSesIdentityPolicyInvalidNameRule) Name() string {
	return "aws_ses_identity_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSesIdentityPolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSesIdentityPolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSesIdentityPolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSesIdentityPolicyInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
