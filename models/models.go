package models

type VWFirm struct {
	ID      int64   `json:"id"`
	LogoImg *string `json:"logo_img"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
}
