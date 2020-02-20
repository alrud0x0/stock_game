package model

type Stock struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Uid       string    `gorm:"column:uid"                  json:"uid"`
	Name      string    `gorm:"column:name"   	        json:"name"`
	Price     float64   `gorm:"column:price"                json:"price"`
}

func GetStockTable() string {
	return "stock"
}

func GetAllStocks() ([]Stock, error) {
	stocks := []Stock{}
	err := DB().Table(GetStockTable()).Find(&stocks).Error
	if err != nil {
		return stocks, err
	}
	return stocks, nil
}

func GetStock(id int) (Stock, error) {
	stock := Stock{}
	err := DB().Table(GetStockTable()).Where("id = ?", id).Find(&stock).Error
	if err != nil {
		return stock, err
	}
	return stock, nil
}

func GetStockByUid(uid string) (Stock, error) {
	stock := Stock{}
	err := DB().Table(GetStockTable()).Where("uid = ?", uid).Find(&stock).Error
	if err != nil {
		return stock, err
	}
	return stock, nil
}

func AddStock(uid string, name string, price float64) (Stock, error) {
	stock := Stock{Uid: uid, Name: name, Price: price}
	err := DB().Table(GetStockTable()).Create(&stock).Error

	return stock, err
}

func UpdateStock(id int, price float64) (Stock, error) {
	stock := Stock{}
	stock, err := GetStock(id)
	if err != nil {
		return stock, err
	}

	err = DB().Table(GetStockTable()).Where("id = ?", id).Update("price", price).Error
	stock.Price = price

	return stock, err
}

func DeleteStock(id int) (Stock, error) {
	stock := Stock{}
	err := DB().Table(GetStockTable()).Where("id = ?", id).Delete(&stock).Error

	return stock, err
}
