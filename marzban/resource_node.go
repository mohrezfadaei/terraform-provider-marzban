package marzban

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNodeCreate,
		ReadContext:   resourceNodeRead,
		UpdateContext: resourceNodeUpdate,
		DeleteContext: resourceNodeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  62050,
			},
			"api_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  62051,
			},
			"usage_coefficient": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  1.0,
			},
			"add_as_new_host": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"node_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	name := d.Get("name").(string)
	address := d.Get("address").(string)
	port := d.Get("port").(int)
	apiPort := d.Get("api_port").(int)
	usageCoefficient := d.Get("usage_coefficient").(float64)
	addAsNewHost := d.Get("add_as_new_host").(bool)

	nodeID, err := client.CreateNode(name, address, port, apiPort, usageCoefficient, addAsNewHost)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", nodeID))
	return resourceNodeRead(ctx, d, m)
}

func resourceNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	NodeID := d.Id()

	node, err := client.GetNode(NodeID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", node.Name)
	d.Set("address", node.Address)
	d.Set("port", node.Port)
	d.Set("api_port", node.ApiPort)
	d.Set("usage_coefficient", node.UsageCoefficient)
	d.Set("add_as_new_host", node.AddAsNewHost)

	return nil
}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	nodeID := d.Id()

	if d.HasChanges("name", "address", "port", "api_port", "usage_coefficient", "add_as_new_host") {
		name := d.Get("name").(string)
		address := d.Get("address").(string)
		port := d.Get("port").(int)
		apiPort := d.Get("api_port").(int)
		usageCoefficient := d.Get("usage_coefficient").(float64)
		addAsNewHost := d.Get("add_as_new_host").(bool)

		err := client.UpdateNode(nodeID, name, address, port, apiPort, usageCoefficient, addAsNewHost)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceNodeRead(ctx, d, m)
}

func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	nodeID := d.Id()

	err := client.DeleteNode(nodeID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
