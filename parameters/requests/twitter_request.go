package requests

// GET Request (CRC Check)
type TwitterWebhookGetRequest struct {
	CrcToken string `json:"crc_token" mapstructure:"crc_token"`
}

func NewTwitterWebhookGetRequest() TwitterWebhookGetRequest {
	return TwitterWebhookGetRequest{}
}

func (r *TwitterWebhookGetRequest) Validate() error {
	return nil
}

// POST Request
type TwitterWebhookPostRequest struct {
	DirectMessageEvents []struct {
		Type          string `json:"type"`
		ID            string `json:"id"`
		MessageCreate struct {
			SenderID    string `json:"sender_id"`
			MessageData struct {
				Text               string `json:"text"`
				QuickReplyResponse struct {
					MetaData string `json:"metadata"`
				} `json:"quick_reply_response"`
			} `json:"message_data"`
		} `json:"message_create"`
	} `json:"direct_message_events"`
}

func NewTwitterWebhookPostRequest() TwitterWebhookPostRequest {
	return TwitterWebhookPostRequest{}
}

func (r *TwitterWebhookPostRequest) Validate() error {
	return nil
}

type QuickReplyOption struct {
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	Metadata    string `json:"metadata"`
}

func newQuickReplyOption(label, metadata string) QuickReplyOption {
	return QuickReplyOption{
		Label:    label,
		Metadata: metadata,
	}
}

type PostDirectMessageRequest struct {
	Event struct {
		Type          string `json:"type"`
		MessageCreate struct {
			Target struct {
				RecipientID string `json:"recipient_id"`
			} `json:"target"`
			MessageData struct {
				Text       string `json:"text"`
				Ctas       *Ctas  `json:"ctas,omitempty"`
				Attachment struct {
					Type  string `json:"type"`
					Media struct {
						ID string `json:"id"`
					} `json:"media"`
				} `json:"attachment"`
				QuickReply *QuickReply `json:"quick_reply,omitempty"`
			} `json:"message_data"`
		} `json:"message_create"`
	} `json:"event"`
}

type QuickReply struct {
	Type    string             `json:"type"`
	Options []QuickReplyOption `json:"options"`
}

type Ctas []struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	URL   string `json:"url"`
}
