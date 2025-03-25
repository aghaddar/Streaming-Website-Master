package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (s *CommentService) CreateComment(comment *models.Comments) error {
	if comment == nil {
		return errors.New("comment cannot be nil")
	}
	return s.db.Create(comment).Error
}
func (s *CommentService) GetCommentByID(id uint) (*models.Comments, error) {
	var comment models.Comments
	if err := s.db.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (s *CommentService) UpdateComment(comment *models.Comments) error {
	return s.db.Save(comment).Error
}

func (s *CommentService) DeleteComment(id uint) error {
	return s.db.Delete(&models.Comments{}, id).Error
}

func (s *CommentService) ListComments() ([]models.Comments, error) {
	var comments []models.Comments
	if err := s.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
