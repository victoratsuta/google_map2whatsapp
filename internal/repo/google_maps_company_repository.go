package repo

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
	"github.com/victoratsuta/google_map2whatsapp/pkg/google_maps"
)

type GoogleMapsCompaniesRepository struct {
	client google_maps.GoogleMapsApiClientInterface
}

func NewGoogleMapsCompaniesRepository(client google_maps.GoogleMapsApiClientInterface) *GoogleMapsCompaniesRepository {
	return &GoogleMapsCompaniesRepository{
		client: client,
	}
}

func (r *GoogleMapsCompaniesRepository) GetByLocation(location string) (entity.CompanyCollection, error) {

	companies := entity.NewCompanyCollection()
	pageToken := ""
	currentIteration := 1
	maxIterations := 10       // this is a safety limit of google maps api call per one search
	minimumAllowedPlaces := 5 // If we got less than 5 places, we can assume that we reached the end of the list

	for {
		request, err := google_maps.NewSearchPlaceRequest(location, pageToken)

		if err != nil {
			return nil, fmt.Errorf("call to google maps: %s", err)
		}

		response, err := r.client.SearchPlace(request)
		if err != nil {
			return nil, fmt.Errorf("call to google maps: %s", err)
		}

		for _, place := range response.Places {

			company, err := entity.NewCompany(
				place.DisplayName.Text,
				strings.TrimPrefix(strings.ReplaceAll(place.InternationalPhoneNumber, " ", ""), "+"),
				"stub",
			)

			if err != nil {
				log.Warn().Err(err).Msg("Error during company creation from google maps response")
				continue
			}
			companies.Add(place.Id, company)
		}

		if len(response.Places) < minimumAllowedPlaces || currentIteration >= maxIterations {
			break
		} else {
			pageToken = response.NextPageToken
			currentIteration++
		}

	}

	return companies, nil
}
