package controllers

import (
	"github.com/revel/revel"
	"strconv"

	"backend/app/model"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func GetUserByParam(userIdString string) (model.User, error) {
	userId, err := strconv.Atoi(userIdString)

	user, err := model.GetUser(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetGameByParam(gameIdString string) (model.Game, error) {
	gameId, err := strconv.Atoi(gameIdString)
	game, err := model.GetGame(gameId)
	if err != nil {
		return game, err
	}

	return game, nil
}

func GetStockByParam(stockId int) (model.Stock, error) {
	stock, err := model.GetStock(stockId)

	if err != nil {
		return stock, err
	}

	return stock, nil
}

func ConvertGameStocksResponse(gdList []model.GameData) (model.GameStocksResponse, error) {
	gsr := model.GameStocksResponse{}
	stockKeyUsers := make(map[int][]int)
	stockKeyStartPrice := make(map[int]int)
	userStockReasons := make(map[int]map[int]string)

	for i, gamedata := range gdList {
		if i == 0 {
			game := model.Game{}
			game, err := model.GetGame(gamedata.GameId)
			if err == nil {
				gsr.Game = game
			}
		}
		_, ok := stockKeyUsers[gamedata.StockId]
		if ok {
			userList := stockKeyUsers[gamedata.StockId]
			userList = append(userList, gamedata.UserId)

			stockKeyUsers[gamedata.StockId] = userList
		} else {
			stockKeyUsers[gamedata.StockId] = []int{gamedata.UserId}
			stockKeyStartPrice[gamedata.StockId] = gamedata.StartPrice
		}


		_, exist := userStockReasons[gamedata.UserId]
		if exist {
			userStockReasons[gamedata.UserId][gamedata.StockId] = gamedata.Reason
		} else {
			userStockReasons[gamedata.UserId] = map[int]string{gamedata.StockId: gamedata.Reason}
		}
	}

	for stockId, stockUsers := range stockKeyUsers {
		stock, _ := model.GetStock(stockId)
		users := []model.StockUser{}
		for _, userId := range stockUsers {
			user := model.User{}
			user, err := model.GetUser(userId)

			reason := userStockReasons[user.Id][stockId]
			if err == nil {
				users = append(users, model.StockUser{User: model.UserDisplay{Id: user.Id, Name: user.Name, Email: user.Email}, Reason: reason})
			}
		}

		startPrice, _ := stockKeyStartPrice[stockId]

		gamestocks := model.GameStocks{}
		gamestocks.Stock = stock
		gamestocks.Users = users
		gamestocks.StartPrice = startPrice

		gsr.Stocks = append(gsr.Stocks, gamestocks)
	}

	return gsr, nil
}
