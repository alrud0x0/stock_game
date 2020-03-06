package model

type GameData struct {
	Id            int      `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	GameId        int      `gorm:"column:gameId"               json:"gameId"`
	StockId       int      `gorm:"column:stockId"              json:"stockId"`
	StockCode     string   `gorm:"column:stockCode"            json:"code"`
	StockName     string   `gorm:"column:stockName"            json:"name"`
	Start         int      `gorm:"column:start"                json:"start"`
	Current       int      `gorm:"column:current"              json:"current"`
	Vote          int      `gorm:"column:vote"                 json:"vote"`
	Rank          int      `gorm:"column:rank"                 json:"rank"`
}

type GameDataResponse struct {
	Code	  string    `json:"code"`
	Name	  string    `json:"name"`
	Start	  int 	    `json:"start"`
	Current	  int 	    `json:"start"`
	Vote	  int 	    `json:"vote"`
	Rank      int       `json:"rank"`
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
	err := DB().Table(GetGameDataTable()).Where("gameId = ?", gameId).Find(&gameDatas).Order("vote desc").Error
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

func GetGameDataByGameStock(gameId int, stockId int) (GameData, error) {
	gameData := GameData{}
	err := DB().Table(GetGameDataTable()).Where("gameId = ? AND stockId = ?", gameId, stockId).Find(&gameData).Error
	if err != nil {
		return gameData, err
	}
	return gameData, nil
}

func AddGameData(gameId int, stockId int, stockCode string, stockName string, start int) (GameData, error) {
	gameData := GameData{GameId: gameId, StockId: stockId, StockCode: stockCode, StockName: stockName, Start: start, Current: start, Vote: 1, Rank: 0}
	err := DB().Table(GetGameDataTable()).Create(&gameData).Error

	return gameData, err
}

func UpdateGameData(id int, change int) (GameData, error) {
	gameData := GameData{}
	gameData, err := GetGameData(id)
	if err != nil {
		return gameData, err
	}

	vote := 0
	if change > 0 {
		vote = gameData.Vote + 1
	} else {
		vote = gameData.Vote - 1
	}

	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("vote", vote).Error
	gameData.Vote = vote

	return gameData, err
}

func DeleteGameData(id int) (GameData, error) {
	gameData := GameData{}
	err := DB().Table(GetGameDataTable()).Where("id = ?", id).Delete(&gameData).Error

	return gameData, err
}
