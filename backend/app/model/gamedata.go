package model

type GameData struct {
	Id            int      `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	GameId        int      `gorm:"column:gameId"               json:"gameId"`
	UserId        int      `gorm:"column:userId"               json:"userId"`
	StockId       int      `gorm:"column:stockId"              json:"stockId"`
	StartPrice    int      `gorm:"column:startPrice"           json:"startPrice"`
	Reason        string   `gorm:"column:reason"               json:"reason"`
}

type GameStocksResponse struct {
	Id            int       	`json:"id"`
	Stocks        [] GameStocks     `json:"stocks"`
	Game          Game      	`json:"game"`
}

type GameStocks struct {
	Users         [] StockUser   `json:"users"`
	Stock         Stock     `json:"stock"`
	StartPrice    int       `json:"game"`
}

type GameDataByGameIdResponse struct {
	StockInfos        [] StockInfo     `json:"stocksinfos"`
	Game              Game      	   `json:"game"`
}

type StockInfo struct {
	Stock         Stock     	`json:"stock"`
	StartPrice    int       	`json:"startPrice"`
	StockUser     []StockUser 	`json:"stockusers"`
}

type StockUser struct {
	User	      UserDisplay	`json:"userdisplay"`
	Reason        string    	`json:"reason"`
}

func GetGameDataTable() string {
	return "gamedata"
}

func GetAllGameDatas() ([]GameData, error) {
	gameDatas := []GameData{}
	err := DB().Table(GetGameDataTable()).Find(&gameDatas).Error
	if err != nil {
		return gameDatas, err
	}
	return gameDatas, nil
}

func GetGameDatasByGameId(gameId int) ([]GameData, error) {
	gameDatas := []GameData{}
	err := DB().Table(GetGameDataTable()).Where("gameId = ?", gameId).Find(&gameDatas).Error
	if err != nil {
		return gameDatas, err
	}
	return gameDatas, nil
}

func GetGameData(id int) (GameData, error) {
	gameData := GameData{}
	err := DB().Table(GetGameDataTable()).Where("id = ?", id).Find(&gameData).Error
	if err != nil {
		return gameData, err
	}
	return gameData, nil
}

func AddGameData(gameId int, userId int, stockId int, startPrice int, reason string) (GameData, error) {
	gameData := GameData{GameId: gameId, UserId: userId, StockId: stockId, StartPrice: startPrice, Reason: reason}
	err := DB().Table(GetGameDataTable()).Create(&gameData).Error

	return gameData, err
}

func UpdateData(id int, gameId int, userId int, stockId int, startPrice int, reason string) (GameData, error) {
	gameData := GameData{}
	gameData, err := GetGameData(id)
	if err != nil {
		return gameData, err
	}

	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("gameId", gameId).Error
	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("userId", userId).Error
	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("stockId", stockId).Error
	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("startPrice", startPrice).Error
	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("reason", reason).Error

	gameData.GameId = gameId
	gameData.UserId = userId
	gameData.StockId = stockId
	gameData.StartPrice = startPrice

	return gameData, err
}

func DeleteGameData(id int) (GameData, error) {
	gameData := GameData{}
	err := DB().Table(GetGameDataTable()).Where("id = ?", id).Delete(&gameData).Error

	return gameData, err
}
