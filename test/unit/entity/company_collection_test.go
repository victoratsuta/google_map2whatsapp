package entity

import (
	"testing"

	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

func TestCompanyCollectionCount(t *testing.T) {
	collection := entity.NewCompanyCollection()

	// Create test companies
	company1, _ := entity.NewCompany("Test1", "123", "test1.com")
	company2, _ := entity.NewCompany("Test2", "456", "test2.com")
	company3, _ := entity.NewCompany("Test3", "789", "test3.com")

	collection.Add("1", company1)
	collection.Add("2", company2)
	collection.Add("3", company3)

	expected := 3
	if expected != collection.Count() {
		t.Errorf("Expected %d companies, got %d", expected, collection.Count())
	}
}
