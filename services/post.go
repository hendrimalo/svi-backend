package services

import (
	"context"
	"errors"
	"fmt"
	models "svi-backend/models/posts"
	"svi-backend/repository"
	"time"

	"gorm.io/gorm"
)

type PostServiceImpl struct {
	pRepo repository.PostRepositoryInterface
}

type PostServiceInterface interface {
	Create(ctx context.Context, req models.CreatePostInput) error
	GetById(ctx context.Context, id int) (models.Post, error)
	Get(ctx context.Context, fp models.Filter) ([]models.Post, int, error)
	Update(ctx context.Context, req models.UpdatePostInput, id int) error
	Delete(ctx context.Context, id int) error
}

func NewPostService(db *gorm.DB, pRepo repository.PostRepositoryInterface) PostServiceInterface {
	return &PostServiceImpl{
		pRepo: pRepo,
	}
}

func (p *PostServiceImpl) Create(ctx context.Context, req models.CreatePostInput) error {
	ts := time.Now()
	post := models.Post{
		Title:     req.Title,
		Content:   req.Content,
		Category:  req.Category,
		CreatedAt: &ts,
		UpdatedAt: &ts,
		Status:    req.Status,
	}

	err := p.pRepo.CreatePost(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostServiceImpl) Get(ctx context.Context, fp models.Filter) ([]models.Post, int, error) {
	var posts []models.Post
	var total int

	posts, err := p.pRepo.GetPost(ctx, fp)
	if err != nil {
		return posts, total, err
	}
	total, err = p.pRepo.CountPost(ctx, fp)
	return posts, total, nil
}

func (p *PostServiceImpl) GetById(ctx context.Context, id int) (models.Post, error) {
	var post models.Post

	post, err := p.pRepo.FindById(ctx, post, id)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (p *PostServiceImpl) Update(ctx context.Context, req models.UpdatePostInput, id int) error {
	ts := time.Now()
	post := models.Post{
		Title:     req.Title,
		Content:   req.Content,
		Category:  req.Category,
		UpdatedAt: &ts,
		Status:    req.Status,
	}

	_, err := p.pRepo.FindById(ctx, post, id)
	if err != nil {
		return err
	}

	post.Id = id
	err = p.pRepo.UpdateById(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostServiceImpl) Delete(ctx context.Context, id int) error {
	var post models.Post

	v, err := p.pRepo.FindById(ctx, post, id)
	if err != nil {
		return err
	}

	if v.Status == "trash" {
		msg := fmt.Sprintf("error: post id %d has been deleted", id)
		return errors.New(msg)
	}

	err = p.pRepo.DeleteById(ctx, post, id)
	if err != nil {
		return err
	}
	return nil
}
