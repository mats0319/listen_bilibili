// Generate File, Should not Edit.
// Author: mario. https://github.com/mats9693
// Version: goc_ts v0.2.1

export class List {
    playlists: Array<Playlist> = new Array<Playlist>();
}

export class Playlist {
    id: string = "";
    name: string = "";
    description: string = "";
    music_list: Array<Music> = new Array<Music>();
}

export class Music {
    id: string = "";
    name: string = "";
    bv: string = "";
    description: string = "";
    volume: number = 0;
}

export class GetListReq {
    reload_list: boolean = false;
}

export class GetListRes {
    list: List = new List();
    err: string = "";
}

export class GetOriginURLReq {
    music_id: string = "";
}

export class GetOriginURLRes {
    url: string = "";
    name: string = "";
    volume: number = 0;
    err: string = "";
}

export class ModifyListReq {
    list: string = "";
}

export class ModifyListRes {
    err: string = "";
}
