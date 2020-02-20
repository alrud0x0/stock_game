package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"strconv"

	"backend/app/model"
)

type StockController struct {
	*revel.Controller
}

func (sCtrl StockController) Add() revel.Result {
	var jsonData map[string]interface{}
	sCtrl.Params.BindJSON(&jsonData)

	uid := jsonData["uid"].(string)
	name := jsonData["name"].(string)
	price := jsonData["price"].(float64)

	stock, err := model.AddStock(uid, name, price)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return sCtrl.RenderText("Error inserting stock into database!")
	} else {
		return sCtrl.RenderJSON(stock)
	}
}

func (sCtrl StockController) List() revel.Result {
	stocks, err := model.GetAllStocks()
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return sCtrl.RenderText("Error fetching stock into database!")
	} else {
		return sCtrl.RenderJSON(stocks)
	}
}

func (sCtrl StockController) Get() revel.Result {
	stockIdString := sCtrl.Params.Get("stockId")
	stockId, err := strconv.Atoi(stockIdString)

	stock, err := model.GetStock(stockId)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return sCtrl.RenderText("Error fetching stock into database!")
	} else {
		return sCtrl.RenderJSON(stock)
	}
}

func (sCtrl StockController) Update() revel.Result {
	stockIdString := sCtrl.Params.Get("stockId")
	stockId, err := strconv.Atoi(stockIdString)

	var jsonData map[string]interface{}
	sCtrl.Params.BindJSON(&jsonData)
	price := jsonData["price"].(float64)

	stock, err := model.UpdateStock(stockId, price)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return sCtrl.RenderText("Error update stock into database!")
	} else {
		return sCtrl.RenderJSON(stock)
	}
}

func (sCtrl StockController) Delete() revel.Result {
	stockIdString := sCtrl.Params.Get("stockId")
	stockId, err := strconv.Atoi(stockIdString)

	stock, err := model.DeleteStock(stockId)
	if err != nil {
		description := fmt.Sprint(err)
		revel.AppLog.Warn(description)
		return sCtrl.RenderText("Error delete stock into database!")
	} else {
		return sCtrl.RenderJSON(stock)
	}
}
