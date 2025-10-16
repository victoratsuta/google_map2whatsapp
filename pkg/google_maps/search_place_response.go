package google_maps

type SearchPlaceResponse struct {
	Places []struct {
		InternationalPhoneNumber string `json:"internationalPhoneNumber"`
		Id                       string `json:"id"`
		DisplayName              struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"displayName"`
	} `json:"places"`
	NextPageToken string `json:"nextPageToken"`
}
