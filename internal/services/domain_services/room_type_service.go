package domain_services

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/internal/repositories"
)

type RoomTypeService struct {
	Repository *repositories.RoomTypeRepository
}

func NewRoomTypeService() *RoomTypeService {
	return &RoomTypeService{}
}

// Create creates new RoomType.
func (s *RoomTypeService) Create(roomType *models.RoomType) (*models.RoomType, error) {

	return s.Repository.Create(roomType)
}

// Update updates RoomType.
func (s *RoomTypeService) Update(roomType *models.RoomType) (*models.RoomType, error) {

	return s.Repository.Update(roomType)
}

// Find returns RoomType and if it does not find the RoomType, it returns nil.
func (s *RoomTypeService) Find(id uint64) (*models.RoomType, error) {

	return s.Repository.Find(id)
}

// FindAll returns paginates list of hotel types.
func (s *RoomTypeService) FindAll(input *dto.PaginationInput) (*commons.PaginatedList, error) {

	return s.Repository.FindAll(input)
}

// Delete removes hotel type by given id.
func (s *RoomTypeService) Delete(id uint64) error {

	return s.Repository.Delete(id)
}