// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodecommitRepositoryInvalidDefaultBranchRule checks the pattern is valid
type AwsCodecommitRepositoryInvalidDefaultBranchRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodecommitRepositoryInvalidDefaultBranchRule returns new rule with default attributes
func NewAwsCodecommitRepositoryInvalidDefaultBranchRule() *AwsCodecommitRepositoryInvalidDefaultBranchRule {
	return &AwsCodecommitRepositoryInvalidDefaultBranchRule{
		resourceType:  "aws_codecommit_repository",
		attributeName: "default_branch",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Name() string {
	return "aws_codecommit_repository_invalid_default_branch"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Check(runner tflint.Runner) error {
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
					"default_branch must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"default_branch must be 1 characters or higher",
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
