package entity

import (
	"errors"
	"strconv"
)

type company struct {
	name           string
	phoneNumber    string
	googleMapsLink string
}

type Company interface {
	Name() string
	PhoneNumber() string
	GoogleMapsLink() string
}

func NewCompany(name, phoneNumber, googleMapsLink string) (Company, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if phoneNumber == "" {
		return nil, errors.New("phoneNumber cannot be empty")
	}

	if _, err := strconv.Atoi(phoneNumber); err != nil {
		return nil, errors.New("phoneNumber should consist only of integers")
	}
	return &company{
		name:           name,
		phoneNumber:    phoneNumber,
		googleMapsLink: googleMapsLink,
	}, nil
}

func (c *company) Name() string {
	return c.name
}

func (c *company) PhoneNumber() string {
	return c.phoneNumber
}

func (c *company) GoogleMapsLink() string {
	return c.googleMapsLink
}
