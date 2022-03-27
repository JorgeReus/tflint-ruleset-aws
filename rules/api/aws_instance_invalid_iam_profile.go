// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsInstanceInvalidIAMProfileRule checks whether attribute value actually exists
type AwsInstanceInvalidIAMProfileRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsInstanceInvalidIAMProfileRule returns new rule with default attributes
func NewAwsInstanceInvalidIAMProfileRule() *AwsInstanceInvalidIAMProfileRule {
	return &AwsInstanceInvalidIAMProfileRule{
		resourceType:  "aws_instance",
		attributeName: "iam_instance_profile",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsInstanceInvalidIAMProfileRule) Name() string {
	return "aws_instance_invalid_iam_profile"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceInvalidIAMProfileRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceInvalidIAMProfileRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceInvalidIAMProfileRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsInstanceInvalidIAMProfileRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by ListInstanceProfiles
func (r *AwsInstanceInvalidIAMProfileRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

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

		if !r.dataPrepared {
			log.Print("[DEBUG] invoking ListInstanceProfiles")
			var err error
			r.data, err = runner.AwsClient.ListInstanceProfiles()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking ListInstanceProfiles; %w", err)
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid IAM profile name.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	}

	return nil
}
