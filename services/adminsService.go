package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

func (s *AdminService) CreateAdmin(admin *models.Admin) error {
	return s.db.Create(admin).Error
}

func (s *AdminService) GetAdminByID(id uint) (*models.Admin, error) {
	var admin models.Admin
	if err := s.db.First(&admin, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func (s *AdminService) UpdateAdmin(admin *models.Admin) error {
	return s.db.Save(admin).Error
}

func (s *AdminService) DeleteAdmin(id uint) error {
	return s.db.Delete(&models.Admin{}, id).Error
}

func (s *AdminService) ListAdmins() ([]models.Admin, error) {
	var admins []models.Admin
	if err := s.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}
