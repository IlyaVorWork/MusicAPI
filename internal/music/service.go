package music

import (
	"MusicAPI/internal/pkg"
	"strings"
)

type IRepository interface {
	GetTracks(group, song, date, text, link string, page, size int) ([]Track, error)
	GetTrack(group, song string) (*Track, error)
	DeleteTrack(group, song string) error
	UpdateTrack(track Track) error
	AddTrack(track Track) error

	FetchTrackInfo(group, song string) (*TrackInfo, error)
}

type Service struct {
	repository IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) GetTracks(group, song, date, text, link string, page, size int) ([]Track, error) {
	tracks, err := service.repository.GetTracks(group, song, date, text, link, page, size)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (service *Service) GetTrackText(group, song string, page, verseCount int) (*string, error) {
	track, err := service.repository.GetTrack(group, song)
	if err != nil {
		return nil, pkg.TrackNotExistError
	}

	verses := strings.Split(track.Text, "\\n\\n")

	var text string
	if page-1+verseCount > len(verses) {
		text = strings.Join(verses[(page-1)*verseCount:], "\\n\\n")
	} else {
		text = strings.Join(verses[(page-1)*verseCount:page-1+verseCount], "\\n\\n")
	}

	return &text, nil
}

func (service *Service) DeleteTrack(data DeleteTrackDTO) error {
	_, err := service.repository.GetTrack(data.Group, data.Song)
	if err != nil {
		return pkg.TrackNotExistError
	}

	err = service.repository.DeleteTrack(data.Group, data.Song)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) UpdateTrack(data UpdateTrackDTO) error {
	track, err := service.repository.GetTrack(data.Group, data.Song)
	if err != nil {
		return pkg.TrackNotExistError
	}

	track.ReleaseDate = data.NewReleaseDate
	track.Text = data.NewText
	track.Link = data.NewLink

	err = service.repository.UpdateTrack(*track)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) AddTrack(data AddTrackDTO) error {
	_, err := service.repository.GetTrack(data.Group, data.Song)
	if err == nil {
		return pkg.TrackExistError
	}

	trackInfo, err := service.repository.FetchTrackInfo(data.Group, data.Song)
	if err != nil {
		return err
	}

	var track Track
	track.Group = data.Group
	track.Song = data.Song

	track.ReleaseDate = trackInfo.ReleaseDate
	track.Text = trackInfo.Text
	track.Link = trackInfo.Link

	err = service.repository.AddTrack(track)
	if err != nil {
		return err
	}

	return nil
}
