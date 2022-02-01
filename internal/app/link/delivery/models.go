package delivery

type Link struct {
	ID    int64  `json:"-"`
	Link  string `json:"link"`
	Token string `json:"-"`
}

type LinkFilter struct {
	Token *string `param:"token"`
}
