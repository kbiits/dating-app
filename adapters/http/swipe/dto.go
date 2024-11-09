package http_swipe

type SwipeProfileReq struct {
	ProfileID string `json:"profile_id" validate:"required"`
	IsLiked   bool   `json:"is_liked" validate:"required"`
}
