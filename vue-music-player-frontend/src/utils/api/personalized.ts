import type { PersonalizedAlbum, PersonalizedArtist, PersonalizedSong } from '@/models/personalized'
import myAxios from './myAxios'

//同样的api
export async function apiGetPersonalizeSongs() {
  const result = await myAxios.get<{
    code: number
    songs: PersonalizedSong[]
    artists: null
    albums: null
  }>('discover', { type: 1 })
  return result
}

export async function apiGetPersonalizeAlbums() {
  const result = await myAxios.get<{
    code: number
    songs: null
    artists: null
    albums: PersonalizedAlbum[]
  }>('discover', { type: 2 })
  return result
}

export async function apiGetPersonalizeArtists() {
  const result = await myAxios.get<{
    code: number
    songs: null
    artists: PersonalizedArtist[]
    albums: null
  }>('discover', { type: 3 })
  return result
}
