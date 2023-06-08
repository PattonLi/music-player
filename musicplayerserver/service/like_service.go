package service

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/model"
)

type LikesService struct {
	likesdao *dao.LikeDao
}

// 获取用户所有收藏
func (l *LikesService) GetUserLikes(userid int) ([]model.SongInfo, []model.AlbumInfo, []model.ArtistInfo, []model.PlaylistInfo, error) {
	likes, err := l.likesdao.GetUserLikes(userid)
	var song_id []int
	var album_id []int
	var artist_id []int
	var playlist_id []int
	var songs []model.SongInfo
	var albums []model.AlbumInfo
	var artists []model.ArtistInfo
	var playlists []model.PlaylistInfo
	songdao := &dao.Songdao{}
	albumdao := &dao.AlbumDao{}
	artistdao := &dao.ArtistDao{}
	platlistdao := &dao.PlaylistDao{}

	for i := 0; i < len(likes); i++ {
		if likes[i].Type == 1 {
			song_id = append(song_id, likes[i].Song_id)
		}

		if likes[i].Type == 2 {
			artist_id = append(artist_id, likes[i].Artist_id)
		}

		if likes[i].Type == 3 {
			album_id = append(album_id, likes[i].Album_id)
		}

		if likes[i].Type == 4 {
			playlist_id = append(playlist_id, likes[i].Playlist_id)
		}
	}

	for i := 0; i < len(song_id); i++ {
		song, _ := songdao.GetSongDetail(song_id[i])
		songs = append(songs, song)
	}

	for i := 0; i < len(album_id); i++ {
		album, _ := albumdao.GetAlbumById(album_id[i])
		albums = append(albums, album)
	}

	for i := 0; i < len(artist_id); i++ {
		artist, _ := artistdao.GetInfoById(artist_id[i])
		artists = append(artists, artist)
	}

	for i := 0; i < len(playlist_id); i++ {
		playlist, _ := platlistdao.GetInfoById(playlist_id[i])
		playlists = append(playlists, playlist)
	}

	return songs, albums, artists, playlists, err
}

// 添加用户收藏
func (l *LikesService) AddUserLikes(like model.LikesInfo) error {
	err := l.likesdao.AddUserLikes(like)
	return err
}

// 删除用户收藏
func (l *LikesService) DeleteUserLikes(like model.LikesInfo) error {
	err := l.likesdao.DeleteUserLikes(like)
	return err
}
