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

type ChosenInlineResult struct {
    ResultId string `json:"result_id"`
    From *User `json:"from"`
    Query string `json:"query"`
}



//InputFile
