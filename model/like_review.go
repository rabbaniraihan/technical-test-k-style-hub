package model

type Like_review struct {
	Id       int `gorm:"primaryKey" json:"id"`
	ReviewId int `json:"id_review"`
	MemberId int `json:"id_member"`
}
