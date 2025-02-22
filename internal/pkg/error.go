package pkg

import "errors"

var (
	TrackNotExistError          = errors.New("track does not exist")
	TrackExistError             = errors.New("track already exist")
	NoTrackInfoError            = errors.New("track info cannot be found")
	GroupOrSongNotProvidedError = errors.New("group or song was not provided")
)
