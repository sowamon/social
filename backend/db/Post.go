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
func Get(cursor int) (models.IResponse, int) {
	cn := Conn()

	var m []models.Post

	cn = cn.Limit(20).Order("id DESC")

	if cursor != 0 {
		cn = cn.Where("id < ?", cursor)
	}

	err := cn.Find(&m).Error

	if err != nil {
		return models.Response(nil, err.Error()), 400
	}
	return models.Response(m, "success"), 200
}
