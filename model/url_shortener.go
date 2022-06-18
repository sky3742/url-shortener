package model

type Url_Shortener struct {
	Id            string `json:"id"`
	Destination   string `json:"destination"`
	RedirectCount int    `json:"redirect_count"`
	CreatedAt     int64  `json:"created_at"`
}

func (url *Url_Shortener) IncrementCount() {
	url.RedirectCount++
}
