// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchLogGroupInvalidKmsKeyIDRule checks the pattern is valid
type AwsCloudwatchLogGroupInvalidKmsKeyIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchLogGroupInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsCloudwatchLogGroupInvalidKmsKeyIDRule() *AwsCloudwatchLogGroupInvalidKmsKeyIDRule {
	return &AwsCloudwatchLogGroupInvalidKmsKeyIDRule{
		resourceType:  "aws_cloudwatch_log_group",
		attributeName: "kms_key_id",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogGroupInvalidKmsKeyIDRule) Name() string {
	return "aws_cloudwatch_log_group_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogGroupInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogGroupInvalidKmsKeyIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogGroupInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogGroupInvalidKmsKeyIDRule) Check(runner tflint.Runner) error {
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
					"kms_key_id must be 256 characters or less",
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
