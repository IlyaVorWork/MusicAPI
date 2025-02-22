package music

import (
	"MusicAPI/internal/pkg"
	"database/sql"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"os"
)

type Repository struct {
	db     *sql.DB
	client *http.Client
}

func NewRepository(db *sql.DB, client *http.Client) *Repository {
	return &Repository{
		db:     db,
		client: client,
	}
}

func (repository *Repository) GetTracks(group, song, date, text, link string, page, size int) ([]Track, error) {
	rows, err := repository.db.Query("SELECT * FROM tracks WHERE tracks.group LIKE $1 AND song LIKE $2 AND release_date LIKE $3 AND text LIKE $4 AND link LIKE $5 LIMIT $6 OFFSET $7", "%"+group+"%", "%"+song+"%", "%"+date+"%", "%"+text+"%", "%"+link+"%", size, (page-1)*size)
	if err != nil {
		return nil, err
	}

	tracks := make([]Track, 0)
	for rows.Next() {
		var track Track
		err = rows.Scan(&track.Group, &track.Song, &track.ReleaseDate, &track.Text, &track.Link)
		if err != nil {
			return nil, err
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (repository *Repository) GetTrack(group, song string) (*Track, error) {
	var track Track

	err := repository.db.QueryRow("SELECT * FROM tracks WHERE tracks.group = $1 AND song = $2", group, song).Scan(&track.Group, &track.Song, &track.ReleaseDate, &track.Text, &track.Link)
	if err != nil {
		return nil, err
	}

	return &track, nil
}

func (repository *Repository) DeleteTrack(group, song string) error {
	_, err := repository.db.Exec("DELETE FROM tracks WHERE tracks.group = $1 AND song = $2", group, song)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) UpdateTrack(track Track) error {
	_, err := repository.db.Exec("UPDATE tracks SET release_date = $1, text = $2, link = $3 WHERE tracks.group = $4 AND song = $5", track.ReleaseDate, track.Text, track.Link, track.Group, track.Song)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) AddTrack(track Track) error {
	_, err := repository.db.Exec("INSERT INTO tracks(\"group\", song, release_date, text, link) VALUES($1, $2, $3, $4, $5)", track.Group, track.Song, track.ReleaseDate, track.Text, track.Link)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) FetchTrackInfo(group, song string) (*TrackInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info?group=%s&song=%s", os.Getenv("MUSIC_INFO_ENDPOINT"), group, song), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	res, err := repository.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Status != "200 OK" {
		return nil, pkg.NoTrackInfoError
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	trackInfo := &TrackInfo{}
	err = json.Unmarshal(body, trackInfo)
	if err != nil {
		return nil, err
	}

	return trackInfo, nil
}
