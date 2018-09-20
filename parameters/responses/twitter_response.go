package responses

type GetTwitterWebHookCrcCheckResponse struct {
	Token string `json:"response_token"`
}

func NewGetTwitterWebHookCrcCheckResponse() GetTwitterWebHookCrcCheckResponse {
	return GetTwitterWebHookCrcCheckResponse{}
}
