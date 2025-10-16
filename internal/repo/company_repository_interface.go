package repo

import (
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

type CompaniesRepositoryInterface interface {
	GetByLocation(location string) (entity.CompanyCollection, error)
}
