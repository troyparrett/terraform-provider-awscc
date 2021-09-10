// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package groundstation_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGroundStationMissionProfileDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GroundStation::MissionProfile", "awscc_groundstation_mission_profile", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSGroundStationMissionProfileDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GroundStation::MissionProfile", "awscc_groundstation_mission_profile", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}