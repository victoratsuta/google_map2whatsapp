package entity

import (
	"testing"

	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
)

func TestCompanyValidateNumber(t *testing.T) {
	_, err := entity.NewCompany("Test1", "wefewrgerg", "test1.com")

	expected := "phoneNumber should consist only of integers"
	if expected != err.Error() {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}

func TestCompanyValidateName(t *testing.T) {
	_, err := entity.NewCompany("", "232435234", "test1.com")

	expected := "name cannot be empty"
	if expected != err.Error() {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}

func TestCompanyValidateNumberNotEmpty(t *testing.T) {
	_, err := entity.NewCompany("dfwrgferg", "", "test1.com")

	expected := "phoneNumber cannot be empty"
	if expected != err.Error() {
		t.Errorf("Expected %s, got %s", expected, err.Error())
	}
}
