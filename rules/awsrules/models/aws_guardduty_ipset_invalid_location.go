// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsGuarddutyIpsetInvalidLocationRule checks the pattern is valid
type AwsGuarddutyIpsetInvalidLocationRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyIpsetInvalidLocationRule returns new rule with default attributes
func NewAwsGuarddutyIpsetInvalidLocationRule() *AwsGuarddutyIpsetInvalidLocationRule {
	return &AwsGuarddutyIpsetInvalidLocationRule{
		resourceType:  "aws_guardduty_ipset",
		attributeName: "location",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyIpsetInvalidLocationRule) Name() string {
	return "aws_guardduty_ipset_invalid_location"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyIpsetInvalidLocationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyIpsetInvalidLocationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyIpsetInvalidLocationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyIpsetInvalidLocationRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"location must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"location must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
