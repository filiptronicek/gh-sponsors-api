from bs4 import BeautifulSoup
import requests as req
import re
import json
from http.server import BaseHTTPRequestHandler

from os import environ, getenv

if environ.get("gh_token") is None:
    from dotenv import load_dotenv
    load_dotenv()

def getUsrDetails(u):
    headers = {"Authorization": f'token {getenv("gh_token")}'}
    reqs = req.get("https://api.github.com/users/"+u).text
    return json.loads(reqs)

def getSponsorNames(u: str):
    usr = u.split("?u=")[1].split("HTTP")[0].replace(" ", "")
    totalUsers = []
    url = f'https://github.com/{usr}?tab=sponsoring'
    resp = req.get(url)
    usrCount = 0
    sponsors = 0
    if resp.history:
        sponsors = None
    else:
        htmlGH = BeautifulSoup(resp.text, 'html.parser')
        count = htmlGH.select("div.d-table > div > a > img:nth-child(1)")

        for handle in count:
            handle['alt'] = handle['alt'].replace('@', '')
            totalUsers.append({"handle": handle['alt'],"avatar": handle['src'], "profile": "https://github.com/"+handle['alt']})


    if sponsors == None:
        d = "Eror: GitHub Sponsors aren't setup with this user."
    d = json.dumps(totalUsers)
    return  '{"sponsorees": '+d+"}"
class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','application/json')
        self.send_header('Access-Control-Allow-Origin', '*')
        self.end_headers()
        message = getSponsorNames(self.requestline)
        self.wfile.write(str(message).encode())
        return