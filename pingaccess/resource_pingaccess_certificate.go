package pingaccess

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func resourcePingAccessCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingAccessCertificateCreate,
		Read:   resourcePingAccessCertificateRead,
		Update: resourcePingAccessCertificateUpdate,
		Delete: resourcePingAccessCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePingAccessCertificateCreate(d *schema.ResourceData, m interface{}) error {
	input := pa.ImportTrustedCertInput{
		Body: pa.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
	}

	svc := m.(*pa.Client).Certificates

	result, _, err := svc.ImportTrustedCert(&input)
	if err != nil {
		return fmt.Errorf("Error creating Certificate: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Certificates
	input := &pa.GetTrustedCertInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTrustedCert(input)
	if err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	input := pa.UpdateTrustedCertInput{
		Body: pa.X509FileImportDocView{
			Alias:    String(d.Get("alias").(string)),
			FileData: String(d.Get("file_data").(string)),
		},
		Id: d.Id(),
	}

	svc := m.(*pa.Client).Certificates

	result, _, err := svc.UpdateTrustedCert(&input)
	if err != nil {
		return fmt.Errorf("Error creating Certificate: %s", err)
	}

	d.SetId(result.Id.String())
	return resourcePingAccessCertificateReadResult(d, result)
}

func resourcePingAccessCertificateDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pa.Client).Certificates
	_, err := svc.DeleteTrustedCertCommand(&pa.DeleteTrustedCertCommandInput{Id: d.Id()})
	if err != nil {
		return fmt.Errorf("Error deleting virtualhost: %s", err)
	}
	return nil
}

func resourcePingAccessCertificateReadResult(d *schema.ResourceData, rv *pa.TrustedCertView) error {
	if err := d.Set("alias", rv.Alias); err != nil {
		return err
	}
	return nil
}
