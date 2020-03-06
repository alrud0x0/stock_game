package model

type UserGameData struct {
	Id            	int      `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	UserId          int      `gorm:"column:userId"               json:"userId"`
	GameId          int      `gorm:"column:gameId"               json:"gameId"`
	GameDataId      int      `gorm:"column:gamedataId"           json:"gamedataId"`
}

func GetUserGameDataTable() string {
	return "user_gamedata"
}

func GetAllUserGameDatas() ([]UserGameData, error) {
	userGamedataList := []UserGameData{}
	err := DB().Table(GetUserGameDataTable()).Find(&userGamedataList).Error
	if err != nil {
		return userGamedataList, err
	}
	return userGamedataList, nil
}

func GetUserGameDatasByUserId(userId int) ([]UserGameData, error) {
	userGamedataList := []UserGameData{}
	err := DB().Table(GetUserGameDataTable()).Where("userId = ?", userId).Find(&userGamedataList).Error
	if err != nil {
		return userGamedataList, err
	}
	return userGamedataList, nil
}

func GetUserGameDatasByUserIdGameId(userId int, gameId int) ([]UserGameData, error) {
	userGamedataList := []UserGameData{}
	err := DB().Table(GetUserGameDataTable()).Where("userId = ? AND gameId = ?", userId, gameId).Find(&userGamedataList).Error
	if err != nil {
		return userGamedataList, err
	}
	return userGamedataList, nil
}

func AddUserGameData(userId int, gameId int, gamedataId int) (UserGameData, error) {
	userGamedata := UserGameData{UserId: userId, GameId: gameId, GameDataId: gamedataId}
	err := DB().Table(GetUserGameDataTable()).Create(&userGamedata).Error

	return userGamedata, err
}

func DeleteUserGameData(userId int, gameId int, stockId int) (UserGameData, error) {
	userGamedata := UserGameData{}

	gamedata, err := GetGameDataByGameStock(gameId, stockId)
	if err != nil {
		return userGamedata, err
	}

	err = DB().Table(GetUserGameDataTable()).Where("userId = ? AND gamedataId = ?", userId, gamedata.Id).Delete(&userGamedata).Error

	return userGamedata, err
}
