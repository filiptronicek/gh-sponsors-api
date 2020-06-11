from bs4 import BeautifulSoup
import requests as req
import re
import json
from http.server import BaseHTTPRequestHandler
usr = "yg"


def getSponsorNames(u: str):
    usr = u.split("?u=")[1].split("HTTP")[0].replace(" ", "")
    url = f'https://github.com/sponsors/{usr}'
    resp = req.get(url)
    sponsors = 0
    if resp.history:
        sponsors = None
    else:
        htmlGH = BeautifulSoup(resp.text, 'html.parser')
        count = htmlGH.select("div.mr-1 > a > img")
        users = []
        for handle in count:
            handle['alt'] = handle['alt'].replace('@', '')
            users.append({"handle": handle['alt'],"avatar": handle['src'], "profile": "https://github.com/"+handle['alt']})
        d = users
        d = json.dumps(d)
    if sponsors == None:
        d = "Eror: GitHub Sponsors aren't setup with this user."
    return d
class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','application/json')
        self.send_header('Access-Control-Allow-Origin', '*')
        self.end_headers()
        message = getSponsorNames(self.requestline)
        self.wfile.write(str(message).encode())
        return