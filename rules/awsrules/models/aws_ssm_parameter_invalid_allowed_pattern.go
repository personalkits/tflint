// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmParameterInvalidAllowedPatternRule checks the pattern is valid
type AwsSsmParameterInvalidAllowedPatternRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSsmParameterInvalidAllowedPatternRule returns new rule with default attributes
func NewAwsSsmParameterInvalidAllowedPatternRule() *AwsSsmParameterInvalidAllowedPatternRule {
	return &AwsSsmParameterInvalidAllowedPatternRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "allowed_pattern",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidAllowedPatternRule) Name() string {
	return "aws_ssm_parameter_invalid_allowed_pattern"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidAllowedPatternRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidAllowedPatternRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidAllowedPatternRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidAllowedPatternRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"allowed_pattern must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
