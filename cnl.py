#!/usr/bin/python
# -*- coding: utf-8 -*-

# pip install pycrypto

import os
import http.server
import socketserver
import re
import logging
from urllib.parse import parse_qs
from Crypto.Cipher import AES
import base64
port = 9666
addr = "127.0.0.1"

logger = logging.getLogger('cnlpy')
logger.setLevel(logging.DEBUG)
cn_log = logging.StreamHandler()
cn_log.setLevel(logging.DEBUG)
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
cn_log.setFormatter(formatter)
logger.addHandler(cn_log)

class ClickHandler(http.server.BaseHTTPRequestHandler):

    def __init__(self, request, client_address, server):
        self.source=''
        self.urls=''
        self.jk=''
        self.password=''
        self.package=''
        self.method=''

        super().__init__(request, client_address, server)

    # not all implemented jet.. kept this funcs for later commits
    def do_GET(self):
        path = self.path.strip("/").lower()
        self.map = [
        (r"add$", self.add),
        #~ (r"addcrypted$", self.addcrypted),
        (r"addcrypted2$", self.addcrypted2),
        #~ (r"flashgot", self.flashgot),
        #~ (r"crossdomain\.xml", self.crossdomain),
        #~ (r"checkSupportForUrl", self.checksupport),
        (r"jdcheck.js", self.jdcheck),
        (r"", self.flash)
        ]
        func = None
        for r, f in self.map:
            if re.match(r"(flash(got)?/?)?"+r, path):
                func = f
                break
        if func:
            try:
                func()
            except Exception as e :
                raise(e)
                self.send_error(500, str(e))
        else:
            self.send_error(404, "Not Found")

    def do_POST(self):
        length=int(self.headers['Content-Length'])
        if not length:
            logger.error("No Content-Lenth in POST Request")
            return

        self.post = parse_qs(self.rfile.read(length).decode('utf-8'))
        self.do_GET()

    def log_message(self, format, *args):
        # silent mode ;-)
        return

    def success(self):
        self.send_response(200, "OK")
        self.end_headers()

    def response(self, response_fn):
        if response_fn is None:
            self.send_error(404, "Not Found")
            return
        try:
            self.send_response(200, "OK")
            self.end_headers()
            self.wfile.write(response_fn.encode())
        except Exception as e:
            self.send_error(500, str(e))
            return

    def debug(self):
        print("DEBUG")
        print("Address: %s"%(self.client_address,))
        print("Path: %s"%(self.path,))
        print("ReqVers: %s"%(self.request_version,))
        print("Headers: %s"%(self.headers,))

    def get_post(self, name, default=""):
        if name in self.post:
            return self.post[name][0]
        else:
            return default

    def aes_decrypt(self,key, data):
        '''
        decrypt cnl2 crypted
        1. base64 decode
        2. base16 decode key
        3. AES 128 (16 block_size) decryption, iv=key
        '''
        key_bytes = 16
        key = base64.b16decode(key)
        assert len(key) == key_bytes
        # Create AES-CBC cipher. Key is IV
        aes = AES.new(key, AES.MODE_CBC, key)
        # Decrypt and return the plaintext.
        return aes.decrypt(base64.b64decode(data)).decode('utf-8').strip()

    def __str__(self):
        rep="Method: %s\n"%(self.method,)
        if self.package:
            rep+="Package: %s\n"%(self.package,)
        rep+="passwords: %s\nsource: %s\nurls:\n%s\n"%(self.passwords,self.source,self.urls)
        return rep

    def add(self):
        self.method="plain"
        self.urls = self.get_post('urls')
        self.passwords = self.get_post('passwords')
        self.source = self.get_post('source')
        print(self)
        self.success()

    def addcrypted2(self):
        self.method="addcrypted2"
        self.passwords = self.get_post("passwords")
        self.source = self.get_post("source")
        self.package = self.get_post("package")
        jk = self.get_post("jk")
        crypted = self.get_post("crypted")
        m=re.search("\{.?return.?'(\d+)'.?\}",jk)
        if m:
            key = m.group(1)
        else:
            print("No Key found...")
            self.send_error(500, str(e))

        self.urls=self.aes_decrypt(key,crypted)
        print(self)
        self.success()

    def flash(self):
        self.response("JDownloader\r\n")

    def jdcheck(self):
        self.response("jdownloader=true;\nvar version='10629';\n\r\n")

if __name__ == "__main__":
    socketserver.TCPServer.allow_reuse_address = True
    with socketserver.TCPServer((addr, port), ClickHandler) as httpd:
        print("serving at port", port)
        try:
            httpd.serve_forever()
        except KeyboardInterrupt:
            pass
        finally:
            httpd.server_close()
