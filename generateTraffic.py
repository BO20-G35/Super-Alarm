#!/usr/bin/env python3
import requests
import sched, time

url_getCookie = 'http://127.0.0.1:8081/getCookie'
url_targetTraffic = 'http://127.0.0.1:8081/toggle?k=JHDGFUAYEG23RIUETYWERY3RSDFV23RGUE'

session = requests.Session()

# Sets JWT token as session cookie
session.get(url_getCookie)
print("got cookie")

# Trigger alarm toggle func() every minute
schedule = sched.scheduler(time.time, time.sleep)
def triggerTraffic(sc):
    response = session.get(url_targetTraffic)
    print(response.text)
    schedule.enter(60, 1, triggerTraffic, (sc,))

schedule.enter(1, 1, triggerTraffic, (schedule,))
schedule.run()
