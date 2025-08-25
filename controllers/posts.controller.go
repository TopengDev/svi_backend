package controllers

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/topengdev/svi_backend/controllers/validatos"
	"github.com/topengdev/svi_backend/initializers"
	"github.com/topengdev/svi_backend/interfaces"
	"github.com/topengdev/svi_backend/models"
	"gorm.io/gorm"
)

// POST ~ /article/
func PostsCreate(c *gin.Context){
	
	// Binds payload
	var payload interfaces.ICreatePostDTO 
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validates the payload
	if valMsg, isErr := validators.ValidateCreatePost(payload); isErr {
		c.JSON(http.StatusBadRequest, gin.H{"error": valMsg})
		return
	}

	// Create the new post entry
	newPost := models.Post{
		Title: payload.Title,
		Content: payload.Content,
		Category: payload.Category,
		Status: payload.Status,
		Created_date:time.Now().Format("2006-01-02"),
	}


	// Insert the new post entry to db
	if err := initializers.DB.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// POST/PUT/PATCH ~ /article/:id
func PostUpdate(c *gin.Context) {
	// Parse id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	// Binds payload
	var payload interfaces.IUpdatePostDTO
	payload.Id = int(id)
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validates payload
	if valMsg, isErr := validators.ValidateUpdatePost(payload); isErr {
		c.JSON(http.StatusBadRequest, gin.H{"error": valMsg})
		return
	}

	// Find existing entry
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch article"})
		return
	}

	// Update the fields if provided
	if payload.Title != "" {
		post.Title = payload.Title
	}
	if payload.Content != "" {
		post.Content = payload.Content
	}
	if payload.Category != "" {
		post.Category = payload.Category
	}
	if payload.Status != "" {
		post.Status = payload.Status
	}
	post.Updated_date = time.Now().Format("2006-01-02")

	// Update the database entry
	if err := initializers.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DELETE ~ /article/:id
func PostDelete(c *gin.Context) {
	// parse id
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	// Soft delete the database entry
	res := initializers.DB.Delete(&models.Post{}, id)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete article"})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GET ~ /article/:id
func PostGetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch article"})
		return
	}

	// reshape response
	
	c.JSON(http.StatusOK, gin.H{
		"Title": post.Title,
		"Content": post.Content,
		"Category": post.Category,
		"Status": post.Status,
	})
}


// GET ~ /article/:limit/:offset
func PostsList(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a positive integer"})
		return
	}
	// Max limit
	if limit > 100 {
		limit = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offset must be a non-negative integer"})
		return
	}

	var posts []models.Post
	if err := initializers.DB.
		Limit(limit).
		Offset(offset).
		Order("id DESC").
		Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
		return
	}

	// total count
	var total int64
	if err := initializers.DB.Model(&models.Post{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count posts"})
		return
	}

	// reshape response
	type PostResp struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Category string `json:"category"`
		Status   string `json:"status"`
	}
	resp := make([]PostResp, len(posts))
	for i, p := range posts {
		resp[i] = PostResp{
			Title:    p.Title,
			Content:  p.Content,
			Category: p.Category,
			Status:   p.Status,
		}
	}

	c.JSON(http.StatusOK, resp)
}
