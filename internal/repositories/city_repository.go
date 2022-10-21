package repositories

import (
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"reservation-api/pkg/database/tenant_database_resolver"
)

type CityRepository struct {
	ConnectionResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewCityRepository(connectionResolver *tenant_database_resolver.TenantDatabaseResolver) *CityRepository {
	return &CityRepository{
		ConnectionResolver: connectionResolver,
	}
}

func (r *CityRepository) Create(city *models.City, tenantID uint64) (*models.City, error) {
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Create(&city); tx.Error != nil {
		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Update(city *models.City, tenantID uint64) (*models.City, error) {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Updates(&city); tx.Error != nil {
		return nil, tx.Error
	}

	return city, nil
}

func (r *CityRepository) Delete(id uint64, tenantID uint64) error {

	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if err := db.Model(&models.City{}).Where("id=?", id).Delete(&models.City{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *CityRepository) Find(id uint64, tenantID uint64) (*models.City, error) {

	model := models.City{}
	db := r.ConnectionResolver.GetTenantDB(tenantID)

	if tx := db.Where("id=?", id).Find(&model); tx.Error != nil {

		return nil, tx.Error
	}

	if model.Id == 0 {
		return nil, nil
	}

	return &model, nil
}

func (r *CityRepository) FindAll(input *dto.PaginationFilter) (*commons.PaginatedResult, error) {
	db := r.ConnectionResolver.GetTenantDB(input.TenantID)
	return paginatedList(&models.City{}, db, input)
}
