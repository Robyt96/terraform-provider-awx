---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_workflow_job_template Resource - terraform-provider-awx"
subcategory: ""
description: |-
  
---

# awx_workflow_job_template (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of this workflow job template. (string, required)

### Optional

- `allow_simultaneous` (Boolean)
- `ask_inventory_on_launch` (Boolean)
- `ask_limit_on_launch` (Boolean)
- `ask_scm_branch_on_launch` (Boolean)
- `ask_variables_on_launch` (Boolean)
- `description` (String) Optional description of this workflow job template.
- `inventory_id` (String) Inventory applied as a prompt, assuming job template prompts for inventory. (id, default=``)
- `limit` (String)
- `organization_id` (Number) The organization used to determine access to this template. (id, default=``)
- `scm_branch` (String)
- `survey_enabled` (Boolean)
- `variables` (String)
- `webhook_credential` (String)
- `webhook_service` (String)

### Read-Only

- `id` (String) The ID of this resource.
