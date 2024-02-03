package db

import (
	"backend/dto"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// @Summary Post
// @Description Post
// @ID post
// @Accept  json
// @Produce  json
// @Param username path dto.Post true "Post"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/post [post]
func Post(c echo.Context) error {
	cn := Conn()
	rq := new(dto.Post)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)

	p := models.Post{OwnerID: claims.UserId, Content: rq.Content, Attach: rq.Attach}
	cn.Create(&p)

	return c.JSON(http.StatusOK, models.Response(nil, "Successfully Posted"))
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
