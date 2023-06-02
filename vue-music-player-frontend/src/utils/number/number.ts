import dayjs from 'dayjs'

export function numberToDuration(during: number) {
  const s = Math.floor(during) % 60
  during = Math.floor(during / 60)
  const i = during % 60

  const ii = i < 10 ? `0${i}` : i
  const ss = s < 10 ? `0${s}` : s

  return ii + ':' + ss
}

//返回日期格式
export function toDate(this: number, format: string = 'YYYY-MM-DD'): string {
  return dayjs(this).format(format)
}
