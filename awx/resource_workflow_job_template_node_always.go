/*
*TBD*

# Example Usage

```hcl
resource "random_uuid" "workflow_node_k3s_uuid" {}

	resource "awx_workflow_job_template_node_always" "k3s" {
	  workflow_job_template_id      = awx_workflow_job_template.default.id
	  workflow_job_template_node_id = awx_workflow_job_template_node.default.id
	  unified_job_template_id       = awx_job_template.k3s.id
	  inventory_id                  = awx_inventory.default.id
	  identifier                    = random_uuid.workflow_node_k3s_uuid.result
	}

```
*/
package awx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/robyt96/goawx/client"
)

func resourceWorkflowJobTemplateNodeAlways() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowJobTemplateNodeAlwaysCreate,
		ReadContext:   resourceWorkflowJobTemplateNodeRead,
		UpdateContext: resourceWorkflowJobTemplateNodeUpdate,
		DeleteContext: resourceWorkflowJobTemplateNodeDelete,
		Schema:        workflowJobNodeSchema,
	}
}
func resourceWorkflowJobTemplateNodeAlwaysCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.WorkflowJobTemplateNodeAlwaysService
	return createNodeForWorkflowJob(awxService, ctx, d, m)
}
