package controllers

import (
	"go-api-arch-mvc-template/api"
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ハンドラー定義
// ハンドラーはhttpリクエストを受け取ってレスポンスを返す処理
// こいつにモデルの処理（crud）を呼び出すメソッドを実装していく
type AlbumHandler struct{}

// DeleteAlbumId implements api.ServerInterface.
func (a *AlbumHandler) DeleteAlbumId(c *gin.Context, id int) {
	panic("unimplemented")
}

// GetAlbumId implements api.ServerInterface.
func (a *AlbumHandler) GetAlbumId(c *gin.Context, id int) {
	panic("unimplemented")
}

// UpdateAlbumId implements api.ServerInterface.
func (a *AlbumHandler) UpdateAlbumId(c *gin.Context, id int) {
	panic("unimplemented")
}

// gin.Context => reqやresの情報保持
func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody api.CreateAlbumJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	createAlbum, err := models.CreateAlbum(
		requestBody.Title,
		requestBody.ReleaseDate.Time,
		string(requestBody.Category.Name))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createAlbum)
}

func (a *AlbumHandler) GetAlbum(c *gin.Context, ID int) {
	album, err := models.GetAlbum(ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (a *AlbumHandler) UpdateAlbumById(c *gin.Context, ID int) {
	var requestBody api.UpdateAlbumIdJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	album, err := models.GetAlbum(ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	if requestBody.Category != nil {
		// カテゴリ名更新
		album.Category.Name = string(requestBody.Category.Name)
	}
	if requestBody.Title != nil {
		// タイトル更新
		album.Title = *requestBody.Title
	}
	if err := album.Save(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (a *AlbumHandler) DeleteAlbumById(c *gin.Context, ID int) {
	album := models.Album{ID: ID}

	if err := album.Delete(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
