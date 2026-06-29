package handlers

import (
	"kampus-api/config"
	"kampus-api/models"

	"github.com/gin-gonic/gin"
)

// GetJurusan godoc
// @Summary Ambil semua jurusan
// @Description Menampilkan seluruh data jurusan
// @Tags Jurusan
// @Produce json
// @Success 200 {array} models.Jurusan
// @Router /jurusan [get]
func GetJurusan(c *gin.Context) {

	var jurusan []models.Jurusan

	config.DB.Find(&jurusan)

	c.JSON(200, gin.H{
		"data": jurusan,
	})
}