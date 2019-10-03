// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsIotThingInvalidThingTypeNameRule checks the pattern is valid
type AwsIotThingInvalidThingTypeNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIotThingInvalidThingTypeNameRule returns new rule with default attributes
func NewAwsIotThingInvalidThingTypeNameRule() *AwsIotThingInvalidThingTypeNameRule {
	return &AwsIotThingInvalidThingTypeNameRule{
		resourceType:  "aws_iot_thing",
		attributeName: "thing_type_name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9:_-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIotThingInvalidThingTypeNameRule) Name() string {
	return "aws_iot_thing_invalid_thing_type_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIotThingInvalidThingTypeNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIotThingInvalidThingTypeNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIotThingInvalidThingTypeNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIotThingInvalidThingTypeNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"thing_type_name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"thing_type_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`thing_type_name does not match valid pattern ^[a-zA-Z0-9:_-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
