from bs4 import BeautifulSoup
import requests as req
import re
import json
from http.server import BaseHTTPRequestHandler

from os import environ, getenv

if environ.get("gh_token") is None:
    from dotenv import load_dotenv
    load_dotenv()

def getSponsorNames(u: str):
    usr = u.split("?u=")[1].split("HTTP")[0].replace(" ", "")
    totalUsers = []
    for i in range(1, 100):
        url = f'https://www.buymeacoffee.com/{usr}?page={i}&notification=1'
        resp = req.get(url)
        usrCount = 0
        sponsors = 0
        if resp.history:
            sponsors = None
        else:
            htmlGH = BeautifulSoup(resp.text, 'html.parser')
            count = htmlGH.select("div.mr-1 > a > img")

            for handle in count:
                usrCount += 1
                handle['alt'] = handle['alt'].replace('@', '')
                totalUsers.append({"handle": handle['alt'],"avatar": handle['src'], "profile": "https://github.com/"+handle['alt']})

            if usrCount == 0:
                break

        if sponsors == None:
            d = "Eror: GitHub Sponsors aren't setup with this user."
    d = json.dumps(totalUsers)
    return  '{"sponsors": '+d+"}"

class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','application/json')
        self.send_header('Access-Control-Allow-Origin', '*')
        self.end_headers()
        message = getSponsorNames(self.requestline)
        self.wfile.write(str(message).encode())
        return