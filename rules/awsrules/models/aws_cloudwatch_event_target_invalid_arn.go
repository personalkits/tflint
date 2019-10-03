// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchEventTargetInvalidArnRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchEventTargetInvalidArnRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidArnRule() *AwsCloudwatchEventTargetInvalidArnRule {
	return &AwsCloudwatchEventTargetInvalidArnRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "arn",
		max:           1600,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidArnRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"arn must be 1600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
