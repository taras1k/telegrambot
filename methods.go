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

func (api *API) GetUpdates(u *UpdateObj) (*[]Update, error) {
    var ret []Update
    err := api.callPostMethod("getUpdates", u, &ret)
    return &ret, err
}
