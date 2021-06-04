package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/atest"
)

func TestAccAWSCodeCommitRepositoryDataSource_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acctest-%d", acctest.RandInt())
	resourceName := "aws_codecommit_repository.default"
	datasourceName := "data.aws_codecommit_repository.default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { atest.PreCheck(t) },
		ErrorCheck: atest.ErrorCheck(t, codecommit.EndpointsID),
		Providers:  atest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsCodeCommitRepositoryDataSourceConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "clone_url_http", resourceName, "clone_url_http"),
					resource.TestCheckResourceAttrPair(datasourceName, "clone_url_ssh", resourceName, "clone_url_ssh"),
					resource.TestCheckResourceAttrPair(datasourceName, "repository_name", resourceName, "repository_name"),
				),
			},
		},
	})
}

func testAccCheckAwsCodeCommitRepositoryDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "aws_codecommit_repository" "default" {
  repository_name = "%s"
}

data "aws_codecommit_repository" "default" {
  repository_name = aws_codecommit_repository.default.repository_name
}
`, rName)
}
