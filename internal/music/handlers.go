package music

import (
	"MusicAPI/internal/pkg"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type IService interface {
	GetTracks(group, song, date, text, link string, page, size int) ([]Track, error)
	GetTrackText(group, song string, page, verseCount int) (*string, error)
	DeleteTrack(data DeleteTrackDTO) error
	UpdateTrack(data UpdateTrackDTO) error
	AddTrack(data AddTrackDTO) error
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

// Handle of GetTracks
// @Tags Track
// @Description Returns a list of tracks paginated by 10 items per page by default
// @Produce json
// @Param group query string false "filter for 'group' field"
// @Param song query string false "filter for 'song' field"
// @Param date query string false "filter for 'date' field"
// @Param text query string false "filter for 'text' field"
// @Param link query string false "filter for 'link' field"
// @Param page query int false "page number"
// @Param size query int false "number of items per page"
// @Success 200 {object} pkg.TracksRes
// @Failure 500 {object} pkg.ErrorRes
// @Router /tracks [get]
func (handler *Handler) GetTracks(c *gin.Context) {
	log.Info("Query for tracks library")

	group := c.Query("group")
	song := c.Query("song")
	date := c.Query("date")
	text := c.Query("text")
	link := c.Query("link")

	log.Debugf("Group: %s, Song: %s, Date: %s, Text: %s, Link: %s", group, song, date, text, link)

	page := c.GetInt("page")
	size := c.GetInt("size")

	log.Debugf("Page number: %d, page size: %d", page, size)

	tracks, err := handler.service.GetTracks(group, song, date, text, link, page, size)
	if err != nil {
		log.Error("Error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Tracks were successfully provided")
	c.JSON(http.StatusOK, tracks)
	return
}

// Handle of GetTrackText
// @Tags Track
// @Description Returns a track's text paginated by 1 verse per page by default
// @Produce json
// @Param group query string true "filter for 'group' field"
// @Param song query string true "filter for 'song' field"
// @Param page query int false "page number"
// @Param verseCount query int false "number of verses per page"
// @Success 200 {object} pkg.TrackTextRes
// @Failure 400 {object} pkg.ErrorRes
// @Failure 500 {object} pkg.ErrorRes
// @Router /track/text [get]
func (handler *Handler) GetTrackText(c *gin.Context) {
	log.Info("Query for track's text")
	group := c.Query("group")
	song := c.Query("song")

	if group == "" || song == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": pkg.GroupOrSongNotProvidedError})
		return
	}

	log.Debugf("Group: %s, song: %s", group, song)

	page := c.GetInt("page")
	verseCount := c.GetInt("verseCount")

	log.Debugf("Page number: %d, verses per page: %d", page, verseCount)

	text, err := handler.service.GetTrackText(group, song, page, verseCount)
	if err != nil {
		log.Error("Error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Track's text was successfully provided")
	c.JSON(http.StatusOK, gin.H{"text": text})
	return
}

// Handle of DeleteTrack
// @Tags Track
// @Description Deletes a track from the database by provided group and song name
// @Produce json
// @Param dto body DeleteTrackDTO true "contains group and song name of track to delete"
// @Success 200 {object} pkg.InfoRes
// @Failure 400 {object} pkg.ErrorRes
// @Failure 500 {object} pkg.ErrorRes
// @Router /track/delete [delete]
func (handler *Handler) DeleteTrack(c *gin.Context) {
	log.Info("Query for deleting a track")

	var queryData DeleteTrackDTO
	err := c.ShouldBind(&queryData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debugf("Group: %s, song: %s", queryData.Group, queryData.Song)

	err = handler.service.DeleteTrack(queryData)
	if err != nil {
		log.Error("Error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Track was successfully deleted")
	c.JSON(http.StatusOK, gin.H{"info": "success"})
	return
}

// Handle of UpdateTrack
// @Tags Track
// @Description Updates a track's info in the database
// @Produce json
// @Param dto body UpdateTrackDTO true "contains group and song name and info fields to update"
// @Success 200 {object} pkg.InfoRes
// @Failure 400 {object} pkg.ErrorRes
// @Failure 500 {object} pkg.ErrorRes
// @Router /track/update [post]
func (handler *Handler) UpdateTrack(c *gin.Context) {
	log.Info("Query for updating a track")

	var queryData UpdateTrackDTO
	err := c.ShouldBind(&queryData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debugf("Group: %s, song: %s, new release date: %s, new text: %s, new link: %s", queryData.Group, queryData.Song, queryData.NewReleaseDate, queryData.NewText, queryData.NewLink)

	err = handler.service.UpdateTrack(queryData)
	if err != nil {
		log.Error("Error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Track was successfully updated")
	c.JSON(http.StatusOK, gin.H{"info": "success"})
	return
}

// Handle of AddTrack
// @Tags Track
// @Description Fetches track's info and adds it into the database
// @Produce json
// @Param dto body AddTrackDTO true "contains group and song name of track to add"
// @Success 200 {object} pkg.InfoRes
// @Failure 400 {object} pkg.ErrorRes
// @Failure 500 {object} pkg.ErrorRes
// @Router /track/add [post]
func (handler *Handler) AddTrack(c *gin.Context) {
	log.Info("Query for adding a track")

	var queryData AddTrackDTO
	err := c.ShouldBind(&queryData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Debugf("Group: %s, song: %s", queryData.Group, queryData.Song)

	err = handler.service.AddTrack(queryData)
	if err != nil {
		log.Error("Error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("Track was successfully added")
	c.JSON(http.StatusOK, gin.H{"info": "success"})
	return
}
