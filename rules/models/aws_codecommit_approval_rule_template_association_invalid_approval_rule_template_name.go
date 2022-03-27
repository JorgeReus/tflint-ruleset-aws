// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule checks the pattern is valid
type AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule returns new rule with default attributes
func NewAwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule() *AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule {
	return &AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule{
		resourceType:  "aws_codecommit_approval_rule_template_association",
		attributeName: "approval_rule_template_name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule) Name() string {
	return "aws_codecommit_approval_rule_template_association_invalid_approval_rule_template_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodecommitApprovalRuleTemplateAssociationInvalidApprovalRuleTemplateNameRule) Check(runner tflint.Runner) error {
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
					"approval_rule_template_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"approval_rule_template_name must be 1 characters or higher",
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
