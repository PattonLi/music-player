import type { Album } from './album'
import type { Song } from './song'

export interface Artist {
  artistId: number
  artist: string
  picUrl: string
  songSize: number
  albumSize: number
  profile: string
  location: string
}

export interface ArtistDetail {
  artist: Artist
  songs: Song[]
  albums: Album[]
}
