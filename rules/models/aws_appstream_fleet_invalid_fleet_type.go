// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamFleetInvalidFleetTypeRule checks the pattern is valid
type AwsAppstreamFleetInvalidFleetTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppstreamFleetInvalidFleetTypeRule returns new rule with default attributes
func NewAwsAppstreamFleetInvalidFleetTypeRule() *AwsAppstreamFleetInvalidFleetTypeRule {
	return &AwsAppstreamFleetInvalidFleetTypeRule{
		resourceType:  "aws_appstream_fleet",
		attributeName: "fleet_type",
		enum: []string{
			"ALWAYS_ON",
			"ON_DEMAND",
			"ELASTIC",
		},
	}
}

// Name returns the rule name
func (r *AwsAppstreamFleetInvalidFleetTypeRule) Name() string {
	return "aws_appstream_fleet_invalid_fleet_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamFleetInvalidFleetTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamFleetInvalidFleetTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamFleetInvalidFleetTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamFleetInvalidFleetTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as fleet_type`, truncateLongMessage(val)),
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
