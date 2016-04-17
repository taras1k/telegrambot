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

type ForwardMessageObj struct {
    ChatID string `json:"chat_id"`
    FromChatID string `json:"from_chat_id"`
    MessageID int64 `json:"message_id"`
    DisableNotification bool `json:"disable_notification,omitempty"`
}

type SendPhotoObj struct {
    ChatID string `json:"chat_id"`
    Photo interface{} `json:"photo"`
    Caption string `json:caption,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup *interface{} `json:"reply_markup,omitempty"`
}

type UpdateObj struct {
    Offset int64 `json:"offset,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Timout int64 `json:"timeout,omitempty"`
}
