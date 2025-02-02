---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nobl9 Provider"
subcategory: ""
description: |-
  
---

# nobl9 Provider



## Example Usage

```terraform
terraform {
  required_providers {
    nobl9 = {
      source  = "nobl9.com/nobl9/nobl9"
      version = "0.1.0"
    }
  }
}

provider "nobl9" {
  organization  = "<your org name>"
  project       = "default"
  client_id     = "<client_id>"
  client_secret = "<client_secret>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **client_id** (String) Authentication parameter ClientID.
- **client_secret** (String, Sensitive) Authentication parameter ClientSecret.
- **ingest_url** (String) Nobl9 API URL.
- **okta_auth_server** (String) Authorization service configuration.
- **okta_org_url** (String) Authorization service URL.
- **organization** (String) Nobl9 Organization that contain resources managed by this provider.
- **project** (String) Nobl9 project used when importing resources.
