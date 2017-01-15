package telegrambot

type SendMessageObj struct {
    ChatId string `json:"chat_id"`
    Text string `json:"text"`
    ParseMode string `json:"parse_mode"`
    DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type ForwardMessageObj struct {
    ChatId string `json:"chat_id"`
    FromChatId string `json:"from_chat_id"`
    MessageId int64 `json:"message_id"`
    DisableNotification bool `json:"disable_notification,omitempty"`
}

type SendPhotoObj struct {
    ChatId string `json:"chat_id"`
    Photo interface{} `json:"photo"`
    Caption string `json:"caption,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendAudioObj struct {
    ChatId string `json:"chat_id"`
    Audio interface{} `json:"audio"`
    Duration int64 `json:"duration,omitempty"`
    Performer string `json:"performer,omitempty"`
    Title string `json:"title,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendDocumentObj struct {
    ChatId string `json:"chat_id"`
    Document interface{} `json:"document"`
    Caption string `json:"caption,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendStickerObj struct {
    ChatId string `json:"chat_id"`
    Sticker interface{} `json:"sticker"`
    Caption string `json:"caption,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendVideoObj struct {
    ChatId string `json:"chat_id"`
    Video interface{} `json:"video"`
    Duration int64 `json:"duration,omitempty"`
    Width int64 `json:"width,omitempty"`
    Height int64 `json:"width,omitempty"`
    Caption string `json:"caption,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendVoiceObj struct {
    ChatId string `json:"chat_id"`
    Voice interface{} `json:"voice"`
    Duration int64 `json:"duration,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendLocationObj struct {
    ChatId string `json:"chat_id"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendVenueObj struct {
    ChatId string `json:"chat_id"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Title string `json:"title"`
    Adress string `json:"adress"`
    FoursquareId string `json:"foursquare_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendContactObj struct {
    ChatId string `json:"chat_id"`
    PhoneNumber string `json:"phone_number"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type SendChatActionObj struct {
    ChatId string `json:"chat_id"`
    Action string `json:"action"`
}

type GetUserProfilePhotosObj struct {
    UserId int64 `json:"user_id"`
    Offset int64 `json:"offset,omitempty"`
    Limit int64 `json:"limit,omitempty"`
}

type GetFileObj struct {
    FileId string `json:"file_id"`
}

type KickChatMemberObj struct {
    ChatId string `json:"chat_id"`
    UserId int64 `json:"user_id"`
}

type LeaveChatObj struct {
    ChatId string `json:"chat_id"`
}

type UnbanChatMemberObj struct {
    ChatId string `json:"chat_id"`
    UserId int64 `json:"user_id"`
}

type GetChatObj struct {
    ChatId string `json:"chat_id"`
}

type GetChatAdministratorsObj struct {
    ChatId string `json:"chat_id"`
}

type GetChatMembersCountObj struct {
    ChatId string `json:"chat_id"`
}

type GetChatMemberObj struct {
    ChatId string `json:"chat_id"`
    UserId int64 `json:"user_id"`
}

type AnswerCallbackQueryObj struct {
    CallbackQueryId string `json:"callback_query_id"`
    Text string `json:"text,omitempty"`
    ShowAlert bool `json:"show_alert,omitempty"`
}

type EditMessageTextObj struct {
    ChatId string `json:"chat_id,omitempty"`
    MessageId string `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    Text string `json:"text"`
    ParseMode string `json:"parse_mode,omitempty"`
    DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCaptionObj struct {
    ChatId string `json:"chat_id,omitempty"`
    MessageId string `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    Caption string `json:"caption"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkupObj struct {
    ChatId string `json:"chat_id,omitempty"`
    MessageId string `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}


type AnswerInlineQueryObj struct {
    InlineQueryId string `json:"inline_query_id"`
    Results []interface{} `json:"results"`
    CacheTime int64 `json:"cache_time,omitempty"`
    IsPersonal bool `json:"is_personal,omitempty"`
    NextOffset string `json:"next_offset,omitempty"`
    SwitchPmText string `json:"switch_pm_text,omitempty"`
    SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}

type UpdateObj struct {
    Offset int64 `json:"offset,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Timout int64 `json:"timeout,omitempty"`
}

type SetWebhookObj struct {
    Url string `json:"url"`
    Certificate interface{} `json:"interface,omitempty"`
}
