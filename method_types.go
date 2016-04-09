package telegrambot


type SendMessageObj struct {
    ChatID string `json:"chat_id"`
    Test string `json:"text"`
    ParseMode string `json:"parse_mode"`
    DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup *interface{} `json:"reply_markup,omitempty"`
}

