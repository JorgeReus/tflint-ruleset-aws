// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyAppInvalidAccessTokenRule checks the pattern is valid
type AwsAmplifyAppInvalidAccessTokenRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyAppInvalidAccessTokenRule returns new rule with default attributes
func NewAwsAmplifyAppInvalidAccessTokenRule() *AwsAmplifyAppInvalidAccessTokenRule {
	return &AwsAmplifyAppInvalidAccessTokenRule{
		resourceType:  "aws_amplify_app",
		attributeName: "access_token",
		max:           255,
		min:           1,
		pattern:       regexp.MustCompile(`^(?s).+$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyAppInvalidAccessTokenRule) Name() string {
	return "aws_amplify_app_invalid_access_token"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyAppInvalidAccessTokenRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyAppInvalidAccessTokenRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyAppInvalidAccessTokenRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyAppInvalidAccessTokenRule) Check(runner tflint.Runner) error {
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
					"access_token must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"access_token must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`access_token does not match valid pattern ^(?s).+$`,
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
