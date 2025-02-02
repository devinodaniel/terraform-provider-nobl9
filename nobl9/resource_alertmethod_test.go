package nobl9

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	n9api "github.com/nobl9/nobl9-go"
)

func TestAcc_Nobl9AlertMethod(t *testing.T) {
	cases := []struct {
		name           string
		resourceSuffix string
		configFunc     func(string) string
	}{
		{"test-webhook", "webhook", testWebhookTemplateConfig},
		{"test-webhook-fields", "webhook", testWebhookTemplateFieldsConfig},
		{"test-pagerduty", "pagerduty", testPagerDutyConfig},
		{"test-slack", "slack", testSlackConfig},
		{"test-discord", "discord", testDiscordConfig},
		{"test-opsgenie", "opsgenie", testOpsgenieConfig},
		{"test-servicenow", "servicenow", testServiceNowConfig},
		{"test-jira", "jira", testJiraConfig},
		{"test-teams", "msteams", testTeamsConfig},
		{"test-email", "email", testEmailConfig},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: ProviderFactory(),
				CheckDestroy:      CheckDestory("nobl9_alert_method_"+tc.resourceSuffix, n9api.ObjectAlertMethod),
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(tc.name),
						Check:  CheckObjectCreated(fmt.Sprintf("nobl9_alert_method_%s.%s", tc.resourceSuffix, tc.name)),
					},
				},
			})
		})
	}
}

func testWebhookTemplateConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_webhook" "%s" {
  name        = "%s"
  project     = "%s"
  description = "wehbook"
  url         = "http://web.net"
  template    = "SLO needs attention $slo_name"
}
`, name, name, testProject)
}

func testWebhookTemplateFieldsConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_webhook" "%s" {
  name            = "%s"
  project         = "%s"
  description	  = "wehbook"
  url             = "http://web.net"
  template_fields = [ "slo_name", "slo_details_link" ]
}
`, name, name, testProject)
}

func testPagerDutyConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_pagerduty" "%s" {
  name            = "%s"
  project         = "%s"
  description     = "paderduty"
  integration_key = "84dfcdf19dad8f6c82b7e22afa024065"
}
`, name, name, testProject)
}

func testSlackConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_slack" "%s" {
  name        = "%s"
  project     = "%s"
  description = "slack"
  url         = "https://slack.com"
}
`, name, name, testProject)
}

func testDiscordConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_discord" "%s" {
  name        = "%s"
  project     = "%s"
  description = "discord"
  url         = "https://discord.com"
}
`, name, name, testProject)
}

func testOpsgenieConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_opsgenie" "%s" {
  name        = "%s"
  project     = "%s"
  description = "opsgenie"
  url         = "https://discord.com"
  auth		  = "GenieKey 12345"
}
`, name, name, testProject)
}

func testServiceNowConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_servicenow" "%s" {
  name        = "%s"
  project     = "%s"
  description = "servicenow"
  username    = "nobleUser"
  password    = "very sercret"
  instanceid  = "1"
}
`, name, name, testProject)
}

func testJiraConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_jira" "%s" {
  name        = "%s"
  project     = "%s"
  description = "jira"
  url		  = "https://jira.com"
  username    = "nobleUser"
  apitoken    = "very sercret"
  project_key = "PC"
}
`, name, name, testProject)
}

func testTeamsConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_msteams" "%s" {
  name        = "%s"
  project     = "%s"
  description = "teams"
  url		  = "https://teams.com"
}
`, name, name, testProject)
}

func testEmailConfig(name string) string {
	return fmt.Sprintf(`
resource "nobl9_alert_method_email" "%s" {
  name        = "%s"
  project     = "%s"
  description = "teams"
  to		  = [ "testUser@nobl9.com" ]
  cc		  = [ "testUser@nobl9.com" ]
  bcc		  = [ "testUser@nobl9.com" ]
  subject     = "Test email please ignore"
  body        = "This is just a test email"
}
`, name, name, testProject)
}
