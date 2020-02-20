package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"strconv"
	"time"

	"backend/app/model"
)

type GameController struct {
	*revel.Controller
}

func (gCtrl GameController) Add() revel.Result {
	var jsonData map[string]interface{}
	gCtrl.Params.BindJSON(&jsonData)

	dateString := jsonData["date"].(string)
	date, err := time.Parse(time.RFC3339, dateString)

	game, err := model.AddGame(date)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return gCtrl.RenderText("Error inserting game into database!")
	} else {
		return gCtrl.RenderJSON(game)
	}
}

func (gCtrl GameController) List() revel.Result {
	games, err := model.GetAllGames()
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return gCtrl.RenderText("Error fetching game into database!")
	} else {
		return gCtrl.RenderJSON(games)
	}
}

func (gCtrl GameController) Get() revel.Result {
	game, err := GetGameByParam(gCtrl.Params.Get("gameId"))
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return gCtrl.RenderText("Error fetching game into database!")
	} else {
		return gCtrl.RenderJSON(game)
	}
}

func (gCtrl GameController) GetData() revel.Result {
	game, err := GetGameByParam(gCtrl.Params.Get("gameId"))

	gamedataList, _ := model.GetGameDatasByGameId(game.Id)
	response, _ := ConvertGameStocksResponse(gamedataList)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return gCtrl.RenderText("Error fetching game into database!")
	} else {
		return gCtrl.RenderJSON(response)
	}
}

func (gCtrl GameController) Update() revel.Result {
	gameIdString := gCtrl.Params.Get("gameId")
	gameId, err := strconv.Atoi(gameIdString)

	var jsonData map[string]interface{}
	gCtrl.Params.BindJSON(&jsonData)
	dateString := jsonData["date"].(string)
	date, err := time.Parse(time.RFC3339, dateString)

	game, err := model.UpdateGame(gameId, date)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return gCtrl.RenderText("Error update game into database!")
	} else {
		return gCtrl.RenderJSON(game)
	}
}

func (gCtrl GameController) Delete() revel.Result {
	gameIdString := gCtrl.Params.Get("gameId")
	gameId, err := strconv.Atoi(gameIdString)

	game, err := model.DeleteGame(gameId)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return gCtrl.RenderText("Error delete game into database!")
	} else {
		return gCtrl.RenderJSON(game)
	}
}

