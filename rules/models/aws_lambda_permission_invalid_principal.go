// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaPermissionInvalidPrincipalRule checks the pattern is valid
type AwsLambdaPermissionInvalidPrincipalRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidPrincipalRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidPrincipalRule() *AwsLambdaPermissionInvalidPrincipalRule {
	return &AwsLambdaPermissionInvalidPrincipalRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "principal",
		pattern:       regexp.MustCompile(`^[^\s]+$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidPrincipalRule) Name() string {
	return "aws_lambda_permission_invalid_principal"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidPrincipalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidPrincipalRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidPrincipalRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidPrincipalRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\s]+$`),
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
