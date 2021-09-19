from requests_toolbelt.utils import dump
import requests
import sys

auth = requests.auth.HTTPBasicAuth('<CLIENT_ID>', '<SECRET_TOKEN>')

data = {'grant_type': 'password',
        'username': 'svnpenn',
        'password': sys.argv[1]}

headers = {'User-Agent': 'MyBot/0.0.1'}

resp = requests.post('https://www.reddit.com/api/v1/access_token',
                    auth=auth, data=data, headers=headers)

data = dump.dump_all(resp)
print(data.decode())
