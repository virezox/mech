import base64
import json
import requests

def createpsshfromkid(kid_value):
   kid_value = kid_value.replace('-', '')
   array_of_bytes = bytearray(b'\x00\x00\x002pssh\x00\x00\x00\x00')
   array_of_bytes.extend(bytes.fromhex("edef8ba979d64acea3c827dcd51d21ed"))
   array_of_bytes.extend(b'\x00\x00\x00\x12\x12\x10')
   array_of_bytes.extend(bytes.fromhex(kid_value.replace("-", "")))
   return base64.b64encode(bytes.fromhex(array_of_bytes.hex())).decode('utf-8')

################################################################################

def post_request(payload):
   r = requests.post(
      license_url,
      json=payload,
      headers=selfHeaders,
      timeout=10
   )
   r.raise_for_status()
   output_json = json.loads(r.text)
   return output_json['license']

def decrypter(license_response):
   data = {
      "pssh": pssh,
      "license_response": license_response,
      "license": license_url,
      "headers": selfHeaders,
      "buildInfo": buildinfo
   }
   headers = {
      'user-agent': 'Mozilla/5.0 (Linux; Android 11; M2007J20CG) AppleWebKit/537.36 (KHTML, like Gecko) '
      'Chrome/98.0.4758.60 Mobile Safari/537.36',
      'accept': '*/*',
      'content-type': 'application/json',
   }
   r = requests.post(baseurl + "/decrypter", headers=headers, json=data)
   r.raise_for_status()
   return r.text

selfHeaders = {
   "Accept-Language": "en-US,en;q=0.9",
   "Cache-Control": "no-cache",
   "Connection": "keep-alive",
   "Content-Type": "application/json",
   "Origin": "https://www.channel4.com",
   "Pragma": "no-cache",
   "Referer": "https://www.channel4.com/",
   "Sec-Fetch-Dest": "empty",
   "Sec-Fetch-Mode": "cors",
   "Sec-Fetch-Site": "cross-site",
   "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36",
   "sec-ch-ua-mobile": "?0",
   "sec-ch-ua-platform": "Windows",
}

pssh = createpsshfromkid("00000000-0000-0000-0000-000004246624")

payload = {"request_id":5273616,
"token":"QmlLSlZXWXWKD708zb-Lq6KEEvvmQFAygjf3i3jdvwVdkNUi8Nl2OJSP7VO7_JR71d..."},
"message":"CAQ="}

baseurl = "http://getwvkeys.cc/"
buildinfo = "Xiaomi/nitrogen/nitrogen:10/QKQ1.190910.002/V12.0.1.0.QEDMIXM:user/release-keys"
license_url = "https://c4.eme.lp.aws.redbeemedia.com/wvlicenceproxy-service/widevine/acquire"
data = {"pssh": pssh, "buildInfo": buildinfo}
r = requests.post(baseurl + "/pssh", json=data)
r.raise_for_status()
payload["message"] = base64.b64encode(r.content).decode()
license_response = post_request(payload)
print(decrypter(license_response))
