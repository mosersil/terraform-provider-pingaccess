package license

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type LicenseAPI interface {
	GetLicenseCommand() (output *models.LicenseView, resp *http.Response, err error)
	ImportLicenseCommand(input *ImportLicenseCommandInput) (output *models.LicenseView, resp *http.Response, err error)
}
