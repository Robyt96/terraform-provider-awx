---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_workflow_job_template_node_always Resource - terraform-provider-awx"
subcategory: ""
description: |-
  
---

# awx_workflow_job_template_node_always (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String)
- `unified_job_template_id` (Number)
- `workflow_job_template_id` (Number)
- `workflow_job_template_node_id` (Number) The workflow_job_template_node id from with the new node will start

### Optional

- `all_parents_must_converge` (Boolean)
- `diff_mode` (Boolean)
- `extra_data` (String)
- `inventory_id` (Number) Inventory applied as a prompt, assuming job template prompts for inventory.
- `job_tags` (String)
- `job_type` (String)
- `limit` (String)
- `scm_branch` (String)
- `skip_tags` (String)
- `verbosity` (Number)

### Read-Only

- `id` (String) The ID of this resource.
