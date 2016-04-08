package telegrambot

type API struct {
    Token string
}

func CreateAPI(token string) *API {
    return &API{Token: token}
}

