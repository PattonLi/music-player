package service

import (
	"errors"
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type PlaylistService struct {
	playlistdao *dao.PlaylistDao
}

// 获取热门歌单
func (p *PlaylistService) GetHotPlaylist(pagesize int, currentpage int) ([]model.PlaylistInfo, error, error, int) {
	playlist, err := p.playlistdao.GetAllPlaylist()
	l := len(playlist)
	n := l / pagesize
	remainder := l % pagesize
	var pagenum int
	var remainder_flag bool
	if remainder == 0 {
		pagenum = n
		remainder_flag = false
	} else {
		pagenum = n + 1
		remainder_flag = true
	}

	if currentpage > pagenum {
		err1 := errors.New("当前要获取的歌单分页过大！")
		return nil, err, err1, pagenum
	}

	var playlistpage []model.PlaylistInfo

	if currentpage == pagenum && remainder_flag {
		playlistpage = playlist[(currentpage-1)*pagesize : l]
		return playlistpage, err, nil, pagenum
	}

	playlistpage = playlist[(currentpage-1)*pagesize : currentpage*pagesize]
	return playlistpage, err, nil, pagenum
}

// 获取歌单所有歌曲
func (p *PlaylistService) GetPlaylistAllSongs(playlistid int) (model.PlaylistInfo, []model.SongInfo, bool) {
	playlist, songs, err := p.playlistdao.GetPlaylistAllSongs(playlistid)
	return playlist, songs, err
}
