package dao

import (
	"music-player/musicplayerserver/model"
)

type PlaylistDao struct {
}

// 获取所有歌单
func (p *PlaylistDao) GetAllPlaylist() ([]model.PlaylistInfo, error) {
	var playlist []model.PlaylistInfo
	result := DB.Find(&playlist)
	return playlist, result.Error
}

// 根据playlist_id获取歌单的所有歌曲
func (p *PlaylistDao) GetPlaylistAllSongs(playlistid int) (model.PlaylistInfo, []model.SongInfo, bool) { //bool返回值用以判断有无数据库操作错误
	var playlist model.PlaylistInfo
	var playlist_songs []model.Playlist_Songs
	var songs []model.SongInfo
	var songid int
	result0 := DB.First(&playlist, playlistid)
	result1 := DB.Where("playlist_id = ?", playlistid).Find(&playlist_songs)

	for i := 0; i < len(playlist_songs); i++ {
		var song model.SongInfo
		songid = playlist_songs[i].Song_id
		result2 := DB.First(&song, songid)
		songs = append(songs, song)

		if result2.Error != nil {
			return playlist, nil, false
		}
	}

	if result0.Error == nil && result1.Error == nil {
		return playlist, songs, true
	}

	return playlist, songs, false
}
