package models

type VWFirm struct {
	ID      int64   `json:"id"`
	LogoImg *string `json:"logo_img"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
}

type Region struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
	DESC string `json:"desc"`
}

type BizSize struct {
	ID    int16  `json:"id"`
	TYPE  string `json:"type"`
	BRIEF string `type:"brief"`
}

type RegisterForm struct {
	EMAIL    string `form:"email"`
	FULLNAME string `form:"fullname"`
	PASS     string `form:"password"`
}
