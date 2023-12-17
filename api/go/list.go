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
	ReloadList bool `json:"reload_list"`
}

type GetListRes struct {
	List List   `json:"list"`
	Err  string `json:"err"`
}

const URI_GetOriginURL = "/origin-url/get"

type GetOriginURLReq struct {
	MusicID string `json:"music_id"`
}

type GetOriginURLRes struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	Volume int32  `json:"volume"`
	Err    string `json:"err"`
}

const URI_ModifyList = "/list/modify"

type ModifyListReq struct {
	List string `json:"list"` // json string
}

type ModifyListRes struct {
	Err string `json:"err"`
}
