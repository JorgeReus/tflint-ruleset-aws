// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodebuildProjectInvalidDescriptionRule checks the pattern is valid
type AwsCodebuildProjectInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsCodebuildProjectInvalidDescriptionRule returns new rule with default attributes
func NewAwsCodebuildProjectInvalidDescriptionRule() *AwsCodebuildProjectInvalidDescriptionRule {
	return &AwsCodebuildProjectInvalidDescriptionRule{
		resourceType:  "aws_codebuild_project",
		attributeName: "description",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsCodebuildProjectInvalidDescriptionRule) Name() string {
	return "aws_codebuild_project_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodebuildProjectInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodebuildProjectInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodebuildProjectInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodebuildProjectInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 255 characters or less",
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
