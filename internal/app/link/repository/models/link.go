package models

type Link struct {
	ID    int    `db:"id"`
	Link  string `db:"link"`
	Token string `db:"token"`
}
