// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsChimeVoiceConnectorInvalidNameRule checks the pattern is valid
type AwsChimeVoiceConnectorInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsChimeVoiceConnectorInvalidNameRule returns new rule with default attributes
func NewAwsChimeVoiceConnectorInvalidNameRule() *AwsChimeVoiceConnectorInvalidNameRule {
	return &AwsChimeVoiceConnectorInvalidNameRule{
		resourceType:  "aws_chime_voice_connector",
		attributeName: "name",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsChimeVoiceConnectorInvalidNameRule) Name() string {
	return "aws_chime_voice_connector_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsChimeVoiceConnectorInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsChimeVoiceConnectorInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsChimeVoiceConnectorInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsChimeVoiceConnectorInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
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
