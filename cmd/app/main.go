package main

import (
	"MusicAPI/docs"
	"MusicAPI/internal/music"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/webstradev/gin-pagination/v2/pkg/pagination"
	"net/http"
	"os"
)

// @title MusicAPI
// @version 1.0

func main() {

	err := godotenv.Load("deploy/.env")
	if err != nil {
		log.Fatal(err)
		return
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	if gin.Mode() == gin.DebugMode {
		log.SetLevel(log.DebugLevel)
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	pgDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot open postgres connection ", err)
	}
	defer pgDB.Close()

	client := &http.Client{}

	repository := music.NewRepository(pgDB, client)
	service := music.NewService(repository)
	handler := music.NewHandler(service)

	router := gin.New()
	router.Use(gin.Recovery())

	tracksPaginator := pagination.New(
		pagination.WithMinPageSize(10),
		pagination.WithMaxPageSize(30),
	)

	trackTextPaginator := pagination.New(
		pagination.WithSizeText("verseCount"),
		pagination.WithDefaultPageSize(1),
		pagination.WithMinPageSize(1),
		pagination.WithMaxPageSize(10))

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tracks", tracksPaginator, handler.GetTracks)

	trackRouter := router.Group("/track")
	{
		trackRouter.GET("/text", trackTextPaginator, handler.GetTrackText)
		trackRouter.DELETE("/delete", handler.DeleteTrack)
		trackRouter.POST("/update", handler.UpdateTrack)
		trackRouter.POST("/add", handler.AddTrack)
	}

	addr := ":" + os.Getenv("SERVICE_PORT")
	err = router.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
