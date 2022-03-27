// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayWorkingStorageInvalidDiskIDRule checks the pattern is valid
type AwsStoragegatewayWorkingStorageInvalidDiskIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayWorkingStorageInvalidDiskIDRule returns new rule with default attributes
func NewAwsStoragegatewayWorkingStorageInvalidDiskIDRule() *AwsStoragegatewayWorkingStorageInvalidDiskIDRule {
	return &AwsStoragegatewayWorkingStorageInvalidDiskIDRule{
		resourceType:  "aws_storagegateway_working_storage",
		attributeName: "disk_id",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayWorkingStorageInvalidDiskIDRule) Name() string {
	return "aws_storagegateway_working_storage_invalid_disk_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayWorkingStorageInvalidDiskIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayWorkingStorageInvalidDiskIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayWorkingStorageInvalidDiskIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayWorkingStorageInvalidDiskIDRule) Check(runner tflint.Runner) error {
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
					"disk_id must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"disk_id must be 1 characters or higher",
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
