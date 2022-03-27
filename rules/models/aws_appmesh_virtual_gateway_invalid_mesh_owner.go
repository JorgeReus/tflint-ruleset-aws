// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppmeshVirtualGatewayInvalidMeshOwnerRule checks the pattern is valid
type AwsAppmeshVirtualGatewayInvalidMeshOwnerRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppmeshVirtualGatewayInvalidMeshOwnerRule returns new rule with default attributes
func NewAwsAppmeshVirtualGatewayInvalidMeshOwnerRule() *AwsAppmeshVirtualGatewayInvalidMeshOwnerRule {
	return &AwsAppmeshVirtualGatewayInvalidMeshOwnerRule{
		resourceType:  "aws_appmesh_virtual_gateway",
		attributeName: "mesh_owner",
		max:           12,
		min:           12,
	}
}

// Name returns the rule name
func (r *AwsAppmeshVirtualGatewayInvalidMeshOwnerRule) Name() string {
	return "aws_appmesh_virtual_gateway_invalid_mesh_owner"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppmeshVirtualGatewayInvalidMeshOwnerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppmeshVirtualGatewayInvalidMeshOwnerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppmeshVirtualGatewayInvalidMeshOwnerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppmeshVirtualGatewayInvalidMeshOwnerRule) Check(runner tflint.Runner) error {
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
					"mesh_owner must be 12 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"mesh_owner must be 12 characters or higher",
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
