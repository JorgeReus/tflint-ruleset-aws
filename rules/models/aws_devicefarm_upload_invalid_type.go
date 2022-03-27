// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDevicefarmUploadInvalidTypeRule checks the pattern is valid
type AwsDevicefarmUploadInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDevicefarmUploadInvalidTypeRule returns new rule with default attributes
func NewAwsDevicefarmUploadInvalidTypeRule() *AwsDevicefarmUploadInvalidTypeRule {
	return &AwsDevicefarmUploadInvalidTypeRule{
		resourceType:  "aws_devicefarm_upload",
		attributeName: "type",
		enum: []string{
			"ANDROID_APP",
			"IOS_APP",
			"WEB_APP",
			"EXTERNAL_DATA",
			"APPIUM_JAVA_JUNIT_TEST_PACKAGE",
			"APPIUM_JAVA_TESTNG_TEST_PACKAGE",
			"APPIUM_PYTHON_TEST_PACKAGE",
			"APPIUM_NODE_TEST_PACKAGE",
			"APPIUM_RUBY_TEST_PACKAGE",
			"APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE",
			"APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE",
			"APPIUM_WEB_PYTHON_TEST_PACKAGE",
			"APPIUM_WEB_NODE_TEST_PACKAGE",
			"APPIUM_WEB_RUBY_TEST_PACKAGE",
			"CALABASH_TEST_PACKAGE",
			"INSTRUMENTATION_TEST_PACKAGE",
			"UIAUTOMATION_TEST_PACKAGE",
			"UIAUTOMATOR_TEST_PACKAGE",
			"XCTEST_TEST_PACKAGE",
			"XCTEST_UI_TEST_PACKAGE",
			"APPIUM_JAVA_JUNIT_TEST_SPEC",
			"APPIUM_JAVA_TESTNG_TEST_SPEC",
			"APPIUM_PYTHON_TEST_SPEC",
			"APPIUM_NODE_TEST_SPEC",
			"APPIUM_RUBY_TEST_SPEC",
			"APPIUM_WEB_JAVA_JUNIT_TEST_SPEC",
			"APPIUM_WEB_JAVA_TESTNG_TEST_SPEC",
			"APPIUM_WEB_PYTHON_TEST_SPEC",
			"APPIUM_WEB_NODE_TEST_SPEC",
			"APPIUM_WEB_RUBY_TEST_SPEC",
			"INSTRUMENTATION_TEST_SPEC",
			"XCTEST_UI_TEST_SPEC",
		},
	}
}

// Name returns the rule name
func (r *AwsDevicefarmUploadInvalidTypeRule) Name() string {
	return "aws_devicefarm_upload_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmUploadInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDevicefarmUploadInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmUploadInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmUploadInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
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
