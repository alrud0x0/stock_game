#-*- coding:utf-8 -*-

import requests
import json
import sys
import datetime
import FinanceDataReader as fdr


def get_games():
    url = 'http://localhost:9000/api/v1/games'
    response = requests.get(url)

    return json.loads(response.text)

def create_new_game():
    today = datetime.datetime.now()
    start_day = today + datetime.timedelta(days=3)
    start_day = start_day.strftime("%Y-%m-%d") + 'T09:00:00.000Z'
    url = 'http://localhost:9000/api/v1/games'
    response = requests.post(url, headers={'Content-Type': 'application/json'}, data=json.dumps({'date': start_day}))
    print(response.text)

def deactive_game(game):
    url = 'http://localhost:9000/api/v1/games/%d' % game['id']
    response = requests.put(url,
         headers={'Content-Type': 'application/json'},
         data=json.dumps({'date': game['date'], 'active': 0}))

    print(response.text)

#######

def close_game():
    # get all games
    games = get_games()
    last_game = games[0]

    # TODO
    # calculate game result

def close_vote():
    # get all games
    games = get_games()
    last_game = games[0]

    if last_game['active'] == 1:
        deactive_game(last_game)