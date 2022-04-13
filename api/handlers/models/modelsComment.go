package models

type CreateCommentReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
type GetCommentReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
type UpdateCommentReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
type DeleteCommentReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
type GetCommentRes struct {
	UserId string `json:"user_id"`
	Text   string `json:"text"`
}
