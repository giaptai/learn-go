package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VWFirm struct {
	ID      int64   `json:"id"`
	LogoImg *string `json:"logo_img"`
	Name    string  `json:"name" form:"company"`
	Region  string  `json:"region" form:"region"`
}

type Firm struct {
	ID      int64   `json:"id"`
	Firm    string  `json:"firm" form:"company"`
	LogoImg *string `json:"logo_img" form:"logo_img"`
	Region  string  `json:"region" form:"region_id"`
	BizSize string  `json:"biz_size" form:"biz_size"`
	Website string  `json:"website" form:"website"`
	Address string  `json:"address" form:"address"`
	Holding string  `json:"holding" form:"holding"`
}

type Region struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
	DESC string `json:"desc"`
}

type BizSize struct {
	ID    int16  `json:"id"`
	TYPE  string `json:"type"`
	BRIEF string `json:"brief"`
}

type RegisterForm struct {
	EMAIL    string `form:"email"`
	FULLNAME string `form:"fullname"`
	PASS     string `form:"password"`
}

type Opinion struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FIRMID     int                `form:"firm_id" bson:"firm_id" json:"firm_id"`
	HEADER     string             `form:"header" bson:"header" json:"header"`
	LEVEL      string             `form:"level" bson:"level" json:"level"`
	CONTENT    string             `form:"content" bson:"content" json:"content"`
	TOTALSCORE float32            `form:"total_score" bson:"total_score" json:"total_score"`
	CREATE_AT  time.Time          `bson:"create_at" json:"create_at"`
}

func (o *Opinion) DefaultOpinion() Opinion {
	*o = Opinion{
		ID:         o.ID,
		FIRMID:     o.FIRMID,
		HEADER:     o.HEADER,
		LEVEL:      o.LEVEL,
		CONTENT:    o.CONTENT,
		TOTALSCORE: o.TOTALSCORE,
		CREATE_AT:  time.Now().UTC().Add(7 * time.Hour),
	}
	return *o
}
