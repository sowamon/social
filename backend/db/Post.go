package db

import (
	"backend/models"
)

// @Summary Post
// @Description Post
// @ID post
// @Accept  json
// @Produce  json
// @Param username path post.PostDTO true "Post"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/post [post]
func Post(ownerId uint, content string, attach string) (models.IResponse, int) {
	cn := Conn()

	p := models.Post{OwnerID: ownerId, Content: content, Attach: attach}
	cn.Create(&p)

	return models.Response(nil, "Successfully Posted"), 200
}

// @Summary Get Post
// @Description Get Posts
// @ID getPost
// @Accept  json
// @Produce  json
// @Param username path post.GetPostDTO true "Post"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/post [get]
func AuthGetPosts(cursor int) (models.IResponse, int) {
	cn := Conn()

	var m []models.Post

	err := cn.Debug().Limit(20).Order("id DESC").Find(&m, "id <= ?", cursor).Error
	if err == nil && len(m) != 0 {
		return models.Response(m, "success"), 200
	} else {
		return models.Response(nil, "no message"), 400
	}
}
