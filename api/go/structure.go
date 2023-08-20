package api

type List struct {
	ID        string
	Name      string
	Author    string
	Playlists []Playlist
}

type Playlist struct {
	ID        string
	Name      string
	MusicList []Music
}

type Music struct {
	ID     string
	Name   string
	Bv     string
	volume int32
}

type GetListRes struct {
	List string
	Err  string
}

type GetOriginURLRes struct {
	URL string
	Err string
}
