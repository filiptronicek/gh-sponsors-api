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
    for i in range(0, 100):
        url = f'https://www.buymeacoffee.com/{usr}?page={i}&notification=1'
        resp = req.get(url)
        print(resp.text)
        usrCount = 0
        sponsors = 0
        if resp.history:
            sponsors = None
        else:
            htmlGH = BeautifulSoup(resp.text, 'html.parser')
            count = htmlGH.select("span.av-heavy")
            print(count)
            for handle in count:
                usrCount += 1
                print(handle)
                totalUsers.append({name: handle.string })

            if usrCount == 0:
                break

        if sponsors == None:
            d = "This user doesn't have BMC"
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