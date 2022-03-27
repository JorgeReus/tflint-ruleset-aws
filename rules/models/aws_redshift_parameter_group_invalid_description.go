// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftParameterGroupInvalidDescriptionRule checks the pattern is valid
type AwsRedshiftParameterGroupInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftParameterGroupInvalidDescriptionRule returns new rule with default attributes
func NewAwsRedshiftParameterGroupInvalidDescriptionRule() *AwsRedshiftParameterGroupInvalidDescriptionRule {
	return &AwsRedshiftParameterGroupInvalidDescriptionRule{
		resourceType:  "aws_redshift_parameter_group",
		attributeName: "description",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftParameterGroupInvalidDescriptionRule) Name() string {
	return "aws_redshift_parameter_group_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftParameterGroupInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftParameterGroupInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftParameterGroupInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftParameterGroupInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 2147483647 characters or less",
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
