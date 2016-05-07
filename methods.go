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
            params["chat_id"] = ph.ChatID
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
            params["chat_id"] = au.ChatID
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
            params["chat_id"] = doc.ChatID
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
            params["chat_id"] = st.ChatID
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

func (api *API) GetUpdates(u *UpdateObj) (*[]Update, error) {
    var ret []Update
    err := api.callPostMethod("getUpdates", u, &ret)
    return &ret, err
}
