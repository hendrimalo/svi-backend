package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	mPost "svi-backend/models/posts"
	"svi-backend/services"

	"github.com/gin-gonic/gin"
)

type PostControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PostControllerImpl struct {
	pService services.PostServiceInterface
}

func NewPostController(pService services.PostServiceInterface) PostControllerInterface {
	return &PostControllerImpl{
		pService: pService,
	}
}

func (p *PostControllerImpl) Create(c *gin.Context) {
	var req mPost.CreatePostInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := p.pService.Create(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func (p *PostControllerImpl) Get(c *gin.Context) {
	status := c.Query("status")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	filter := mPost.Filter{
		Limit:  intLimit,
		Offset: offset,
		Status: status,
	}

	posts, total, err := p.pService.Get(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"meta": gin.H{
			"total": total,
		},
	})
}

func (p *PostControllerImpl) GetById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param id cannot be empty",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	post, err := p.pService.GetById(c, idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func (post *PostControllerImpl) Update(c *gin.Context) {
	var req mPost.UpdatePostInput

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param id cannot be empty",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = post.pService.Update(c, req, idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("updated post %s success", id),
	})
}

func (post *PostControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param id cannot be empty",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	err = post.pService.Delete(c, idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("deleted post %d success", idInt),
	})
}
