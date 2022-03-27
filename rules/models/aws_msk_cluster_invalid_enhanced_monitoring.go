// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMskClusterInvalidEnhancedMonitoringRule checks the pattern is valid
type AwsMskClusterInvalidEnhancedMonitoringRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMskClusterInvalidEnhancedMonitoringRule returns new rule with default attributes
func NewAwsMskClusterInvalidEnhancedMonitoringRule() *AwsMskClusterInvalidEnhancedMonitoringRule {
	return &AwsMskClusterInvalidEnhancedMonitoringRule{
		resourceType:  "aws_msk_cluster",
		attributeName: "enhanced_monitoring",
		enum: []string{
			"DEFAULT",
			"PER_BROKER",
			"PER_TOPIC_PER_BROKER",
			"PER_TOPIC_PER_PARTITION",
		},
	}
}

// Name returns the rule name
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Name() string {
	return "aws_msk_cluster_invalid_enhanced_monitoring"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMskClusterInvalidEnhancedMonitoringRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as enhanced_monitoring`, truncateLongMessage(val)),
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
