package api

type List struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	Playlists []Playlist `json:"playlists"`
}

type Playlist struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	MusicList []Music `json:"musicList"`
}

type Music struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Bv     string `json:"bv"`
	Volume int32  `json:"volume"`
}

type GetListRes struct {
	List string `json:"list"`
	Err  string `json:"err"`
}

type GetOriginURLRes struct {
	URL string `json:"url"`
	Err string `json:"err"`
}
