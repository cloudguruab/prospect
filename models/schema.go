package models

type Scores struct {
  ID uint `json:"id" gorm:"primary_key"`
  User string `json:"user"`
  Text string `json:"text"`
  Harmful string `json:"harmful"`
}
