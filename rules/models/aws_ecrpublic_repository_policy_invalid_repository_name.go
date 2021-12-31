// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule checks the pattern is valid
type AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule returns new rule with default attributes
func NewAwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule() *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule {
	return &AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule{
		resourceType:  "aws_ecrpublic_repository_policy",
		attributeName: "repository_name",
		max:           205,
		min:           2,
		pattern:       regexp.MustCompile(`^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Name() string {
	return "aws_ecrpublic_repository_policy_invalid_repository_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrpublicRepositoryPolicyInvalidRepositoryNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"repository_name must be 205 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"repository_name must be 2 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
