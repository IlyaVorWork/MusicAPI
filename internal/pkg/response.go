package pkg

type Track struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type ErrorRes struct {
	Error string `json:"error"`
}

type InfoRes struct {
	Info string `json:"info"`
}

type TracksRes struct {
	Tracks []Track `json:"tracks"`
}

type TrackTextRes struct {
	Text string `json:"text"`
}
