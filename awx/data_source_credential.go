/*
Use this data source to query Credential by ID.

# Example Usage

```hcl
*TBD*
```
*/
package awx

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/robyt96/goawx/client"
	"golang.org/x/exp/slices"
)

func dataSourceCredentialByName() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCredentialByNameRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tower_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceCredentialByNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*awx.AWX)

	name := d.Get("name").(string)
	creds, err := client.CredentialsService.ListCredentials(map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to fetch credentials",
			Detail:   "Unable to fetch credentials from AWX API",
		})
		return diags
	}

	idx := slices.IndexFunc(creds, func(c *awx.Credential) bool { return c.Name == name })

	if idx < 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find credentials",
			Detail:   "Unable to find credentials with given name from AWX API",
		})
		return diags
	}

	cred := creds[idx]

	d.Set("username", cred.Inputs["username"])
	d.Set("kind", cred.Kind)
	d.Set("name", cred.Name)
	d.Set("tower_id", cred.ID)
	d.SetId(strconv.Itoa(cred.ID))
	// d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
