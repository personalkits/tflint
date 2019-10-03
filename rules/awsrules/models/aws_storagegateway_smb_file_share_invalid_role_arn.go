// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewaySmbFileShareInvalidRoleArnRule checks the pattern is valid
type AwsStoragegatewaySmbFileShareInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewaySmbFileShareInvalidRoleArnRule returns new rule with default attributes
func NewAwsStoragegatewaySmbFileShareInvalidRoleArnRule() *AwsStoragegatewaySmbFileShareInvalidRoleArnRule {
	return &AwsStoragegatewaySmbFileShareInvalidRoleArnRule{
		resourceType:  "aws_storagegateway_smb_file_share",
		attributeName: "role_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Name() string {
	return "aws_storagegateway_smb_file_share_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewaySmbFileShareInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
