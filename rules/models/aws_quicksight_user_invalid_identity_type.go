// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightUserInvalidIdentityTypeRule checks the pattern is valid
type AwsQuicksightUserInvalidIdentityTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsQuicksightUserInvalidIdentityTypeRule returns new rule with default attributes
func NewAwsQuicksightUserInvalidIdentityTypeRule() *AwsQuicksightUserInvalidIdentityTypeRule {
	return &AwsQuicksightUserInvalidIdentityTypeRule{
		resourceType:  "aws_quicksight_user",
		attributeName: "identity_type",
		enum: []string{
			"IAM",
			"QUICKSIGHT",
			"IAM_IDENTITY_CENTER",
		},
	}
}

// Name returns the rule name
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Name() string {
	return "aws_quicksight_user_invalid_identity_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

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

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as identity_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
