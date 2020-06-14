package pingaccess

import (
	"context"
	"fmt"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessAuthTokenManagement() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAuthTokenManagementCreate,
		ReadContext:   resourcePingAccessAuthTokenManagementRead,
		UpdateContext: resourcePingAccessAuthTokenManagementUpdate,
		DeleteContext: resourcePingAccessAuthTokenManagementDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingAccessAuthTokenManagementSchema(),
	}
}

func resourcePingAccessAuthTokenManagementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"issuer": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "PingAccessAuthToken",
			Description: "The issuer value to include in auth tokens. PingAccess inserts this value as the iss claim within the auth tokens.",
		},
		"key_roll_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "This field is true if key rollover is enabled. When false, PingAccess will not rollover keys at the configured interval.",
		},
		"key_roll_period_in_hours": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     24,
			Description: "The interval (in hours) at which PingAccess will roll the keys. Key rollover updates keys at regular intervals to ensure the security of signed auth tokens.",
		},
		"signing_algorithm": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "P-256",
			Description: "The signing algorithm used when creating signed auth tokens.",
		},
	}
}

func resourcePingAccessAuthTokenManagementCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("auth_token_management")
	return resourcePingAccessAuthTokenManagementUpdate(ctx, d, m)
}

func resourcePingAccessAuthTokenManagementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).AuthTokenManagement
	result, _, err := svc.GetAuthTokenManagementCommand()
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read AuthTokenManagement: %s", err))
	}

	return resourcePingAccessAuthTokenManagementReadResult(d, result)
}

func resourcePingAccessAuthTokenManagementUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).AuthTokenManagement
	input := pa.UpdateAuthTokenManagementCommandInput{
		Body: *resourcePingAccessAuthTokenManagementReadData(d),
	}
	result, _, err := svc.UpdateAuthTokenManagementCommand(&input)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to update AuthTokenManagement: %s", err))
	}

	d.SetId("auth_token_management")
	return resourcePingAccessAuthTokenManagementReadResult(d, result)
}

func resourcePingAccessAuthTokenManagementDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(*pa.Client).AuthTokenManagement
	_, err := svc.DeleteAuthTokenManagementCommand()
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to delete AuthTokenManagement: %s", err))

	}
	return nil
}

func resourcePingAccessAuthTokenManagementReadResult(d *schema.ResourceData, input *pa.AuthTokenManagementView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "issuer", input.Issuer, &diags)
	setResourceDataBoolWithDiagnostic(d, "key_roll_enabled", input.KeyRollEnabled, &diags)
	setResourceDataIntWithDiagnostic(d, "key_roll_period_in_hours", input.KeyRollPeriodInHours, &diags)
	setResourceDataStringWithDiagnostic(d, "signing_algorithm", input.SigningAlgorithm, &diags)
	return diags
}

func resourcePingAccessAuthTokenManagementReadData(d *schema.ResourceData) *pa.AuthTokenManagementView {
	atm := &pa.AuthTokenManagementView{
		Issuer:               String(d.Get("issuer").(string)),
		KeyRollEnabled:       Bool(d.Get("key_roll_enabled").(bool)),
		KeyRollPeriodInHours: Int(d.Get("key_roll_period_in_hours").(int)),
		SigningAlgorithm:     String(d.Get("signing_algorithm").(string)),
	}

	return atm
}
