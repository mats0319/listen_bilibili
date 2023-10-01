package api

type List struct {
	Playlists []Playlist `json:"playlists" yaml:"playlists"`
}

type Playlist struct {
	ID          string  `json:"id" yaml:"id"`
	Name        string  `json:"name" yaml:"name"`
	Description string  `json:"description" yaml:"description"`
	MusicList   []Music `json:"music_list" yaml:"music_list"`
}

type Music struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Bv          string `json:"bv" yaml:"bv"`
	Description string `json:"description" yaml:"description"`
	Volume      int32  `json:"volume" yaml:"volume"`
}

const URI_GetList = "/list/get"

type GetListReq struct {
}

type GetListRes struct {
	List List   `json:"list" yaml:"list"`
	Err  string `json:"err" yaml:"err"`
}

const URI_GetOriginURL = "/origin-url/get"

type GetOriginURLReq struct {
	MusicID string `json:"music_id" yaml:"music_id"`
}

type GetOriginURLRes struct {
	URL    string `json:"url" yaml:"url"`
	Volume int32  `json:"volume" yaml:"volume"`
	Err    string `json:"err" yaml:"err"`
}
