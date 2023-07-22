/*
Use this data source to query Credential Type by ID.

# Example Usage

```hcl
*TBD*
```
*/
package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/robyt96/goawx/client"
	"golang.org/x/exp/slices"
)

func dataSourceCredentialTypeByName() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCredentialTypeByNameRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inputs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"injectors": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCredentialTypeByNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*awx.AWX)
	name := d.Get("name").(string)
	credTypes, _, err := client.CredentialTypeService.ListCredentialTypes(map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find credential types",
			Detail:   fmt.Sprintf("Unable to fetch credential types. Error: %s", err.Error()),
		})
	}

	idx := slices.IndexFunc(credTypes, func(c *awx.CredentialType) bool { return c.Name == name })

	if idx < 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to fetch credential type",
			Detail:   fmt.Sprintf("Unable to fetch credential type with name: %s. Not found", name),
		})
		return diags
	}

	credType := credTypes[idx]

	d.Set("name", credType.Name)
	d.Set("description", credType.Description)
	d.Set("kind", credType.Kind)
	d.Set("inputs", credType.Inputs)
	d.Set("injectors", credType.Injectors)
	d.SetId(strconv.Itoa(credType.ID))

	return diags
}
