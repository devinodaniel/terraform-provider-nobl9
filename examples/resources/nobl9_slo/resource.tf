resource "nobl9_project" "this" {
  display_name = "Test Terraform"
  name         = "test-terraform"
  description  = "An example terraform project"
}

resource "nobl9_service" "this" {
  name         = "foo-front-page"
  project      = nobl9_project.this.name
  display_name = "${nobl9_project.this.display_name} Front Page"
  description  = "Front page service"
}

resource "nobl9_slo" "this" {
  name             = "${nobl9_project.this.name}-latency"
  service          = nobl9_service.this.name
  budgeting_method = "Occurrences"
  project          = nobl9_project.this.name

  alert_policies = [
    "foo-front-page-latency"
  ]

  time_window {
    unit       = "Day"
    count      = 30
    is_rolling = true
  }

  objective {
    target       = 0.99
    display_name = "OK"
    value        = 2000
    op           = "gte"
  }

  indicator {
    name = "test-terraform-prom-agent"
    raw_metric {
      prometheus {
        promql = <<EOT
        latency_west_c7{code="ALL",instance="localhost:3000",job="prometheus",service="globacount"}
        EOT
      }
    }
  }
}

