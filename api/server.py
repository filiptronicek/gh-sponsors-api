from http.server import BaseHTTPRequestHandler
from main import getSponsorCount
class handler(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','text/plain')
        self.end_headers()
        message = getSponsorCount()
        self.wfile.write(message.encode())
        return