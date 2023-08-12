package constant

import "time"

type CommentActionParams struct {
	UserID      int64
	VideoId     int64  `form:"video_id" binding:"required"`
	ActionType  int64  `form:"action_type" binding:"required,oneof=1 2"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type FavoriteActionParams struct {
	UserID     int64
	VideoId    int64 `form:"video_id" binding:"required"`
	ActionType int64 `form:"action_type" binding:"required,oneof=1 2"`
}

type RelationActionParams struct {
	UserID     int64
	ToUserID   int64 `form:"to_user_id" binding:"required"`
	ActionType int64 `form:"action_type" binding:"required,oneof=1 2"`
}

type Message struct {
	Id         int64
	ToUserId   int64  `form:"to_user_id" binding:"required"`
	FromUserId int64  `form:"from_user_id"`
	Content    string `form:"content" binding:"required"`
	CreateTime time.Time
	ActionType int64 `form:"action_type" binding:"required,oneof=1 2"`
}
