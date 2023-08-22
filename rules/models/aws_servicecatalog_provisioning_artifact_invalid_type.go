// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProvisioningArtifactInvalidTypeRule checks the pattern is valid
type AwsServicecatalogProvisioningArtifactInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsServicecatalogProvisioningArtifactInvalidTypeRule returns new rule with default attributes
func NewAwsServicecatalogProvisioningArtifactInvalidTypeRule() *AwsServicecatalogProvisioningArtifactInvalidTypeRule {
	return &AwsServicecatalogProvisioningArtifactInvalidTypeRule{
		resourceType:  "aws_servicecatalog_provisioning_artifact",
		attributeName: "type",
		enum: []string{
			"CLOUD_FORMATION_TEMPLATE",
			"MARKETPLACE_AMI",
			"MARKETPLACE_CAR",
			"TERRAFORM_OPEN_SOURCE",
			"TERRAFORM_CLOUD",
		},
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProvisioningArtifactInvalidTypeRule) Name() string {
	return "aws_servicecatalog_provisioning_artifact_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProvisioningArtifactInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProvisioningArtifactInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProvisioningArtifactInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProvisioningArtifactInvalidTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

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

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
