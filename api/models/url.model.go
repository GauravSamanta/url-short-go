package models

type Url struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	LongUrl string `json:"longUrl"`
	ShortId string `json:"shortId" gorm:"unique not null index"`
	Clicks  int    `json:"clicks" gorm:"default:0"`
}
