package music

type Track struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type DeleteTrackDTO struct {
	Group string `json:"group" required:"true"`
	Song  string `json:"song" required:"true"`
}

type UpdateTrackDTO struct {
	Group          string `json:"group" required:"true"`
	Song           string `json:"song" required:"true"`
	NewReleaseDate string `json:"newReleaseDate"`
	NewText        string `json:"newText"`
	NewLink        string `json:"newLink"`
}

type AddTrackDTO struct {
	Group string `json:"group" required:"true"`
	Song  string `json:"song" required:"true"`
}

type TrackInfo struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
