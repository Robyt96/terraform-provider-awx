/*
*TBD*

# Example Usage

```hcl

	resource "awx_schedule" "default" {
	  name                      = "schedule-test"
	  rrule                     = "DTSTART;TZID=Europe/Paris:20211214T120000 RRULE:INTERVAL=1;FREQ=DAILY"
	  unified_job_template_id   = awx_job_template.baseconfig.id
	  extra_data                = <<EOL

organization_name: testorg
EOL
}
```
*/
package awx

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/robyt96/goawx/client"
)

func resourceInventorySourceSchedule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventorySourceScheduleCreate,
		ReadContext:   resourceInventorySourceScheduleRead,
		UpdateContext: resourceInventorySourceScheduleUpdate,
		DeleteContext: resourceInventorySourceScheduleDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rrule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"unified_job_template_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"inventory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"extra_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Extra data to be pass for the schedule (YAML format)",
			},
		},
	}
}

func resourceInventorySourceScheduleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesSchedulesService

	result, err := awxService.CreateInventorySourcesSchedule(d.Get("unified_job_template_id").(int), map[string]interface{}{
		"name":        d.Get("name").(string),
		"rrule":       d.Get("rrule").(string),
		"description": d.Get("description").(string),
		"enabled":     d.Get("enabled").(bool),
		"extra_data":  unmarshalYaml(d.Get("extra_data").(string)),
	}, map[string]string{})
	if err != nil {
		log.Printf("Fail to Create Schedule %v", err)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Schedule",
			Detail:   fmt.Sprintf("Schedule failed to create %s", err.Error()),
		})
		return diags
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceInventorySourceScheduleRead(ctx, d, m)
}

func resourceInventorySourceScheduleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.ScheduleService
	id, diags := convertStateIDToNummeric("Update Schedule", d)
	if diags.HasError() {
		return diags
	}

	params := make(map[string]string)
	_, err := awxService.GetByID(id, params)
	if err != nil {
		return buildDiagNotFoundFail("schedule", id, err)
	}

	_, err = awxService.Update(id, map[string]interface{}{
		"name":        d.Get("name").(string),
		"rrule":       d.Get("rrule").(string),
		"description": d.Get("description").(string),
		"enabled":     d.Get("enabled").(bool),
		"extra_data":  unmarshalYaml(d.Get("extra_data").(string)),
	}, map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update Schedule",
			Detail:   fmt.Sprintf("Schedule with name %s failed to update %s", d.Get("name").(string), err.Error()),
		})
		return diags
	}

	return resourceInventorySourceScheduleRead(ctx, d, m)
}

func resourceInventorySourceScheduleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.ScheduleService
	id, diags := convertStateIDToNummeric("Read schedule", d)
	if diags.HasError() {
		return diags
	}

	res, err := awxService.GetByID(id, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail("schedule", id, err)

	}
	d = setScheduleResourceData(d, res)
	return nil
}

func resourceInventorySourceScheduleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.ScheduleService
	id, diags := convertStateIDToNummeric(diagElementHostTitle, d)
	if diags.HasError() {
		return diags
	}

	if _, err := awxService.Delete(id); err != nil {
		return buildDiagDeleteFail(
			diagElementHostTitle,
			fmt.Sprintf("id %v, got %s ",
				id, err.Error()))
	}
	d.SetId("")
	return nil
}
