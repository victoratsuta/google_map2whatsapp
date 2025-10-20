package google_maps

import (
	"fmt"
)

type searchPlaceRequest struct {
	location  string
	pageToken string
}

type SearchPlaceRequest interface {
	Location() string
	PageToken() string
	HasPageToken() bool
}

func NewSearchPlaceRequest(location string, pageToken string) (SearchPlaceRequest, error) {
	if location == "" {
		return nil, fmt.Errorf("location cannot be empty")
	}

	return &searchPlaceRequest{
		location:  location,
		pageToken: pageToken,
	}, nil
}

func (c *searchPlaceRequest) Location() string {
	return c.location
}

func (c *searchPlaceRequest) PageToken() string {
	return c.pageToken
}

func (c *searchPlaceRequest) HasPageToken() bool {
	return c.pageToken != ""
}
