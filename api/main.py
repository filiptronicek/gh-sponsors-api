from bs4 import BeautifulSoup
import requests as req
import re
import json
from http.server import BaseHTTPRequestHandler


def getSponsorCount(u: str):
    usr = u.split("?u=")[1].split("HTTP")[0].replace(" ", "")
    url = f'https://github.com/sponsors/{usr}'
    resp = req.get(url)

    if resp.history:
        sponsors = None
    else:
        htmlGH = BeautifulSoup(resp.text, 'html.parser')
        count = htmlGH.select_one("div.border-top:nth-child(2) > p:nth-child(1)").get_text()

        for txt in count.split(" "):
            x = re.search("[0-9]", txt)
            if x:
                sponsors = int(txt)
    if sponsors == None:
        sponsors = "Eror: GitHub Sponsors aren't setup with this user."
    return json.dumps({"count":sponsors})

class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','application/json')
        self.end_headers()
        message = getSponsorCount(self.requestline)
        self.wfile.write(str(message).encode())
        return