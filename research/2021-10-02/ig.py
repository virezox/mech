from datetime import datetime
from requests_toolbelt.utils import dump
import os
import requests

url = 'https://www.instagram.com/accounts/login/'
login_url = 'https://www.instagram.com/accounts/login/ajax/'
time = int(datetime.now().timestamp())
response = requests.get(url, auth=None, headers={'User-Agent': 'Mozilla'})
csrf = response.cookies['csrftoken']
username = 'srpen6'
password = os.environ['PASS']

payload = {
   'username': username,
   'enc_password': f'#PWD_INSTAGRAM_BROWSER:0:{time}:{password}',
   'queryParams': {},
   'optIntoOneTap': 'false'
}

login_header = {
   "User-Agent": "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko)"
                 " Chrome/77.0.3865.120 Safari/537.36",
   "X-Requested-With": "XMLHttpRequest",
   "Referer": "https://www.instagram.com/accounts/login/",
   "x-csrftoken": csrf
}

response = requests.post(login_url, data=payload, headers=login_header)
data = dump.dump_all(response)
print(data.decode())
