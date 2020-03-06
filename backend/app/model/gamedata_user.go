package model

type GameDataUser struct {
	Id            int      `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	GameDateId    int      `gorm:"column:gamedataId"           json:"gamedataId"`
	UserId        int      `gorm:"column:userId"               json:"userId"`
	Reason        string   `gorm:"column:reason"               json:"reason"`
}

type GameDataUserDisplay struct {
	Id	int 		`json:"id"`
	User 	UserDisplay     `json:"user"`
	Reason  string          `json:"reason"`
}

func GetGameDataUserTable() string {
	return "gamedata_user"
}

func GetAllGameDataUsers() ([]GameDataUser, error) {
	gamedataUsers := []GameDataUser{}
	err := DB().Table(GetGameDataUserTable()).Find(&gamedataUsers).Error
	if err != nil {
		return gamedataUsers, err
	}
	return gamedataUsers, nil
}

func GetGameDatasUserByGameDataId(gamedataId int) ([]GameDataUser, error) {
	gamedataUsers := []GameDataUser{}
	err := DB().Table(GetGameDataUserTable()).Where("gamedataId = ?", gamedataId).Find(&gamedataUsers).Error
	if err != nil {
		return gamedataUsers, err
	}
	return gamedataUsers, nil
}

func GetGameDataUser(id int) (GameDataUser, error) {
	gamedataUser := GameDataUser{}
	err := DB().Table(GetGameDataUserTable()).Where("id = ?", id).Find(&gamedataUser).Error
	if err != nil {
		return gamedataUser, err
	}
	return gamedataUser, nil
}

func AddGameDataUser(gamedataId int, userId int, reason string) (GameDataUser, error) {
	gamedataUser := GameDataUser{GameDateId: gamedataId, UserId: userId, Reason: reason}
	err := DB().Table(GetGameDataUserTable()).Create(&gamedataUser).Error

	return gamedataUser, err
}

func UpdateGameDataUser(id int, reason string) (GameDataUser, error) {
	gamedataUser, err := GetGameDataUser(id)
	if err != nil {
		return gamedataUser, err
	}

	err = DB().Table(GetGameDataTable()).Where("id = ?", id).Update("reason", reason).Error
	gamedataUser.Reason = reason
	return gamedataUser, err
}

func DeleteGameDataUser(id int) (GameDataUser, error) {
	gamedataUser := GameDataUser{}
	err := DB().Table(GetGameDataTable()).Where("id = ?", id).Delete(&gamedataUser).Error

	return gamedataUser, err
}
