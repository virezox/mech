import json
import os
import requests
import sys
from requests_toolbelt.utils import dump

headers = {
	'X-Requested-With': 'com.bandcamp.android',
	'Content-Type': 'application/json',
	'User-Agent': 'Dalvik/2.1.0 (Linux; U; Android 9; Unknown Device)',
	'Host': 'bandcamp.com',
}

data = {"platform":"a","version":191577}
resp = requests.post(
   'https://bandcamp.com/api/mobile/24/bootstrap_data',
   headers=headers, json=data
)
data = dump.dump_all(resp)
print(data.decode())
