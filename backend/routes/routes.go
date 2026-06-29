package routes

import (
	"kampus-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/mahasiswa", handlers.GetMahasiswa)
	r.GET("/jurusan", handlers.GetJurusan)
	r.POST("/mahasiswa", handlers.CreateMahasiswa)
	r.PUT("/mahasiswa/:id", handlers.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", handlers.DeleteMahasiswa)
	r.GET("/search", handlers.SearchMahasiswa)

}
