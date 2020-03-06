#-*- coding:utf-8 -*-

import requests
import json
import sys
import datetime
import FinanceDataReader as fdr

import warnings
warnings.simplefilter(action = "ignore", category = FutureWarning)


def get_stocks():
    url = 'http://localhost:9000/api/v1/stocks'
    resp = requests.get(url)
    if resp.status_code == 200:
        stock_list = json.loads(resp.text)
        stock_dicts = {}
        for stock in stock_list:
            stock_dicts[stock['code']] = stock['id']

        return stock_dicts

    return []

def get_krx_stocks():
    stockmaps = {}

    stocks = fdr.StockListing('KRX')
    for stock in stocks.values:
        stockmaps[stock[0]] = stock[1]

    return stockmaps

def compare_stocks(old_dicts, new_dicts):
    codes = list(old_dicts.keys())
    while codes:
        code = codes.pop()
        if code in new_dicts:
            del new_dicts[code]
        else:
            del old_dicts[code]

    return old_dicts, new_dicts

def get_price(code):
    now = datetime.datetime.now()
    today = now.strftime("%Y-%m-%d")
    df = fdr.DataReader(code, today, today)['Close']

    return df.values[0]

def update_stock(id, price):
    url = 'http://localhost:9000/api/v1/stocks/%d' % id
    response = requests.put(url, headers={'Content-Type': 'application/json'}, data=json.dumps({'price': int(price)}))

    return response

def create_stock(code, name, price):
    url = 'http://localhost:9000/api/v1/stocks'
    response = requests.post(url, headers={'Content-Type': 'application/json'}, data=json.dumps({'code': str(code), 'name': name, 'price': int(price)}))

    return response


#################

def run():
    # fetch from server
    stocks = get_stocks()
    if len(stocks.keys()) == 0:
        print('fail to get stocks')
        sys.exit(1)

    # fetch from krx html
    krx_stocks = get_krx_stocks()

    # filter
    exist_dicts, new_stocks_dicts = compare_stocks(stocks, krx_stocks)

    # update
    for stock_code in exist_dicts:
        stock_id = exist_dicts[stock_code]
        stock_price = get_price(stock_code)
        resp = update_stock(stock_id, stock_price)
        # if resp.status_code != 200:
        print(resp.text)

    # create
    for stock_code in new_stocks_dicts:
        stock_name = new_stocks_dicts[stock_code]
        stock_price = get_price(stock_code)
        resp = create_stock(stock_code, stock_name, stock_price)
        if resp.status_code != 200:
            print(resp.text)

    # TODO
    # update