package models

type CreatePostReq struct {
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
type GetPostReq struct {
	UserId string `json:"user_id"`
	PostId string `json:"post_id"`
}
type UpdatePostReq struct {
	PostId      string `json:"post_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type DeletePostReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
type ListPostsReq struct {
	UserId string `json:"user_id"`
}
type LikePostReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
type LikeDeleteReq struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
type Empty struct{}
type GetPostRes struct {
	PostId      string `json:"post_id"`
	UserId      string `json:"usr_id"`
	Title       string `json:"title"`
	Description string `json:"desctiption"`
	Image       string `json:"image"`
	CheckLike   bool   `json:"check_like"`
}
type ListUserPosts struct {
	Posts []GetPostRes `json:"posts"`
}
type ListPostsRes struct {
	Posts []GetPostRes `json:"posts"`
}
type Bool struct {
	Like bool `json:"check_like"`
}
