package telegrambot

import (
    "os"
    "errors"
    "encoding/json"
)

func (api *API) GetMe() (*User, error) {
    var ret User
    err := api.callGetMethod("getMe", &ret)
    return &ret, err
}

func (api *API) SendMessage(sm *SendMessageObj) (*Message, error) {
    var ret Message
    err := api.callPostMethod("sendMessage", sm, &ret)
    return &ret, err
}

func (api *API) ForwardMessage(fm *ForwardMessageObj) (*Message, error) {
    var ret Message
    err := api.callPostMethod("ForwardMessage", fm, &ret)
    return &ret, err
}

func (api *API) SendPhoto(ph *SendPhotoObj) (*Message, error) {
    var err error
    var ret Message
    switch ph.Photo.(type) {
        default:
            return nil, errors.New("Photo should be string or file")
        case string:
            err = api.callPostMethod("sendPhoto", ph, &ret)
        case *os.File:
            params := make(map[string]string)
            params["chat_id"] = ph.ChatId
            params["caption"] = ph.Caption
            if ph.DisableNotification == true {
                params["disable_notification"] = "True"
            }
            if ph.DisableNotification == false {
                params["disable_notification"] = "False"
            }
            if ph.ReplyToMessageId > 0{
                params["reply_to_message_id"] = string(ph.ReplyToMessageId)
            }

            if ph.ReplyMarkup != nil {
                markup, err := json.Marshal(ph.ReplyMarkup)
                if err != nil {
                    return nil, errors.New("bad reply markup")
                }
                params["reply_markup"] = string(markup)
            }
            err = api.callPostMultipartMethod("sendPhoto", params, ph.Photo.(*os.File), "photo", &ret)
            if err != nil {
                return nil, err
            }
    }
    return &ret, nil
}

func (api *API) SendAudio(au *SendAudioObj) (*Message, error) {
    var err error
    var ret Message
    switch au.Audio.(type) {
        default:
            return nil, errors.New("Audio should be string or file")
        case string:
            err = api.callPostMethod("sendAudio", au, &ret)
        case *os.File:
            params := make(map[string]string)
            params["chat_id"] = au.ChatId
            if au.Duration > 0 {
                params["duration"] = string(au.Duration)
            }
            if au.Performer != "" {
                params["performer"] = au.Performer
            }
            if au.Title != ""{
                params["title"] = au.Title
            }
            if au.DisableNotification == true {
                params["disable_notification"] = "True"
            }
            if au.DisableNotification == false {
                params["disable_notification"] = "False"
            }
            if au.ReplyToMessageId > 0{
                params["reply_to_message_id"] = string(au.ReplyToMessageId)
            }

            if au.ReplyMarkup != nil {
                markup, err := json.Marshal(au.ReplyMarkup)
                if err != nil {
                    return nil, errors.New("bad reply markup")
                }
                params["reply_markup"] = string(markup)
            }
            err = api.callPostMultipartMethod("sendAudio", params, au.Audio.(*os.File), "audio", &ret)
            if err != nil {
                return nil, err
            }
    }
    return &ret, nil
}

func (api *API) SendDocument(doc *SendDocumentObj) (*Message, error) {
    var err error
    var ret Message
    switch doc.Document.(type) {
        default:
            return nil, errors.New("Document should be string or file")
        case string:
            err = api.callPostMethod("sendDocument", doc, &ret)
        case *os.File:
            params := make(map[string]string)
            params["chat_id"] = doc.ChatId
            if doc.Caption != ""{
                params["caption"] = doc.Caption
            }
            if doc.DisableNotification == true {
                params["disable_notification"] = "True"
            }
            if doc.DisableNotification == false {
                params["disable_notification"] = "False"
            }
            if doc.ReplyToMessageId > 0{
                params["reply_to_message_id"] = string(doc.ReplyToMessageId)
            }

            if doc.ReplyMarkup != nil {
                markup, err := json.Marshal(doc.ReplyMarkup)
                if err != nil {
                    return nil, errors.New("bad reply markup")
                }
                params["reply_markup"] = string(markup)
            }
            err = api.callPostMultipartMethod("sendDocument", params, doc.Document.(*os.File), "document", &ret)
            if err != nil {
                return nil, err
            }
    }
    return &ret, nil
}

