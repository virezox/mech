from hashlib import sha1
from requests import Request, Session
from requests_toolbelt.utils import dump
import hmac
import os
import time

s = Session()

data = {
   'grant_type': 'password',
   'username': '4095486538',
   'password': os.environ['PASSWORD'],
   'username_is_user_id': '1',
   'client_id': '134',
   'client_secret': '1myK12VeCL3dWl9o/ncV2VyUUbOJuNPVJK6bZZJxHvk='
}

req = Request('POST', 'https://bandcamp.com/oauth_login', data=data)
prepped = req.prepare()
resp = s.send(prepped)
print(dump.dump_all(resp).decode())

time.sleep(1)
dm = resp.headers['X-Bandcamp-DM']
body = prepped.body.encode()
body = dm[1:].encode() + body
prepped.headers['X-Bandcamp-DM'] = hmac.new(b'dtmfa', body, sha1).digest().hex()
resp = s.send(prepped)

print(dump.dump_all(resp).decode())
