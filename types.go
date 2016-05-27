package telegrambot

//structs for all https://core.telegram.org/bots/api#available-types

type User struct {
    Id int64 `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    Username string `json:"username,omitempty"`
}

type Chat struct {
    Id int64 `json:"id"`
    Type string `json:"type"`
    Title string `json:"title,omitempty"`
    Username string `json:"username,omitempty"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}

type ChatMember struct {
    User *User `json:"user"`
    Status string `json:"status"`
}

type Message struct {
    MessageId int64 `json:"message_id"`
    From *User `json:"from,omitempty"`
    Date int64 `json:"date"`
    Chat *Chat `json:"chat"`
    ForwardFrom *User `json:"forward_from,omitempty"`
    ForwardDate int64 `json:"forward_date,omitempty"`
    ReplyToMessage *Message `json:"reply_to_message,omitempty"`
    Text string `json:"text,omitempty"`
    Audio *Audio `json:"audio,omitempty"`
    Document *Document `json:"document,omitempty"`
    Photo []*PhotoSize `json:"photo,omitempty"`
    Sticker Sticker `json:"sticker,omitempty"`
    Video *Video `json:"video,omitempty"`
    Voice *Voice `json:"voice,omitempty"`
    Caption string `json:"caption,omitempty"`
    Contact *Contact `json:"contact,omitempty"`
    Location *Location `json:"location,omitempty"`
    NewChatParticipant *User `json:"new_chat_participant,omitempty"`
    LeftChatParticipant *User `json:"left_chat_participant,omitempty"`
    NewChatTitle string `json:"new_chat_title,omitempty"`
    NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`
    DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
    GroupChatCreated bool `json:"group_chat_created,omitempty"`
    SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
    ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
    MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
    MigrateFromChatId int64 `json:"migrate_from_chat_id,omitempty"`
}

type PhotoSize struct {
    FileId string `json:"file_id"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    FileSize int64 `json:"file_size,omitempty"`
}

type Audio struct {
    FileId string `json"file_id"`
    Duration int64 `json:"duration"`
    Performer string `json:"performer,omitempty"`
    Title string `json:"title,omitempty"`
    Mietype string `json:"mime_type"`
    FileSize int64 `json:"file_size,omitempty"`
}

type Document struct {
    FileId string `json:"file_id"`
    Thumb *PhotoSize `json:"thumb,omitempty"`
    FileName string `json:"file_name,omitempty"`
    MimeType string `json:"mime_type,omitempty"`
    FileSize int64 `json:"file_size,omitempty"`
}

type Sticker struct {
    FileId string `json:"file_id"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Thumb *PhotoSize `json:"thumb,omitempty"`
    FileSize int64 `json:"file_size,omitempty"`
}

type Video struct {
    FileId string `json"file_id"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Duration int64 `json:"duration"`
    Thumb *PhotoSize `json:"thumb,omitempty"`
    MimeType string `json:"mime_type,omitempty"`
    FileSize int64 `json:"file_size,omitempty"`
}

type Voice struct {
    FileId string `json"file_id"`
    Duration int64 `json:"duration"`
    MimeType string `json:"mime_type,omitempty"`
    FileSize int64 `json:"file_size,omitempty"`
}

type Contact struct {
    PhoneNumber string `json:"phone_number"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    UserId int64 `json:"user_id,omitempty"`
}

type Location struct {
    Longitude float64 `json:"longitude"`
    Latitude float64 `json:"latitude"`
}

type UserProfilePhotos struct {
    TotalCount int64 `json:"total_count"`
    Photos [][]*PhotoSize `json:"photos"`
}

type File struct {
    FileId string `json"file_id"`
    FileSize int64 `json:"file_size,omitempty"`
    FilePath string `json:"file_path,omitempty"`
}

type ReplyKeyboardMarkup struct {
    Keyboard [][]string `json:"keyboard"`
    ResizeKeyBoard bool `json:"resize_keyboard,omitempty"`
    OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
    Selective bool `json:"selective,omitempty"`
}

type ReplyKeyboardHide struct {
    HideKeyboard bool `json:"hide_keyboard"`
    Selective bool `json:"selective,omitempty"`
}

type ForceReply struct {
    ForceReply bool `json:"force_reply"`
    Selective bool `json:"selective,omitempty"`
}

type InlineQuery struct {
    Id string `json:"id"`
    From *User `json:"from"`
    Query string `json:"query"`
    Offset string `json:"offset"`
}

type Update struct {
    UpdateId int64 `json:"update_id"`
    Message *Message `json:"message,omitempty"`
    InlineQuery *InlineQuery `json:"inline_query,omitempty"`
    ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
}

type InlineKeyboardButton struct {
    Text string `json:"text"`
    Url string `json:"url,omitempty"`
    CallbackData string `json:"callback_data,omitempty"`
    SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
}

type InlineKeyboardMarkup struct {
    InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type CallbackQuery struct {
    Id string `json:"id"`
    From *User `json:"from"`
    Message *Message `json:"message,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    Data string `json:"data"`
}

type AnswerInlineQuery struct {
    InlineQueryId string `json:"inline_query_id"`
    Results []interface{} `json:"results"`
    CacheTime int64 `json:"cache_time,omitempty"`
    IsPersonal bool `json:"is_personal,omitempty"`
    NextOffset string `json:"next_offset,omitempty"`
    SwitchPmText string `json:"switch_pm_text,omitempty"`
    SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}


