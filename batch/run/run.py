from apscheduler.schedulers.blocking import BlockingScheduler

# 실행할 함수
def exec_interval():
    print('exec interval')

def exec_cron():
    print('exec cron')

sched = BlockingScheduler()
# every day 18:00
sched.add_job(exec_cron, 'cron', day_of_week='4', hour=18, minute=00)

# 스케줄링 시작
sched.start()

from apscheduler.schedulers.blocking import BlockingScheduler

from batch import update_stock, update_game

sched = BlockingScheduler()

# every day 18:00
sched.add_job(update_stock.run, 'cron', hour=18, minute=00)

# every friday 17:50
sched.add_job(update_game.close_game, 'cron', day_of_week='4', hour=17, minute=50)

# every friday 18:00
sched.add_job(update_game.create_new_game, 'cron', day_of_week='4', hour=18, minute=00)

# every friday 18:00
sched.add_job(update_game.close_vote, 'cron', day_of_week='0', hour=9, minute=00)

# 스케줄링 시작
sched.start()