package repo

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/victoratsuta/google_map2whatsapp/internal/repo"
	"github.com/victoratsuta/google_map2whatsapp/pkg/google_maps"
)

func TestGetByLocationSinglePage20Items(t *testing.T) {
	body := generateJson(
		[]map[string]string{
			{"id": "1", "name": "Company", "number": "+790551234567"},
			{"id": "2", "name": "Company 2", "number": "+790551234567"},
			{"id": "3", "name": "Company 3", "number": "+790551234567"},
			{"id": "4", "name": "Company 4", "number": "+790551234567"},
			{"id": "5", "name": "Company 5", "number": "+790551234567"},
			{"id": "6", "name": "Company 6", "number": "+790551234567"},
			{"id": "7", "name": "Company 7", "number": "+790551234567"},
			{"id": "8", "name": "Company 8", "number": "+790551234567"},
			{"id": "9", "name": "Company 9", "number": "+790551234567"},
			{"id": "10", "name": "Company 10", "number": "+790551234567"},
		})

	var apiResponse google_maps.SearchPlaceResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("error parsing JSON: %v", err)
	}

	fc := &fakeClient{
		resps: []google_maps.SearchPlaceResponse{apiResponse},
	}

	r := repo.NewGoogleMapsCompaniesRepository(fc)
	companyCollection, err := r.GetByLocation("Firenze")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 10
	if companyCollection.Count() != expected {
		t.Fatalf("expected %d companies, companyCollection %d", expected, companyCollection.Count())
	}
}

func TestGetByLocationMultiplePagesAndStopBecausePageBecomeTooShort(t *testing.T) {
	bodies := [][]byte{
		generateJson(
			[]map[string]string{
				{"id": "1", "name": "Company", "number": "+790551234567"},
				{"id": "2", "name": "Company 2", "number": "+790551234567"},
				{"id": "3", "name": "Company 3", "number": "+790551234567"},
				{"id": "4", "name": "Company 4", "number": "+790551234567"},
				{"id": "5", "name": "Company 5", "number": "+790551234567"},
				{"id": "6", "name": "Company 6", "number": "+790551234567"},
				{"id": "7", "name": "Company 7", "number": "+790551234567"},
				{"id": "8", "name": "Company 8", "number": "+790551234567"},
				{"id": "9", "name": "Company 9", "number": "+790551234567"},
				{"id": "10", "name": "Company 10", "number": "+790551234567"},
			}),
		generateJson(
			[]map[string]string{
				{"id": "11", "name": "Company", "number": "+790551234567"},
				{"id": "12", "name": "Company 2", "number": "+790551234567"},
				{"id": "13", "name": "Company 3", "number": "+790551234567"},
				{"id": "14", "name": "Company 4", "number": "+790551234567"},
				{"id": "15", "name": "Company 5", "number": "+790551234567"},
				{"id": "16", "name": "Company 6", "number": "+790551234567"},
				{"id": "17", "name": "Company 7", "number": "+790551234567"},
				{"id": "18", "name": "Company 8", "number": "+790551234567"},
				{"id": "19", "name": "Company 9", "number": "+790551234567"},
				{"id": "110", "name": "Company 10", "number": "+790551234567"},
			}),
		generateJson(
			[]map[string]string{
				{"id": "21", "name": "Company", "number": "+790551234567"},
				{"id": "22", "name": "Company 2", "number": "+790551234567"},
			}),
	}

	var apiResponses []google_maps.SearchPlaceResponse
	for _, body := range bodies {
		var apiResponse google_maps.SearchPlaceResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			t.Fatalf("error parsing JSON: %v", err)
		}
		apiResponses = append(apiResponses, apiResponse)
	}

	fc := &fakeClient{resps: apiResponses}

	r := repo.NewGoogleMapsCompaniesRepository(fc)
	companyCollection, err := r.GetByLocation("Firenze")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 22
	if companyCollection.Count() != expected {
		t.Fatalf("expected %d companies, companyCollection %d", expected, companyCollection.Count())
	}
}

func TestGetByLocationMultiplePagesAndStopThereAreTooManyIterations(t *testing.T) {
	bodies := make([][]byte, 15)
	for i := range 15 {
		bodies[i] = generateJson(
			[]map[string]string{
				{"id": fmt.Sprintf("%d", i*10+1), "name": "Company", "number": "+790551234567"},
				{"id": fmt.Sprintf("%d", i*10+2), "name": "Company 2", "number": "+790551234567"},
				{"id": fmt.Sprintf("%d", i*10+3), "name": "Company 3", "number": "+790551234567"},
				{"id": fmt.Sprintf("%d", i*10+4), "name": "Company 4", "number": "+790551234567"},
				{"id": fmt.Sprintf("%d", i*10+5), "name": "Company 5", "number": "+790551234567"},
			},
		)
	}

	var apiResponses []google_maps.SearchPlaceResponse
	for _, body := range bodies {
		var apiResponse google_maps.SearchPlaceResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			t.Fatalf("error parsing JSON: %v", err)
		}
		apiResponses = append(apiResponses, apiResponse)
	}

	fc := &fakeClient{resps: apiResponses}

	r := repo.NewGoogleMapsCompaniesRepository(fc)
	companyCollection, err := r.GetByLocation("Firenze")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 50
	if companyCollection.Count() != expected {
		t.Fatalf("expected %d companies, companyCollection %d", expected, companyCollection.Count())
	}
}

func generateJson(params []map[string]string) []byte {
	var places []string
	for _, param := range params {
		places = append(
			places,
			fmt.Sprintf(`{
   "id":"%s",
   "displayName": {
        "text": "%s",
        "languageCode": "en"
   },
   "internationalPhoneNumber":"%s"
	}`,
				param["id"],
				param["name"],
				param["number"],
			))
	}

	body := []byte(fmt.Sprintf(`{
   "places":[%s],
   "nextPageToken":"ChIJV4-1i2TBhkcRKlzBPwTo4wY"
    }`,
		strings.Join(places, ","),
	))
	return body
}

type fakeClient struct {
	resps []google_maps.SearchPlaceResponse
	index int
}

func (f *fakeClient) SearchPlace(_ google_maps.SearchPlaceRequest) (google_maps.SearchPlaceResponse, error) {
	r := f.resps[f.index]
	if f.index < len(f.resps)-1 {
		f.index++
	}
	return r, nil
}
