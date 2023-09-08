// Generate File, Should not Edit.
// Author: mario.

export class GetListReq {
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
  err: string = "";
}

export class List {
  id: string = "";
  name: string = "";
  author: string = "";
  playlists: Array<Playlist> = new Array<Playlist>();
}

export class Playlist {
  id: string = "";
  name: string = "";
  musicList: Array<Music> = new Array<Music>();
}

export class Music {
  id: string = "";
  name: string = "";
  bv: string = "";
  volume: number = 0;
}
