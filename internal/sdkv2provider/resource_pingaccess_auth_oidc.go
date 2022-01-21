package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePingAccessAuthOidc() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingAccessAuthOidcCreate,
		ReadContext:   resourcePingAccessAuthOidcRead,
		UpdateContext: resourcePingAccessAuthOidcUpdate,
		DeleteContext: resourcePingAccessAuthOidcDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingAccessAuthOidcSchema(),
		Description: `Manages the PingAccess Admin UI authentication using OIDC.

-> This resource manages a singleton within PingAccess and as such you should ONLY ever declare one of this resource type. Deleting this resource resets the Auth Token Management configuration to default values.`,
	}
}

func resourcePingAccessAuthOidcSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authn_req_list_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     "ON",
			Description: "The ID of the authentioation requirement list for Administrative SSO login to PingAccess",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "",
		},
		"use_slo": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "",
		},
		"username_attribute_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "sub",
			Description: "",
		},
		"oidc_configuration": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"cache_user_attributes": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "",
					},
					"client_credentials": {
						Type:        schema.TypeList,
						Required:    true,
						MaxItems:    1,
						Description: "Specify the client credentials.",
						Elem:        oAuthClientCredentialsResource(),
					},
					"enable_refresh_user": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "",
					},
					"oidc_login_type": {
						Type:             schema.TypeString,
						Optional:         true,
						Default:          "Code",
						ValidateDiagFunc: validateOidcLoginType,
						Description:      "",
					},
					"pf_session_state_cache_in_seconds": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     60,
						Description: "",
					},
					"pkce_challenge_type": {
						Type:             schema.TypeString,
						Optional:         true,
						Default:          "OFF",
						ValidateDiagFunc: validatePkceChallengeType,
						Description:      "Specify the code_challenge_method to use for PKCE during the Code login flow. OFF signifies to not use PKCE.",
					},
					"refresh_user_info_claims_interval": {
						Type:        schema.TypeInt,
						Optional:    true,
						Default:     60,
						Description: "",
					},
					"scopes": {
						Type:     schema.TypeSet,
						Required: true,
						DefaultFunc: func() (interface{}, error) {
							return []interface{}{"profile"}, nil
						},
						Description: "The list of scopes to be specified in the access request. If not specified, the default scopes ('profile') will be used. The openid scope is implied and does not need to be specified in this list.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"send_request_url_to_provider": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     true,
						Description: "",
					},
					"validate_session_is_alive": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "",
					},
				},
			},
		},
	}
}

func resourcePingAccessAuthOidcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Auth
	_, err := svc.DeleteOidcAuthCommand()
	if err != nil {
		return diag.Errorf("unable to reset OidcAuth: %s", err)
	}
	return nil
}

func resourcePingAccessAuthOidcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Auth
	result, _, err := svc.GetOidcAuthCommand()
	if err != nil {
		return diag.Errorf("unable to read OidcAuth: %s", err)
	}

	return resourcePingAccessOAuthOidcReadResult(d, result)
}

func resourcePingAccessOAuthOidcReadResult(d *schema.ResourceData, input *models.OidcConfigView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataIntWithDiagnostic(d, "authn_req_list_id", input.AuthnReqListId, &diags)
	setResourceDataBoolWithDiagnostic(d, "enabled", input.Enabled, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_slo", input.UseSlo, &diags)
	setResourceDataStringWithDiagnostic(d, "username_attribute_name", input.UsernameAttributeName, &diags)

	return diags
}

func resourcePingAccessAuthOidcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("auth_oidc")
	return resourcePingAccessAuthOidcUpdate(ctx, d, m)
}

func resourcePingAccessAuthOidcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(paClient).Auth
	input := auth.UpdateOidcAuthCommandInput{
		Body: *resourcePingAccessAuthOidcReadData(d),
	}
	result, _, err := svc.UpdateOidcAuthCommand(&input)
	if err != nil {
		return diag.Errorf("unable to update OIDC Authentication: %s", err)
	}

	d.SetId("auth_oidc")
	return resourcePingAccessOAuthOidcReadResult(d, result)
}

func resourcePingAccessAuthOidcReadData(d *schema.ResourceData) *models.OidcConfigView {

	//scopes := expandStringList(d.Get("scopes").(*schema.Set).List())

	oidc := &models.OidcConfigView{
		AuthnReqListId:        Int(d.Get("authn_req_list_id").(int)),
		Enabled:               Bool(d.Get("enabled").(bool)),
		UseSlo:                Bool(d.Get("use_slo").(bool)),
		UsernameAttributeName: String(d.Get("username_attribute_name").(string)),
		OidcConfiguration:     expandAdminWebSessionOidcConfigurationView(d.Get("oidc_configuration").([]interface{})),
	}
	return oidc
}
