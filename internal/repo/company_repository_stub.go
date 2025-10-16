package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

type GoogleMapsCompaniesRepository struct {
	googleMapsApiKey string
}

type CompaniesRepositoryStub struct{}

func NewGoogleMapsCompaniesRepository(googleMapsApiKey string) *GoogleMapsCompaniesRepository {
	return &GoogleMapsCompaniesRepository{
		googleMapsApiKey: googleMapsApiKey,
	}
}

func NewCompaniesRepositoryStub(googleMapsApiKey string) *CompaniesRepositoryStub {
	return &CompaniesRepositoryStub{}
}

type APIResponse struct {
	Places []struct {
		InternationalPhoneNumber string `json:"internationalPhoneNumber"`
		DisplayName              struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"displayName"`
	} `json:"places"`
}

func (r *GoogleMapsCompaniesRepository) GetByLocation(location string) (entity.CompanyCollection, error) {

	apiResponse, _ := getCompaniesFromGoogleMapsApi(
		location,
		"https://places.googleapis.com/v1/places:searchText",
		r.googleMapsApiKey,
	)

	var companies entity.CompanyCollection
	for _, place := range apiResponse.Places {
		companies.Add(
			entity.Company{
				Name:        place.DisplayName.Text,
				PhoneNumber: strings.TrimPrefix(strings.ReplaceAll(place.InternationalPhoneNumber, " ", ""), "+"),
			})
	}

	return companies, nil
}

func getCompaniesFromGoogleMapsApi(location string, url string, googleMapsApiKey string) (APIResponse, error) {
	requestBody := strings.NewReader(fmt.Sprintf(`{
		"textQuery": "%s"
	}`, location))

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return APIResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", googleMapsApiKey)
	req.Header.Set("X-Goog-FieldMask", "places.displayName,places.internationalPhoneNumber")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return APIResponse{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("error reading response: %v", err)
	}

	// Check for non-200 status
	if resp.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse JSON response
	var apiResponse APIResponse
	if err := json.Unmarshal([]byte(body), &apiResponse); err != nil {
		return APIResponse{}, fmt.Errorf("error parsing JSON: %v", err)
	}
	return apiResponse, nil
}

func (r *CompaniesRepositoryStub) GetByLocation(location string) (entity.CompanyCollection, error) {

	collection := entity.CompanyCollection{}

	collection.Add(entity.Company{"Milan Tour Experts", "77054778117"})
	collection.Add(entity.Company{"Italy Travel Pro", "77054778117"})
	collection.Add(entity.Company{"Lombardy Adventures", "77054778117"})
	collection.Add(entity.Company{"Duomo Tours", "77054778117"})
	collection.Add(entity.Company{"La Scala Experiences", "77054778117"})
	collection.Add(entity.Company{"Navigli Excursions", "77054778117"})
	collection.Add(entity.Company{"Da Vinci Explorers", "77054778117"})
	collection.Add(entity.Company{"Sforza Castle Tours", "77054778117"})
	collection.Add(entity.Company{"Last Supper Visits", "77054778117"})
	collection.Add(entity.Company{"Galleria Guides", "77054778117"})

	return collection, nil
}
