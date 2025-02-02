---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nobl9_agent Resource - terraform-provider-nobl9"
subcategory: ""
description: |-
  Agent configuration documentation https://docs.nobl9.com/the-nobl9-agent
---

# nobl9_agent (Resource)

[Agent configuration documentation](https://docs.nobl9.com/the-nobl9-agent)



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **agent_type** (String) Type of an agent. [Supported agent types](https://docs.nobl9.com/the-nobl9-agent)
- **name** (String) Unique name of the resource. Must match [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
- **project** (String) Name of the project the resource is in. Must match [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
- **source_of** (List of String) Source of Metrics and/or Services

### Optional

- **appdynamics_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/appdynamics#appdynamics-agent) (see [below for nested schema](#nestedblock--appdynamics_config))
- **bigquery_config** (Set of String) [Configuration documentation](https://docs.nobl9.com/Sources/bigquery#bigquery-agent)
- **datadog_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/datadog#datadog-agent) (see [below for nested schema](#nestedblock--datadog_config))
- **description** (String) Optional description of the resource.
- **display_name** (String) Display name of the resource.
- **dynatrace_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/dynatrace#dynatrace-agent) (see [below for nested schema](#nestedblock--dynatrace_config))
- **graphite_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/graphite#graphite-agent) (see [below for nested schema](#nestedblock--graphite_config))
- **id** (String) The ID of this resource.
- **lightstep_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/lightstep#lightstep-agent) (see [below for nested schema](#nestedblock--lightstep_config))
- **newrelic_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/new-relic#new-relic-agent) (see [below for nested schema](#nestedblock--newrelic_config))
- **opentsdb_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/opentsdb#opentsdb-agent) (see [below for nested schema](#nestedblock--opentsdb_config))
- **prometheus_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/prometheus#prometheus-agent) (see [below for nested schema](#nestedblock--prometheus_config))
- **splunk_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/splunk#splunk-agent) (see [below for nested schema](#nestedblock--splunk_config))
- **splunk_observability_config** (Block Set, Max: 1) [Configuration documentation](https://docs.nobl9.com/Sources/splunk-observability) (see [below for nested schema](#nestedblock--splunk_observability_config))
- **thousandeyes_config** (Set of String) [Configuration documentation](https://docs.nobl9.com/Sources/thousandeyes#thousandeyes-agent)

### Read-Only

- **status** (Map of String) Status of created agent.

<a id="nestedblock--appdynamics_config"></a>
### Nested Schema for `appdynamics_config`

Required:

- **url** (String) Base URL to a AppDynamics Controller.


<a id="nestedblock--datadog_config"></a>
### Nested Schema for `datadog_config`

Required:

- **site** (String) `com` or `eu`, Datadog SaaS instance, which corresponds to one of their two locations (https://www.datadoghq.com/ in the U.S. or https://datadoghq.eu/ in the European Union)


<a id="nestedblock--dynatrace_config"></a>
### Nested Schema for `dynatrace_config`

Required:

- **url** (String) Dynatrace API URL.


<a id="nestedblock--graphite_config"></a>
### Nested Schema for `graphite_config`

Required:

- **url** (String) API URL endpoint of Graphite's instance.


<a id="nestedblock--lightstep_config"></a>
### Nested Schema for `lightstep_config`

Required:

- **organization** (String) Organization name registered in Lightstep.
- **project** (String) Name of the Lightstep project.


<a id="nestedblock--newrelic_config"></a>
### Nested Schema for `newrelic_config`

Required:

- **account_id** (String) ID number assigned to the New Relic user account


<a id="nestedblock--opentsdb_config"></a>
### Nested Schema for `opentsdb_config`

Required:

- **url** (String) OpenTSDB cluster URL.


<a id="nestedblock--prometheus_config"></a>
### Nested Schema for `prometheus_config`

Required:

- **url** (String) Base URL to Prometheus server.


<a id="nestedblock--splunk_config"></a>
### Nested Schema for `splunk_config`

Required:

- **url** (String) Base API URL of the Splunk Search app.


<a id="nestedblock--splunk_observability_config"></a>
### Nested Schema for `splunk_observability_config`

Required:

- **realm** (String) SplunkObservability Realm.


