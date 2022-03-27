// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOpenzfsFileSystemInvalidStorageTypeRule checks the pattern is valid
type AwsFsxOpenzfsFileSystemInvalidStorageTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsFsxOpenzfsFileSystemInvalidStorageTypeRule returns new rule with default attributes
func NewAwsFsxOpenzfsFileSystemInvalidStorageTypeRule() *AwsFsxOpenzfsFileSystemInvalidStorageTypeRule {
	return &AwsFsxOpenzfsFileSystemInvalidStorageTypeRule{
		resourceType:  "aws_fsx_openzfs_file_system",
		attributeName: "storage_type",
		enum: []string{
			"SSD",
			"HDD",
		},
	}
}

// Name returns the rule name
func (r *AwsFsxOpenzfsFileSystemInvalidStorageTypeRule) Name() string {
	return "aws_fsx_openzfs_file_system_invalid_storage_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOpenzfsFileSystemInvalidStorageTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOpenzfsFileSystemInvalidStorageTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOpenzfsFileSystemInvalidStorageTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOpenzfsFileSystemInvalidStorageTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as storage_type`, truncateLongMessage(val)),
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
