package model

import "time"

type Stock struct {
	Id        	int       `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Code      	string    `gorm:"column:code"                 json:"code"`
	Name      	string    `gorm:"column:name"   	      json:"name"`
	Price     	int       `gorm:"column:price"                json:"price"`
	LastUpdate	time.Time `gorm:"column:lastUpdate"           json:"lastUpdate"`
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

func GetStockByCode(code string) (Stock, error) {
	stock := Stock{}
	err := DB().Table(GetStockTable()).Where("uid = ?", code).Find(&stock).Error
	if err != nil {
		return stock, err
	}
	return stock, nil
}

func AddStock(code string, name string, price int) (Stock, error) {
	stock := Stock{Code: code, Name: name, Price: price, LastUpdate: time.Now()}
	err := DB().Table(GetStockTable()).Create(&stock).Error

	return stock, err
}

func UpdateStock(id int, price int) (Stock, error) {
	stock := Stock{}
	stock, err := GetStock(id)
	if err != nil {
		return stock, err
	}

	err = DB().Table(GetStockTable()).Where("id = ?", id).Update("price", price).Error
	err = DB().Table(GetStockTable()).Where("id = ?", id).Update("lastUpdate", time.Now()).Error
	stock.Price = price

	return stock, err
}

func DeleteStock(id int) (Stock, error) {
	stock := Stock{}
	err := DB().Table(GetStockTable()).Where("id = ?", id).Delete(&stock).Error

	return stock, err
}
