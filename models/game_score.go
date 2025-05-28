package models

import "gorm.io/gorm"

type GameScore struct {
	gorm.Model
	Username string `json:"username"`
	Merges   int    `json:"merges"`
	Duration int    `json:"duration"`
	Reward   int    `json:"reward"`
}
