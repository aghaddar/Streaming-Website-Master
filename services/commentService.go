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

func (s *CommentService) CreateComment(comment *models.Comment) error {
	if comment == nil {
		return errors.New("comment cannot be nil")
	}
	return s.db.Create(comment).Error
}
func (s *CommentService) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := s.db.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (s *CommentService) UpdateComment(comment *models.Comment) error {
	return s.db.Save(comment).Error
}

func (s *CommentService) DeleteComment(id uint) error {
	return s.db.Delete(&models.Comment{}, id).Error
}

func (s *CommentService) ListComments() ([]models.Comment, error) {
	var comments []models.Comment
	if err := s.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
