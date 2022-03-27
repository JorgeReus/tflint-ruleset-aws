// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightGroupInvalidAwsAccountIDRule checks the pattern is valid
type AwsQuicksightGroupInvalidAwsAccountIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightGroupInvalidAwsAccountIDRule returns new rule with default attributes
func NewAwsQuicksightGroupInvalidAwsAccountIDRule() *AwsQuicksightGroupInvalidAwsAccountIDRule {
	return &AwsQuicksightGroupInvalidAwsAccountIDRule{
		resourceType:  "aws_quicksight_group",
		attributeName: "aws_account_id",
		max:           12,
		min:           12,
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightGroupInvalidAwsAccountIDRule) Name() string {
	return "aws_quicksight_group_invalid_aws_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightGroupInvalidAwsAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightGroupInvalidAwsAccountIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightGroupInvalidAwsAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightGroupInvalidAwsAccountIDRule) Check(runner tflint.Runner) error {
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
					"aws_account_id must be 12 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"aws_account_id must be 12 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
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
