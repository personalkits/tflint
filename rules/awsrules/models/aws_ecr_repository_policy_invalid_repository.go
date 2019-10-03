// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsEcrRepositoryPolicyInvalidRepositoryRule checks the pattern is valid
type AwsEcrRepositoryPolicyInvalidRepositoryRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrRepositoryPolicyInvalidRepositoryRule returns new rule with default attributes
func NewAwsEcrRepositoryPolicyInvalidRepositoryRule() *AwsEcrRepositoryPolicyInvalidRepositoryRule {
	return &AwsEcrRepositoryPolicyInvalidRepositoryRule{
		resourceType:  "aws_ecr_repository_policy",
		attributeName: "repository",
		max:           256,
		min:           2,
		pattern:       regexp.MustCompile(`^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrRepositoryPolicyInvalidRepositoryRule) Name() string {
	return "aws_ecr_repository_policy_invalid_repository"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrRepositoryPolicyInvalidRepositoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrRepositoryPolicyInvalidRepositoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrRepositoryPolicyInvalidRepositoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrRepositoryPolicyInvalidRepositoryRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"repository must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"repository must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`repository does not match valid pattern ^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