func (api *API) SendSticker(st *SendStickerObj) (*Message, error) {
    var err error
    var ret Message
    switch st.Sticker.(type) {
        default:
            return nil, errors.New("Sticker should be string or file")
        case string:
            err = api.callPostMethod("sendSticker", st, &ret)
        case *os.File:
            params := make(map[string]string)
            params["chat_id"] = st.ChatId
            if st.Caption != ""{
                params["caption"] = st.Caption
            }
            if st.DisableNotification == true {
                params["disable_notification"] = "True"
            }
            if st.DisableNotification == false {
                params["disable_notification"] = "False"
            }
            if st.ReplyToMessageId > 0{
                params["reply_to_message_id"] = string(st.ReplyToMessageId)
            }

            if st.ReplyMarkup != nil {
                markup, err := json.Marshal(st.ReplyMarkup)
                if err != nil {
                    return nil, errors.New("bad reply markup")
                }
                params["reply_markup"] = string(markup)
            }
            err = api.callPostMultipartMethod("sendSticker", params, st.Sticker.(*os.File), "sticker", &ret)
            if err != nil {
                return nil, err
            }
    }
    return &ret, nil
}

func (api *API) SendVideo(vid *SendVideoObj) (*Message, error) {
    var err error
    var ret Message
    switch vid.Video.(type) {
        default:
            return nil, errors.New("Video should be vidring or file")
        case string:
            err = api.callPostMethod("sendVideo", vid, &ret)
        case *os.File:
            params := make(map[string]string)
            params["chat_id"] = vid.ChatId
            if vid.Caption != ""{
                params["caption"] = vid.Caption
            }
            if vid.DisableNotification == true {
                params["disable_notification"] = "True"
            }
            if vid.DisableNotification == false {
                params["disable_notification"] = "False"
            }
            if vid.ReplyToMessageId > 0{
                params["reply_to_message_id"] = string(vid.ReplyToMessageId)
            }
            if vid.Duration > 0{
                params["duration"] = string(vid.Duration)
            }
            if vid.Width > 0{
                params["width"] = string(vid.Width)
            }
            if vid.Height > 0{
                params["height"] = string(vid.Height)
            }

            if vid.ReplyMarkup != nil {
                markup, err := json.Marshal(vid.ReplyMarkup)
                if err != nil {
                    return nil, errors.New("bad reply markup")
                }
                params["reply_markup"] = string(markup)
            }
            err = api.callPostMultipartMethod("sendVideo", params, vid.Video.(*os.File), "video", &ret)
            if err != nil {
                return nil, err
            }
    }
    return &ret, nil
}

func (api *API) SendVoice(voice *SendVoiceObj) (*Message, error) {
    var err error
    var ret Message
    switch voice.Voice.(type) {
        default:
            return nil, errors.New("Voice should be voicering or file")
        case string:
            err = api.callPostMethod("sendVoice", voice, &ret)
        case *os.File:
            params := make(map[string]string)
            params["chat_id"] = voice.ChatId
            if voice.DisableNotification == true {
                params["disable_notification"] = "True"
            }
            if voice.DisableNotification == false {
                params["disable_notification"] = "False"
            }
            if voice.ReplyToMessageId > 0{
                params["reply_to_message_id"] = string(voice.ReplyToMessageId)
            }
            if voice.Duration > 0{
                params["duration"] = string(voice.Duration)
            }

            if voice.ReplyMarkup != nil {
                markup, err := json.Marshal(voice.ReplyMarkup)
                if err != nil {
                    return nil, errors.New("bad reply markup")
                }
                params["reply_markup"] = string(markup)
            }
            err = api.callPostMultipartMethod("sendVoice", params, voice.Voice.(*os.File), "voice", &ret)
            if err != nil {
                return nil, err
            }
    }
    return &ret, nil
}

func (api *API) SendLocation(loc *SendLocationObj) (*Message, error) {
    var ret Message
    err := api.callPostMethod("sendLocation", loc, &ret)
    return &ret, err
}

