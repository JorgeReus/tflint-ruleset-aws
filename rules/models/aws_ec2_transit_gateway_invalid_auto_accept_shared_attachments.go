// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule checks the pattern is valid
type AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule returns new rule with default attributes
func NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule() *AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule {
	return &AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule{
		resourceType:  "aws_ec2_transit_gateway",
		attributeName: "auto_accept_shared_attachments",
		enum: []string{
			"enable",
			"disable",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule) Name() string {
	return "aws_ec2_transit_gateway_invalid_auto_accept_shared_attachments"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as auto_accept_shared_attachments`, truncateLongMessage(val)),
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
