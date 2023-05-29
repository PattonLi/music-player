import { defineStore, storeToRefs } from 'pinia'
import { apiGetSong } from '@/utils/api/song'
import type { Song } from '@/models/song'

export const usePlayerStore = defineStore('player', {
  state: () => ({
    audio: new Audio(),

    playList: [] as Song[], //播放列表,
    song: {} as Song, //正在播放歌曲

    showPlayList: false,
    loopType: 0, //0 单曲循环 1 列表循环 2列表随机播放
    volume: 60, //音量
    isPlaying: false, //是否正在播放
    ended: false, //是否播放结束
    muted: false, //是否静音
    currentTime: 0, //当前播放时间
    duration: 0 //总播放时长
  }),
  getters: {
    //播放列表歌曲数量
    getPlayListCount: (state) => {
      return state.playList.length
    },
    //正在播放歌曲在播放列表中index
    getPlayListIndex: (state) => {
      return state.playList.findIndex((song) => song.id === state.song.id)
    },
    //获取下一首歌曲
    getNextSong(state): Song {
      if (this.getPlayListIndex === this.getPlayListCount - 1) {
        return state.playList[0] //回到播放列表头部
      } else {
        //正常情况
        const nextIndex: number = this.getPlayListIndex + 1
        return state.playList[nextIndex]
      }
    },
    //获取前一首歌曲
    prevSong(state): Song {
      if (this.getPlayListIndex === 0) {
        return state.playList.last() //回到播放列表尾部
      } else {
        //正常情况
        const prevIndex: number = this.getPlayListIndex - 1
        return state.playList[prevIndex]
      }
    }
  },
  actions: {
    //播放列表里面添加音乐
    pushPlayList(replace: boolean, ...list: Song[]) {
      //替换播放列表
      if (replace) {
        this.playList = list
        return
      }
      //不替换播放列表
      list.forEach((song) => {
        //播放列表中没有该歌曲
        if (this.playList.filter((s) => s.id == song.id).length <= 0) {
          this.playList.push(song)
        }
      })
    },
    //清空播放列表
    clearPlayList() {
      this.song = {} as Song
      this.playList = [] as Song[]
      this.isPlaying = false
      this.ended = false
      this.muted = false
      this.currentTime = 0
      this.showPlayList = false
      this.audio.load() //Resets the audio
      setTimeout(() => {
        this.duration = 0
      }, 500)
    },
    //开始播放
    async play(id: number) {
      if (id == this.song.id) return
      this.isPlaying = false
      //获取歌曲信息
      this.song = await apiGetSong(this.song.id)
      this.audio.src = this.song.url
      this.audio
        .play()
        .then((res) => {
          //播放成功
          this.isPlaying = true
          this.pushPlayList(false, this.song)
        })
        .catch((res) => {
          //播放失败
          console.log(res)
        })
    },
    //重新播放
    rePlay() {
      setTimeout(() => {
        this.currentTime = 0
        this.audio.play()
      }, 1000)
    },
    //下一曲
    nextPlay() {
      switch (this.loopType) {
        //单曲循环
        case 0:
          this.rePlay()
          break
        //列表
        case 1:
          this.play(this.getNextSong.id)
          break
        //随机
        case 2:
          this.play(this.playList.sample().id)
          break
      }
    },
    //上一曲
    prev() {
      this.play(this.prevSong.id)
    },
    //播放、暂停
    togglePlay() {
      //没有歌曲在播放
      if (!this.song.id) return
      this.isPlaying = !this.isPlaying
      //开始播放
      if (!this.isPlaying) {
        this.audio.play()
      } //暂停播放
      else {
        this.audio.pause()
      }
    },

    //切换循环类型
    toggleLoop() {
      if (this.loopType == 2) {
        this.loopType = 0
      } else {
        this.loopType++
      }
    },
    //静音切换
    toggleMuted() {
      this.muted = !this.muted
      this.audio.muted = this.muted
    },
    //音量设置
    setVolume(n: number) {
      //范围约束
      n = n > 100 ? 100 : n
      n = n < 0 ? 0 : n
      this.volume = n
      this.audio.volume = n / 100
    },
    //修改进度条
    onSliderChange(val: any) {
      //待做转换成时间格式
      this.currentTime = val
      // this.audio.currentTime = val
    },
    //定时器更新信息
    interval() {
      //如果在播放
      if (this.isPlaying) {
        this.currentTime = parseInt(this.audio.currentTime.toString())
        this.duration = parseInt(this.audio.duration.toString())
        this.ended = this.audio.ended
      }
    }
  }
})

export const userPlayerInit = () => {
  let timer: any
  const { setVolume, interval, nextPlay } = usePlayerStore()
  const { ended } = storeToRefs(usePlayerStore())

  //监听播放结束
  watch(ended, (ended) => {
    if (!ended) return
    nextPlay()
  })

  //启动定时器
  onMounted(() => {
    setVolume(60)
    console.log('启动定时器')
    timer = setInterval(interval, 1000)
  })
  //清除定时器
  onUnmounted(() => {
    console.log('清除定时器')
    clearInterval(timer)
  })
}
