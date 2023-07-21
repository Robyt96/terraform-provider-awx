/*
*TBD*

# Example Usage

```hcl

	resource "awx_project" "myproj" {
	  name = "My AWX Project"
	  ...
	}

	data "awx_project_role" "proj_admins" {
	  name       = "Admin"
	  project_id = resource.awx_project.myproj.id
	}

```
*/
package awx

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/robyt96/goawx/client"
)

func dataSourceProjectRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRolesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceProjectRolesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	params := make(map[string]string)

	proj_id := d.Get("project_id").(int)

	Project, err := client.ProjectService.GetProjectByID(proj_id, params)
	if err != nil {
		return buildDiagnosticsMessage(
			"Get: Fail to fetch Project",
			"Fail to find the project got: %s",
			err.Error(),
		)
	}

	roleslist := []*awx.ApplyRole{
		Project.SummaryFields.ObjectRoles.UseRole,
		Project.SummaryFields.ObjectRoles.AdminRole,
		Project.SummaryFields.ObjectRoles.UpdateRole,
		Project.SummaryFields.ObjectRoles.ReadRole,
	}

	if roleID, okID := d.GetOk("id"); okID {
		id := roleID.(int)
		for _, v := range roleslist {
			if v != nil && id == v.ID {
				d = setProjectRoleData(d, v)
				return diags
			}
		}
	}

	if roleName, okName := d.GetOk("name"); okName {
		name := roleName.(string)

		for _, v := range roleslist {
			if v != nil && name == v.Name {
				d = setProjectRoleData(d, v)
				return diags
			}
		}
	}

	return buildDiagnosticsMessage(
		"Failed to fetch project role - Not Found",
		"The project role was not found",
	)
}

func setProjectRoleData(d *schema.ResourceData, r *awx.ApplyRole) *schema.ResourceData {
	d.Set("name", r.Name)
	d.SetId(strconv.Itoa(r.ID))
	return d
}
