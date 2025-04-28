package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	if db == nil {
		log.Fatal("Database connection is nil in NewAdminService")
	}
	return &AdminService{db: db}
}

// Register creates a new user account
func (s *AdminService) Register(username, email, password, role string) (*models.Admin, error) {
	// Check if user already exists

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	admin := &models.Admin{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
		CreatedAt:    time.Now(),
	}
	if err := s.db.Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
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