type InlineQueryResultArticle struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Title string `json:"title"`
    InputMessageContent interface{} `json:"input_message_content"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    Url string `json:"url,omitempty"`
    HideUrl bool `json:"hide_ur,omitempty"`
    Description string `json:"description,omitempty"`
    ThumbUrl string `json:"thumb_url,omitempty"`
    ThumbWidth int64 `json:"thumb_width,omitempty"`
    ThumbHeight int64 `json:"thumb_height,omitempty"`

}

type InlineQueryResultPhoto struct {
    Type string `json:"type"`
    Id string `json:"id"`
    PhotoUrl string `json:"photo_url"`
    ThumbUrl string `json:"thumb_url"`
    PhotoWidth int64 `json:"photo_width,omitempty"`
    PhotoHeight int64 `json:"photo_height,omitempty"`
    Title string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content"`
}

type InlineQueryResultGif struct {
    Type string `json:"type"`
    Id string `json:"id"`
    GifUrl string `json:"gif_url"`
    GifWidth int64 `json:"gif_width,omitempty"`
    GifHeight int64 `json:"gif_height,omitempty"`
    ThumbUrl string `json:"thumb_url"`
    Title string `json:"title,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultMpeg4Gif struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Mpeg4Url string `json:"mpeg4_url"`
    Mpeg4Width int64 `json:"mpeg4_width,omitempty"`
    Mpeg4Height int64 `json:"mpeg4_height,omitempty"`
    ThumbUrl string `json:"thumb_url"`
    Title string `json:"title,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
    Type string `json:"type"`
    Id string `json:"id"`
    VideoUrl string `json:"video_url"`
    MimeType string `json:"mime_type"`
    ThumbUrl string `json:"thumb_url"`
    Title string `json:"title"`
    Caption string `json:"caption,omitempty"`
    VideoWidth int64 `json:"video_width,omitempty"`
    VideoHeight int64 `json:"video_height,omitempty"`
    VideoDuration int64 `json:"video_duration,omitempty"`
    Description string `json:"description,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultAudio struct {
    Type string `json:"type"`
    Id string `json:"id"`
    AudioUrl string `json:"audio_url"`
    Title string `json:"title"`
    Performer string `json:"performer,omitempty"`
    AudioDuration int64 `json:"audio_duration,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
    Type string `json:"type"`
    Id string `json:"id"`
    VoiceUrl string `json:"voice_url"`
    Title string `json:"title"`
    VoiceDuration int64 `json:"voice_duration,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultDocument struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Title string `json:"title"`
    Caption string `json:"caption,omitempty"`
    DocumentUrl string `json:"document_url"`
    MimeType string `json:"mime_type"`
    Description string `json:"description,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`
    ThumbUrl string `json:"thumb_url,omitempty"`
    ThumbWidth int64 `json:"thumb_width,omitempty"`
    ThumbHeight int64 `json:"thumb_height,omitempty"`

}

type InlineQueryResultLocation struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Latitude string `json:"latitude"`
    Longitude string `json:"longitude"`
    Title string `json:"title"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`
    ThumbUrl string `json:"thumb_url,omitempty"`
    ThumbWidth int64 `json:"thumb_width,omitempty"`
    ThumbHeight int64 `json:"thumb_height,omitempty"`

}

type InlineQueryResultVenue struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Title string `json:"title"`
    Address string `json:"address"`
    FoursquareId string `json:"foursquare_id,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`
    ThumbUrl string `json:"thumb_url,omitempty"`
    ThumbWidth int64 `json:"thumb_width,omitempty"`
    ThumbHeight int64 `json:"thumb_height,omitempty"`

}

type InlineQueryResultContact struct {
    Type string `json:"type"`
    Id string `json:"id"`
    PhoneNumber string `json:"phone_number"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`
    ThumbUrl string `json:"thumb_url,omitempty"`
    ThumbWidth int64 `json:"thumb_width,omitempty"`
    ThumbHeight int64 `json:"thumb_height,omitempty"`

}

type InlineQueryResultCachedPhoto struct {
    Type string `json:"type"`
    Id string `json:"id"`
    PhotoFile_id string `json:"photo_file_id"`
    Title string `json:"title,omitempty"`
    Description string `json:"description,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedGif struct {
    Type string `json:"type"`
    Id string `json:"id"`
    GifFile_id string `json:"gif_file_id"`
    Title string `json:"title,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedMpeg4Gif struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Mpeg4File_id string `json:"mpeg4_file_id"`
    Title string `json:"title,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedSticker struct {
    Type string `json:"type"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedDocument struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Title string `json:"title"`
    DocumentFile_id string `json:"document_file_id"`
    Description string `json:"description,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedVideo struct {
    Type string `json:"type"`
    Id string `json:"id"`
    VideoFile_id string `json:"video_file_id"`
    Title string `json:"title"`
    Description string `json:"description,omitempty"`
    Caption string `json:"caption,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedVoice struct {
    Type string `json:"type"`
    Id string `json:"id"`
    VoiceFile_id string `json:"voice_file_id"`
    Title string `json:"title"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}

type InlineQueryResultCachedAudio struct {
    Type string `json:"type"`
    Id string `json:"id"`
    AudioFile_id string `json:"audio_file_id"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent interface{} `json:"input_message_content,omitempty"`

}


type InputTextMessageContent struct {
    MessageText string `json:"message_text"`
    ParseMode string `json:"parse_mode,omitempty"`
    DisableWebPagePreview bool `json:"disable_web_page_previe,omitempty"`

}

type InputLocationMessageContent struct {
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

type InputVenueMessageContent struct {
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Title string `json:"title"`
    Address string `json:"address"`
    FoursquareId string `json:"foursquare_id,omitempty"`
}

type InputContactMessageContent struct {
    PhoneNumber string `json:"phone_number"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`

}

type ChosenInlineResult struct {
    ResultId string `json:"result_id"`
    From *User `json:"from"`
    Location *Location `json:location,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    Query string `json:"query"`
}



//InputFile
