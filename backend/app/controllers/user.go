package controllers

import (
	"github.com/revel/revel"
	"fmt"
	
	"backend/app/model"
)

type UserController struct {
	*revel.Controller
}

func (uCtrl UserController) Add() revel.Result {
	var jsonData map[string]interface{}
	uCtrl.Params.BindJSON(&jsonData)

	name := jsonData["name"].(string)
	email := jsonData["email"].(string)
	password := jsonData["password"].(string)
	reward := int(jsonData["reward"].(float64))

	user, err := model.AddUser(name, email, password, reward)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Error inserting user into database!")
	} else {
		return uCtrl.RenderJSON(user)
	}
}

func (uCtrl UserController) Login() revel.Result {
	var jsonData map[string]string
	uCtrl.Params.BindJSON(&jsonData)

	name := jsonData["name"]
	password := jsonData["password"]

	user, err := model.LoginUser(name, password)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Invalid User")
	} else {
		return uCtrl.RenderJSON(user)
	}
}

func (uCtrl UserController) Update() revel.Result {
	userIdString := uCtrl.Params.Get("userId")
	user, err := GetUserByParam(userIdString)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Invalid User")
	}

	var jsonData map[string]interface{}
	uCtrl.Params.BindJSON(&jsonData)

	name := jsonData["name"].(string)
	email := jsonData["email"].(string)
	password := jsonData["password"].(string)
	reward := int(jsonData["reward"].(float64))

	user, err = model.UpdateUser(user.Id, name, email, password, reward)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Error update user into database!")
	} else {
		return uCtrl.RenderJSON(user)
	}
}

func (uCtrl UserController) JoinGame() revel.Result {
	user, err := GetUserByParam(uCtrl.Params.Get("userId"))
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Invalid User")
	}

	game, err := GetGameByParam(uCtrl.Params.Get("gameId"))
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Invalid Game")
	}

	var jsonData map[string]interface{}
	uCtrl.Params.BindJSON(&jsonData)
	stockId := int(jsonData["stockId"].(float64))
	reason := jsonData["reason"].(string)

	stock, err := GetStockById(stockId)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("Invalid Stock")
	}

	gamedata, err := model.GetGameDataByGameStock(game.Id, stock.Id)
	if err == nil {
		gamedata, err = model.UpdateGameData(gamedata.Id, 1)
		if err != nil {
			description := fmt.Sprint(err)
			revel.AppLog.Warn(description)
			return uCtrl.RenderText("fail to update gamedata: "+description)
		}
	} else {
		gamedata, err = model.AddGameData(game.Id, stockId, stock.Code, stock.Name, stock.Price)
		if err != nil {
			description := fmt.Sprint(err)
			revel.AppLog.Warn(description)
			return uCtrl.RenderText("fail to insert gamedata: "+description)
		}
	}

	gamedatauser, err := model.AddGameDataUser(gamedata.Id, user.Id, reason)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("fail to insert gamedata_user: "+description)
	}

	_, err = model.AddUserGameData(user.Id, game.Id, gamedata.Id)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return uCtrl.RenderText("fail to insert user_gamedata: "+description)
	} else {
		return uCtrl.RenderJSON(gamedatauser)
	}
}