// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKmsAliasInvalidTargetKeyIDRule checks the pattern is valid
type AwsKmsAliasInvalidTargetKeyIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsKmsAliasInvalidTargetKeyIDRule returns new rule with default attributes
func NewAwsKmsAliasInvalidTargetKeyIDRule() *AwsKmsAliasInvalidTargetKeyIDRule {
	return &AwsKmsAliasInvalidTargetKeyIDRule{
		resourceType:  "aws_kms_alias",
		attributeName: "target_key_id",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Name() string {
	return "aws_kms_alias_invalid_target_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKmsAliasInvalidTargetKeyIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"target_key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_key_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
