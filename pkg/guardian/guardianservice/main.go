package guardianservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/guardian/guardianhandler"
	"github.com/skyerus/history-api/pkg/models"
)

const searchEndpoint string = "/search"

// Search ...
func Search(rawQuery string) (*models.GuardianContent, customerror.Error) {
	var guardianContent models.GuardianContent
	request, err := http.NewRequest("GET", os.Getenv("GUARDIAN_BASE_URL")+searchEndpoint, nil)
	if err != nil {
		return nil, customerror.NewGenericHTTPError(err)
	}
	request.URL.RawQuery = rawQuery
	handler := guardianhandler.NewGuardianHandler()
	response, customErr := handler.SendRequest(request)
	if customErr != nil {
		return nil, customErr
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, customerror.NewGenericHTTPError(err)
	}
	err = json.Unmarshal(body, &guardianContent)
	if err != nil {
		return nil, customerror.NewGenericHTTPError(err)
	}

	return &guardianContent, nil
}
