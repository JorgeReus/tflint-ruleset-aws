// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerWorkteamInvalidWorkteamNameRule checks the pattern is valid
type AwsSagemakerWorkteamInvalidWorkteamNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerWorkteamInvalidWorkteamNameRule returns new rule with default attributes
func NewAwsSagemakerWorkteamInvalidWorkteamNameRule() *AwsSagemakerWorkteamInvalidWorkteamNameRule {
	return &AwsSagemakerWorkteamInvalidWorkteamNameRule{
		resourceType:  "aws_sagemaker_workteam",
		attributeName: "workteam_name",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerWorkteamInvalidWorkteamNameRule) Name() string {
	return "aws_sagemaker_workteam_invalid_workteam_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerWorkteamInvalidWorkteamNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerWorkteamInvalidWorkteamNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerWorkteamInvalidWorkteamNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerWorkteamInvalidWorkteamNameRule) Check(runner tflint.Runner) error {
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
					"workteam_name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"workteam_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`),
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
