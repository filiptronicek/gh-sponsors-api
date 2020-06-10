from bs4 import BeautifulSoup
import requests as req
import re
import json

usr = "yg"

url = f'https://github.com/sponsors/{usr}'
resp = req.get(url)

def getSponsorCount():
    if resp.history:
        sponsors = None
    else:
        htmlGH = BeautifulSoup(resp.text, 'html.parser')
        count = htmlGH.select_one("div.border-top:nth-child(2) > p:nth-child(1)").get_text()

        for txt in count.split(" "):
            x = re.search("[0-9]", txt)

            if x:
                sponsors = int(txt)
    return sponsors

def getSponsorNames():
    if resp.history:
        sponsors = None
    else:
        htmlGH = BeautifulSoup(resp.text, 'html.parser')
        count = htmlGH.select("div.mr-1 > a > img")
        users = []
        for handle in count:
            users.append({"handle": handle['alt'].replace('@', ''),"avatar": handle['src']})
        d = users
        d = json.dumps(d)
        return d