package services

import (
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
)

type CurrencyService struct {
	Repository *repositories.CurrencyRepository
}

// NewCurrencyService returns new CurrencyService
func NewCurrencyService() *CurrencyService {
	return &CurrencyService{}
}

// Create creates new currency.
func (s *CurrencyService) Create(currency *models.Currency) (*models.Currency, error) {

	return s.Repository.Create(currency)
}

// Update updates currency.
func (s *CurrencyService) Update(currency *models.Currency) (*models.Currency, error) {

	return s.Repository.Update(currency)
}

// Find returns currency and if it does not find the currency, it returns nil.
func (s *CurrencyService) Find(id uint64) (*models.Currency, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of currencies
func (s *CurrencyService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// FindBySymbol returns currency by symbol name.
func (s *CurrencyService) FindBySymbol(symbol string) (*models.Currency, error) {

	return s.Repository.FindBySymbol(symbol)
}