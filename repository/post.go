package repository

import (
	"context"
	"fmt"
	models "svi-backend/models/posts"
	"time"

	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	CreatePost(ctx context.Context, req models.Post) error
	GetPost(ctx context.Context, fp models.Filter) ([]models.Post, error)
	UpdateById(ctx context.Context, req models.Post) error
	DeleteById(ctx context.Context, req models.Post, id int) error
	FindById(ctx context.Context, req models.Post, id int) (models.Post, error)
	CountPost(ctx context.Context, fp models.Filter) (int, error)
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepositoryInterface {
	return &PostRepositoryImpl{
		db: db,
	}
}

func (p *PostRepositoryImpl) CreatePost(ctx context.Context, req models.Post) error {
	err := p.db.Select("title", "content", "category", "status", "createdAt", "updatedAt").Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostRepositoryImpl) GetPost(ctx context.Context, fp models.Filter) ([]models.Post, error) {
	var posts []models.Post

	if len(fp.Status) > 0 {
		err := p.db.Where("status = ?", fp.Status).
			Select("id", "title", "content", "category", "status").
			Limit(fp.Limit).Offset(fp.Offset).Find(&posts).
			Error
		if err != nil {
			return posts, err
		}
		return posts, nil
	} else {
		err := p.db.Select("id", "title", "content", "category", "status").
			Limit(fp.Limit).Offset(fp.Offset).Find(&posts).
			Error
		if err != nil {
			return posts, err
		}
		return posts, nil
	}

}

func (p *PostRepositoryImpl) UpdateById(ctx context.Context, req models.Post) error {
	data := models.UpdatePostInput{
		Title:     req.Title,
		Content:   req.Content,
		Category:  req.Category,
		UpdatedAt: time.Now(),
		Status:    req.Status,
	}

	err := p.db.Model(&req).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostRepositoryImpl) DeleteById(ctx context.Context, req models.Post, id int) error {
	err := p.db.Model(&req).Where("id = ?", id).Update("status", "trash").Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostRepositoryImpl) FindById(ctx context.Context, req models.Post, id int) (models.Post, error) {
	var post models.Post

	err := p.db.First(&req, "id = ?", id).Scan(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (p *PostRepositoryImpl) CountPost(ctx context.Context, fp models.Filter) (int, error) {
	var count int

	if len(fp.Status) > 0 {
		raw := p.db.Raw("SELECT COUNT(*) FROM posts WHERE status = ?", fp.Status).Scan(&count)
		if raw.Error != nil {
			return count, raw.Error
		}
		fmt.Println("MASUK 1", count)
		return count, nil
	} else {
		raw := p.db.Raw("SELECT COUNT(*) FROM posts").Scan(&count)
		if raw.Error != nil {
			return count, raw.Error
		}
		fmt.Println("MASUK 2", count)
		return count, nil
	}
}
