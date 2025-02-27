---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_team Resource - terraform-provider-awx"
subcategory: ""
description: |-
  
---

# awx_team (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of this Team
- `organization_id` (Number) Numeric ID of the Team organization

### Optional

- `description` (String) Optional description of this Team.
- `role_entitlement` (Block Set) Set of role IDs of the role entitlements (see [below for nested schema](#nestedblock--role_entitlement))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--role_entitlement"></a>
### Nested Schema for `role_entitlement`

Required:

- `role_id` (Number)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `update` (String)
