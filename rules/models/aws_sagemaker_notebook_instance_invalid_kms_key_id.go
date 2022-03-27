// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule checks the pattern is valid
type AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerNotebookInstanceInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsSagemakerNotebookInstanceInvalidKmsKeyIDRule() *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule {
	return &AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule{
		resourceType:  "aws_sagemaker_notebook_instance",
		attributeName: "kms_key_id",
		max:           2048,
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Name() string {
	return "aws_sagemaker_notebook_instance_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Check(runner tflint.Runner) error {
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
					"kms_key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*$`),
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
