/**
 * comment
 */
export interface Comment {
    comment: string;
    commentId: number;
    commentTime: null | string;
    /**
     * 比如说来自天健
     */
    ipAddress: null | string;
    /**
     * 点赞数
     */
    like: number | null;
    /**
     * 仅点赞时使用，是目标评论的ID
     */
    refCommentId: number | null;
    songId: number;
    /**
     * 1代表评论2代表点赞
     */
    type: number;
    userId: number;
}