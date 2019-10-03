// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsTransferUserInvalidHomeDirectoryRule checks the pattern is valid
type AwsTransferUserInvalidHomeDirectoryRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsTransferUserInvalidHomeDirectoryRule returns new rule with default attributes
func NewAwsTransferUserInvalidHomeDirectoryRule() *AwsTransferUserInvalidHomeDirectoryRule {
	return &AwsTransferUserInvalidHomeDirectoryRule{
		resourceType:  "aws_transfer_user",
		attributeName: "home_directory",
		max:           1024,
		pattern:       regexp.MustCompile(`^$|/.*`),
	}
}

// Name returns the rule name
func (r *AwsTransferUserInvalidHomeDirectoryRule) Name() string {
	return "aws_transfer_user_invalid_home_directory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferUserInvalidHomeDirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferUserInvalidHomeDirectoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferUserInvalidHomeDirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferUserInvalidHomeDirectoryRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"home_directory must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`home_directory does not match valid pattern ^$|/.*`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