func (api *API) SendVenue(ven *SendVenueObj) (*Message, error) {
    var ret Message
    err := api.callPostMethod("sendVenue", ven, &ret)
    return &ret, err
}

func (api *API) SendContact(con *SendContactObj) (*Message, error) {
    var ret Message
    err := api.callPostMethod("sendContact", con, &ret)
    return &ret, err
}

func (api *API) SendChatActionObj(ca *SendChatActionObj) (interface{}, error) {
    var ret interface{}
    var err error
    switch ca.Action {
        case
            "typing",
            "upload_photo",
            "record_video",
            "upload_video",
            "record_audio",
            "upload_audio",
            "upload_document",
            "find_location":  err = api.callPostMethod("sendChatAction", ca, &ret)
        default:
            err = errors.New("chat Action not allowed")
    }

       return ret, err
}

func (api *API) GetUserProfilePhotos(upp *GetUserProfilePhotosObj) (*UserProfilePhotos, error) {
    var ret UserProfilePhotos
    err := api.callPostMethod("getUserProfilePhotos", upp, &ret)
    return &ret, err
}

func (api *API) GetFile(f *GetFileObj) (*File, error) {
    var ret File
    err := api.callPostMethod("getFile", f, &ret)
    return &ret, err
}

func (api *API) KickChatMember(kcm *KickChatMemberObj) (*bool, error) {
    var ret bool
    err := api.callPostMethod("kickChatMember", kcm, &ret)
    return &ret, err
}

func (api *API) LeaveChat(lc *LeaveChatObj) (*bool, error) {
    var ret bool
    err := api.callPostMethod("leaveChat", ret, &ret)
    return &ret, err
}

func (api *API) UnbanChatMember(ucm *UnbanChatMemberObj) (*bool, error) {
    var ret bool
    err := api.callPostMethod("unbanChatMember", ucm, &ret)
    return &ret, err
}

func (api *API) GetChat(c *GetChatObj) (*Chat, error) {
    var ret Chat
    err := api.callPostMethod("getChat", c, &ret)
    return &ret, err
}

func (api *API) GetChatAdministrators(ca *GetChatAdministratorsObj) (*[]ChatMember, error) {
    var ret []ChatMember
    err := api.callPostMethod("getChatAdministrators", ca, &ret)
    return &ret, err
}

func (api *API) GetChatMembersCount(cmc *GetChatMembersCountObj) (*int64, error) {
    var ret int64
    err := api.callPostMethod("getChatMembersCount", cmc, &ret)
    return &ret, err
}

func (api *API) GetChatMember(cm *GetChatMemberObj) (*ChatMember, error) {
    var ret ChatMember
    err := api.callPostMethod("getChatMember", cm, &ret)
    return &ret, err
}

func (api *API) AnswerCallbackQuery(qc *AnswerCallbackQueryObj) (*bool, error) {
    var ret bool
    err := api.callPostMethod("answerCallbackQuery", qc, &ret)
    return &ret, err
}

//#TODO test this function for ret values
func (api *API) EditMessageText(mt *EditMessageTextObj) (interface{}, error) {
    var ret interface{}
    err := api.callPostMethod("editMessageText", mt, &ret)
    return &ret, err
}

//#TODO test this function for ret values
func (api *API) EditMessageCaption(mc *EditMessageCaptionObj) (interface{}, error) {
    var ret interface{}
    err := api.callPostMethod("editMessageCaption", mc, &ret)
    return &ret, err
}

//#TODO test this function for ret values
func (api *API) EditMessageReplyMarkup(rm *EditMessageReplyMarkupObj) (interface{}, error) {
    var ret interface{}
    err := api.callPostMethod("editMessageReplyMarkup", rm, &ret)
    return &ret, err
}


//#TODO test this
func (api *API) AnswerInlineQuery(iq *AnswerCallbackQueryObj) (*bool, error) {
    var ret bool
    err := api.callPostMethod("answerCallbackQuery", iq, &ret)
    return &ret, err
}

func (api *API) GetUpdates(u *UpdateObj) (*[]Update, error) {
    var ret []Update
    err := api.callPostMethod("getUpdates", u, &ret)
    return &ret, err
}
