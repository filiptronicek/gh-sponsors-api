from bs4 import BeautifulSoup
import requests as req
import re
import json
from http.server import BaseHTTPRequestHandler
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
    return json.dumps({"count":sponsors})

class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','application/json')
        self.end_headers()
        message = self.requestline
        self.wfile.write(str(message).encode())
        return