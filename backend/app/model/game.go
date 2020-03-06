package model

import (
	"time"
)

type Game struct {
	Id        int       	`gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Date      time.Time     `gorm:"column:date"                 json:"date"`
	Active    int           `gorm:"column:active"               json:"active"`
}

func GetGameTable() string {
	return "game"
}

func GetAllGames() ([]Game, error) {
	games := []Game{}
	err := DB().Table(GetGameTable()).Order("date desc").Find(&games).Error
	if err != nil {
		return games, err
	}
	return games, nil
}

func GetGame(id int) (Game, error) {
	game := Game{}
	err := DB().Table(GetGameTable()).Where("id = ?", id).Find(&game).Error
	if err != nil {
		return game, err
	}
	return game, nil
}

func AddGame(date time.Time) (Game, error) {
	game := Game{Date: date, Active: 1}
	err := DB().Table(GetGameTable()).Create(&game).Error

	return game, err
}

func UpdateGame(id int, date time.Time, active int) (Game, error) {
	game := Game{}
	game, err := GetGame(id)
	if err != nil {
		return game, err
	}

	err = DB().Table(GetGameTable()).Where("id = ?", id).Update("date", date).Error
	err = DB().Table(GetGameTable()).Where("id = ?", id).Update("active", active).Error

	game.Date = date
	game.Active = active

	return game, err
}

func DeleteGame(id int) (Game, error) {
	game := Game{}
	err := DB().Table(GetGameTable()).Where("id = ?", id).Delete(&game).Error

	return game, err
}
