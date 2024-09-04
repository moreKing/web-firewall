package model

type Home struct {
	Assets    int   `json:"assets"`
	Users     int   `json:"users"`
	Online    int   `json:"online"`
	Accounts  int   `json:"accounts"`
	Yesterday []int `json:"yesterday"`
	Today     []int `json:"today"`
	License   int   `json:"license"`
}
