package telegrambot

import (
    "os"
    "fmt"
    "bytes"
    "net/http"
    "io/ioutil"
    "crypto/tls"
    "encoding/json"
    "mime/multipart"
)

const (
    base_url string = "https://api.telegram.org/bot%s/%s"
)

type API struct {
    Token string
    Client *http.Client
}

type APIResult struct {
    Ok bool `json:"ok"`
    Result json.RawMessage `json:"result"`
    ErrorCode int64 `json:"error_code"`
    Description string `json:"description"`
}

type APIError struct {
    ErrorCode int64 `json:"error_code"`
    Description string `json:"description"`
}

func (e APIError) Error() string{
    return fmt.Sprintf("Code:%i Description:%s", e.ErrorCode, e.Description)
}

func (api *API) callGetMethod(method string, ret interface{})  error {
    url := fmt.Sprintf(base_url, api.Token, method)
    resp, err :=  api.Client.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    res := new(APIResult)
    err = json.Unmarshal(data, &res)
    if err != nil {
        return  err
    }
    if !res.Ok {
        return APIError{res.ErrorCode, res.Description}
    }
    err = json.Unmarshal(res.Result, ret)
    if err != nil {
        return  err
    }
    return nil
}

func (api *API) callPostMethod(method string, obj, ret interface{}) error {
    url := fmt.Sprintf(base_url, api.Token, method)
    jsonStr, err := json.Marshal(obj)
    if err != nil {
        return err
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    resp, err := api.Client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    res := new(APIResult)
    err = json.Unmarshal(data, &res)
    if err != nil {
        return  err
    }
    if !res.Ok {
        return APIError{res.ErrorCode, res.Description}
    }
    err = json.Unmarshal(res.Result, ret)
    if err != nil {
        return  err
    }
    return nil
}

func (api *API) callPostMultipartMethod(method string, params map[string]string, f *os.File, paramName string, ret interface{})  error {
    url := fmt.Sprintf(base_url, api.Token, method)
    fileContents, err := ioutil.ReadAll(f)
    if err != nil {
        return err
    }
    fi, err := f.Stat()
    if err != nil {
        return err
    }
    f.Close()
    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(paramName, fi.Name())
    if err != nil {
        return err
    }
    part.Write(fileContents)
    for key, val := range params {
        writer.WriteField(key, val)
    }
    err = writer.Close()
    if err != nil {
        return err
    }
    req, err := http.NewRequest("POST", url, body)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    resp, err := api.Client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    res := new(APIResult)
    err = json.Unmarshal(data, &res)
    if err != nil {
        return  err
    }
    if !res.Ok {
        return APIError{res.ErrorCode, res.Description}
    }
    err = json.Unmarshal(res.Result, ret)
    if err != nil {
        return  err
    }
    return  nil
}

func CreateAPI(token string) *API {
    api := API{Token: token}
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    api.Client = &http.Client{Transport: tr}
    return &api
}

