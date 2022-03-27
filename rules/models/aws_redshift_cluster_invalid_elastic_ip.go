// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftClusterInvalidElasticIPRule checks the pattern is valid
type AwsRedshiftClusterInvalidElasticIPRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftClusterInvalidElasticIPRule returns new rule with default attributes
func NewAwsRedshiftClusterInvalidElasticIPRule() *AwsRedshiftClusterInvalidElasticIPRule {
	return &AwsRedshiftClusterInvalidElasticIPRule{
		resourceType:  "aws_redshift_cluster",
		attributeName: "elastic_ip",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftClusterInvalidElasticIPRule) Name() string {
	return "aws_redshift_cluster_invalid_elastic_ip"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftClusterInvalidElasticIPRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftClusterInvalidElasticIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftClusterInvalidElasticIPRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftClusterInvalidElasticIPRule) Check(runner tflint.Runner) error {
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
					"elastic_ip must be 2147483647 characters or less",
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
