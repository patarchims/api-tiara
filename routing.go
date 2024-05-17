package main

import (
	"io/ioutil"
	"os"
	"vincentcoreapi/app/rest"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// ROUTING APPLICATION
func (s *Service) RoutingAndListen() {

	gin.DefaultWriter = ioutil.Discard
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))

	router.Use(rest.CORSMiddleware())
	router.Use(static.Serve("/apps/files", static.LocalFile("./files", true)))
	// API VERSI 1
	api := router.Group("")

	// API PUBLIC
	apiPublic := api.Group("")

	// API PROTECTED
	apiProtected := api.Group("")
	apiProtected.Use(rest.JwtVerify())

	apiPublic.GET("/gettoken", s.UserHandler.Login)
	apiProtected.POST("/status-antrean", s.AntrianHandler.GetStatusAntrian)
	apiProtected.POST("/sisa-antrean", s.AntrianHandler.GetSisaAntrian)
	apiProtected.POST("/batal-antrean", s.AntrianHandler.BatalAntrean)
	apiProtected.POST("/check-in", s.AntrianHandler.CheckIn)
	apiProtected.POST("/pasien-baru", s.AntrianHandler.RegisterPasienBaru)
	apiProtected.POST("/get-jadwal-operasi", s.AntrianHandler.GetJadwalOperasi)
	apiProtected.POST("/list-jadwal-operasi", s.AntrianHandler.GetKodeBookingOperasi)
	apiProtected.POST("/ambil-antrean", s.AntrianHandler.AmbilAntrean)

	// NEW FITUR, ANTREAN FARMASI
	apiProtected.POST("/ambil-antrean-farmasi", s.FarmasiHandler.AmbilAntreanFarmasi)
	apiProtected.POST("/status-antrean-farmasi", s.FarmasiHandler.StatusAntreanFarmasi)

	// ==========================
	router.Run(os.Getenv("DEPLOY_PORT"))

}
