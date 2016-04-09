package telegrambot

import (
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
    data, err := api.callPostMethod("SendMessage", sm)
    ret := new(Message)
    err = json.Unmarshal(*data, &ret)
    if err != nil {
        return nil, err
    }
    return ret, nil

}
