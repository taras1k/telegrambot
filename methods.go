package telegrambot

import (
    "errors"
    "encoding/json"
)

func (api *API) GetMe() (*User, error) {
    data, err := api.callGetMethod("getMe")
    ret := new(User)
    err = json.Unmarshal(*data, &ret)
    if err != nil {
        return nil, err
    }
    return ret, nil
}

func (api *API) SendMessage(sm *SendMessageObj) (*Message, error) {
    data, err := api.callPostMethod("sendMessage", sm)
    ret := new(Message)
    err = json.Unmarshal(*data, &ret)
    if err != nil {
        return nil, err
    }
    return ret, nil
}

func (api *API) ForwardMessage(fm *ForwardMessageObj) (*Message, error) {
    data, err := api.callPostMethod("forwardMessage", fm)
    ret := new(Message)
    err = json.Unmarshal(*data, &ret)
    if err != nil {
        return nil, err
    }
    return ret, nil
}

func (api *API) SendPhoto(ph *SendPhotoObj) (*Message, error) {
    var data *[]byte
    var err error
    switch ph.Photo.(type) {
        default:
            return nil, errors.New("Photo should be string or file")
        case string:
            data, err = api.callPostMethod("sendPhoto", ph)
        //write case file
    }

    ret := new(Message)
    err = json.Unmarshal(*data, &ret)
    if err != nil {
        return nil, err
    }
    return ret, nil
}
