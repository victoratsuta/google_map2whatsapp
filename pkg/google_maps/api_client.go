package google_maps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GoogleMapsApiClientInterface interface {
	SearchPlace(searchPlace SearchPlaceRequest) (SearchPlaceResponse, error)
}

type GoogleMapsHttpApiClient struct {
	googleMapsApiKey string
	baseUrl          string
}

func NewGoogleMapsHttpApiClient(googleMapsApiKey, baseUrl string) *GoogleMapsHttpApiClient {
	return &GoogleMapsHttpApiClient{
		googleMapsApiKey: googleMapsApiKey,
		baseUrl:          baseUrl,
	}
}

func (client *GoogleMapsHttpApiClient) SearchPlace(searchPlace SearchPlaceRequest) (SearchPlaceResponse, error) {
	const itemsPerPage = 20
	var fieldMask = []string{
		"places.id",
		"places.displayName",
		"places.internationalPhoneNumber",
		"nextPageToken",
	}
	requestBody := strings.NewReader(fmt.Sprintf(`{
		"textQuery": "%s",
		"pageToken": "%s",
		"pageSize" : "%d",
	}`, searchPlace.Location(), searchPlace.PageToken(), itemsPerPage))

	req, err := http.NewRequestWithContext(context.TODO(), "POST", client.baseUrl, requestBody)
	if err != nil {
		return SearchPlaceResponse{}, fmt.Errorf("error creating request to google maps search: %w", err)
	}

	req.Header.Set("X-Goog-Api-Key", client.googleMapsApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-FieldMask", strings.Join(fieldMask, ","))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return SearchPlaceResponse{}, fmt.Errorf("error making request to google maps search: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchPlaceResponse{}, fmt.Errorf("error reading response from to google maps search: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return SearchPlaceResponse{}, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse SearchPlaceResponse
	if err = json.Unmarshal(body, &apiResponse); err != nil {
		return SearchPlaceResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return SearchPlaceResponse{}, err
	}

	return apiResponse, nil
}
