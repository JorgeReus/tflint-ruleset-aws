// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTimestreamwriteTableInvalidDatabaseNameRule checks the pattern is valid
type AwsTimestreamwriteTableInvalidDatabaseNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsTimestreamwriteTableInvalidDatabaseNameRule returns new rule with default attributes
func NewAwsTimestreamwriteTableInvalidDatabaseNameRule() *AwsTimestreamwriteTableInvalidDatabaseNameRule {
	return &AwsTimestreamwriteTableInvalidDatabaseNameRule{
		resourceType:  "aws_timestreamwrite_table",
		attributeName: "database_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsTimestreamwriteTableInvalidDatabaseNameRule) Name() string {
	return "aws_timestreamwrite_table_invalid_database_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTimestreamwriteTableInvalidDatabaseNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTimestreamwriteTableInvalidDatabaseNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTimestreamwriteTableInvalidDatabaseNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTimestreamwriteTableInvalidDatabaseNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_.-]+$`),
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
