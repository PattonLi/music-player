/**
 * comment
 */
export interface Comment {
  comment: string
  commentId: number
  commentTime: null | string
  /**
   * 点赞数
   */
  like: number | null
  nickname: string| null
  picUrl:string
}
