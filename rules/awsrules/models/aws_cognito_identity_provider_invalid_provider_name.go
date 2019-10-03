// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsCognitoIdentityProviderInvalidProviderNameRule checks the pattern is valid
type AwsCognitoIdentityProviderInvalidProviderNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoIdentityProviderInvalidProviderNameRule returns new rule with default attributes
func NewAwsCognitoIdentityProviderInvalidProviderNameRule() *AwsCognitoIdentityProviderInvalidProviderNameRule {
	return &AwsCognitoIdentityProviderInvalidProviderNameRule{
		resourceType:  "aws_cognito_identity_provider",
		attributeName: "provider_name",
		max:           32,
		min:           1,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{S}\p{N}\p{P}]+$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoIdentityProviderInvalidProviderNameRule) Name() string {
	return "aws_cognito_identity_provider_invalid_provider_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoIdentityProviderInvalidProviderNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoIdentityProviderInvalidProviderNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoIdentityProviderInvalidProviderNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoIdentityProviderInvalidProviderNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"provider_name must be 32 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"provider_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`provider_name does not match valid pattern ^[\p{L}\p{M}\p{S}\p{N}\p{P}]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
