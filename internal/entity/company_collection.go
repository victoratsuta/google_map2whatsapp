package entity

type Company struct {
	Name        string
	PhoneNumber string
}

type CompanyCollection struct {
	Companies []Company
}

func (c *CompanyCollection) Add(company Company) {
	c.Companies = append(c.Companies, company)
}

func (c *CompanyCollection) Get() []Company {
	return c.Companies
}

func (c *CompanyCollection) Count() int {
	return len(c.Companies)
}

func NewCompanyCollection() *CompanyCollection {
	return &CompanyCollection{
		Companies: make([]Company, 0),
	}
}
