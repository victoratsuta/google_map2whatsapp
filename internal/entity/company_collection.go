package entity

type companyCollection struct {
	companies map[string]Company
}

type CompanyCollection interface {
	Add(key string, company Company)
	Get() map[string]Company
	Count() int
}

func (c *companyCollection) Add(key string, company Company) {
	c.companies[key] = company
}

func (c *companyCollection) Get() map[string]Company {
	return c.companies
}

func (c *companyCollection) Count() int {
	return len(c.companies)
}

func NewCompanyCollection() CompanyCollection {
	return &companyCollection{
		companies: map[string]Company{},
	}
}
