import Mock from 'mockjs'

const Random = Mock.Random

const generateRandomSongInfo = () => {
  const songInfo = {
    name: Random.cword(3, 7),
    artist: Random.cword(2, 5),
    album: Random.cword(2, 5),
    time: Random.date(),
    style: Random.cword(2, 5)
  }
  return songInfo
}

const generateMultipleSongInfo = (count) => {
  const songInfoList = []
  for (let i = 0; i < count; i++) {
    const songInfo = generateRandomSongInfo()
    songInfoList.push(songInfo)
  }
  return songInfoList
}

Mock.mock(RegExp('/adminUser/songInfo' + '.*'), 'get', () => {
  let data = generateMultipleSongInfo(10)
  return {
    code: 200,
    totals: 250,
    songInfo: data
  }
})
