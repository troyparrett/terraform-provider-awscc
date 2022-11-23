// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigatewayv2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSApiGatewayV2Authorizer_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGatewayV2::Authorizer", "awscc_apigatewayv2_authorizer", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
