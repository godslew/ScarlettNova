package requests

type GetPingRequest struct {
	ID      int64  `form:"id" json:"id"`
	Message string `form:"message" json:"message"`
}

type PostPingRequest struct {
	ID      int64  `form:"id" json:"id"`
	Message string `form:"message" json:"message"`
}
