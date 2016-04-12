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

func (api *API) callGetMethod(method string) (*[]byte, error) {
    url := fmt.Sprintf(base_url, api.Token, method)
    resp, err :=  api.Client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return &data, nil
}

func (api *API) callPostMethod(method string, obj interface{}) (*[]byte, error) {
    url := fmt.Sprintf(base_url, api.Token, method)
    jsonStr, err := json.Marshal(obj)
    if err != nil {
        return nil, err
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    resp, err := api.Client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return &data, nil
}

func (api *API) callPostMultipartMethod(method string, params map[string]string, f *os.File, paramName string) (*[]byte, error) {
    url := fmt.Sprintf(base_url, api.Token, method)
    fileContents, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, err
    }
    fi, err := f.Stat()
    if err != nil {
        return nil, err
    }
    f.Close()
    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(paramName, fi.Name())
    if err != nil {
        return nil, err
    }
    part.Write(fileContents)
    for key, val := range params {
        writer.WriteField(key, val)
    }
    err = writer.Close()
    if err != nil {
        return nil, err
    }
    req, err := http.NewRequest("POST", url, body)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    resp, err := api.Client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return &data, nil
}
func CreateAPI(token string) *API {
    api := API{Token: token}
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    api.Client = &http.Client{Transport: tr}
    return &api
}

