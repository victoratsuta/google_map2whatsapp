package repo

import (
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

type CompaniesRepositoryStub struct{}

func NewCompaniesRepositoryStub() *CompaniesRepositoryStub {
	return &CompaniesRepositoryStub{}
}

func (*CompaniesRepositoryStub) GetByLocation(_ string) (entity.CompanyCollection, error) {
	collection := entity.NewCompanyCollection()

	company, _ := entity.NewCompany("Milan Tour Experts", "77054778117", "test1.com")
	collection.Add("1", company)
	collection.Add("2", company)

	return collection, nil
}
