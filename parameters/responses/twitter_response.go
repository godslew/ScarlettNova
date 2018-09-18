package responses

type TwitterWebHookCrcCheckResponse struct {
	Token string `json:"response_token"`
}

func NewTwitterWebHookCrcCheckResponse() TwitterWebHookCrcCheckResponse {
	return TwitterWebHookCrcCheckResponse{}
}
