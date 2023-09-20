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
	MusicList []Music `json:"music_list"`
}

type Music struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Bv     string `json:"bv"`
	Volume int32  `json:"volume"`
}

const URI_getList = "/list/get"

type GetListReq struct {
}

type GetListRes struct {
	List string `json:"list"`
	Err  string `json:"err"`
}

const URI_getOriginalURL = "/originURL/get"

type GetOriginURLReq struct {
	MusicID string `json:"music_id"`
}

type GetOriginURLRes struct {
	URL string `json:"url"`
	Err string `json:"err"`
}
