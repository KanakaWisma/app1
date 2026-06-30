package handlers

import (
	"kampus-api/config"
	"kampus-api/models"

	"github.com/gin-gonic/gin"
)

// GetMahasiswa godoc
// @Summary Ambil semua mahasiswa
// @Description Menampilkan seluruh data mahasiswa
// @Tags Mahasiswa
// @Produce json
// @Success 200 {array} models.Mahasiswa
// @Router /mahasiswa [get]
func GetMahasiswa(c *gin.Context) {

	var mahasiswa []models.Mahasiswa

	result := config.DB.Preload("Jurusan").Find(&mahasiswa)

	c.JSON(200, gin.H{
		"data":  mahasiswa,
		"rows":  result.RowsAffected,
		"error": result.Error,
	})
}

// CreateMahasiswa godoc
// @Summary Tambah mahasiswa
// @Description Menambahkan data mahasiswa baru
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param mahasiswa body models.Mahasiswa true "Data Mahasiswa"
// @Success 201 {object} models.Mahasiswa
// @Router /mahasiswa [post]
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

// UpdateMahasiswa godoc
// @Summary Update mahasiswa
// @Description Mengubah data mahasiswa berdasarkan ID
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path int true "ID Mahasiswa"
// @Param mahasiswa body models.Mahasiswa true "Data Mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [put]
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

	c.JSON(200, gin.H{
		"message": "Data berhasil diupdate",
	})
}

// DeleteMahasiswa godoc
// @Summary Hapus mahasiswa
// @Description Menghapus data mahasiswa berdasarkan ID
// @Tags Mahasiswa
// @Produce json
// @Param id path int true "ID Mahasiswa"
// @Success 200 {object} map[string]interface{}
// @Router /mahasiswa/{id} [delete]
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

// SearchMahasiswa godoc
// @Summary Cari mahasiswa
// @Description Mencari mahasiswa berdasarkan nama
// @Tags Mahasiswa
// @Produce json
// @Param search query string true "Nama Mahasiswa"
// @Success 200 {array} models.Mahasiswa
// @Router /search [get]
func SearchMahasiswa(c *gin.Context) {

	search := c.Query("search")

	var mahasiswa []models.Mahasiswa

	config.DB.Where(
		"nama ILIKE ?",
		"%"+search+"%",
	).Find(&mahasiswa)

	c.JSON(200, mahasiswa)

}
