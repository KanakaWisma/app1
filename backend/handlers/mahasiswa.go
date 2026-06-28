package handlers

import (
	"kampus-api/config"
	"kampus-api/models"

	"github.com/gin-gonic/gin"
)

func GetMahasiswa(c *gin.Context) {

	var mahasiswa []models.Mahasiswa

	result := config.DB.Find(&mahasiswa)

	c.JSON(200, gin.H{
		"data":  mahasiswa,
		"rows":  result.RowsAffected,
		"error": result.Error,
	})
}

func CreateMahasiswa(c *gin.Context) {

	var mahasiswa models.Mahasiswa

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&mahasiswa)

	c.JSON(201, mahasiswa)

}

func UpdateMahasiswa(c *gin.Context) {

	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	if err := config.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Mahasiswa tidak ditemukan",
		})
		return
	}

	var input models.Mahasiswa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Model(&mahasiswa).Updates(input)

	c.JSON(200, mahasiswa)

}

func DeleteMahasiswa(c *gin.Context) {

	id := c.Param("id")

	var mahasiswa models.Mahasiswa

	if err := config.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Mahasiswa tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&mahasiswa)

	c.JSON(200, gin.H{
		"message": "Data berhasil dihapus",
	})

}

func SearchMahasiswa(c *gin.Context) {

	search := c.Query("search")

	var mahasiswa []models.Mahasiswa

	config.DB.Where(
		"nama ILIKE ?",
		"%"+search+"%",
	).Find(&mahasiswa)

	c.JSON(200, mahasiswa)

}
