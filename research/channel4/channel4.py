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

def post_request(payload):
   r = requests.post(
      license_url,
      json=payload,
      headers={"Content-Type": "application/json"},
      timeout=10
   )
   r.raise_for_status()
   output_json = json.loads(r.text)
   return output_json['license']

def decrypter(license_response):
   data = {
      "buildInfo": buildinfo,
      "headers": '',
      "license": license_url,
      "license_response": license_response,
      "pssh": pssh,
   }
   headers = {'content-type': 'application/json'}
   r = requests.post(baseurl + "/decrypter", headers=headers, json=data)
   r.raise_for_status()
   return r.text

pssh = createpsshfromkid("00000000-0000-0000-0000-000004246624")
baseurl = "http://getwvkeys.cc/"
buildinfo = "Xiaomi/nitrogen/nitrogen:10/QKQ1.190910.002/V12.0.1.0.QEDMIXM:user/release-keys"
license_url = "https://c4.eme.lp.aws.redbeemedia.com/wvlicenceproxy-service/widevine/acquire"

data = {"pssh": pssh, "buildInfo": buildinfo}
r = requests.post(baseurl + "/pssh", json=data)
r.raise_for_status()

payload = {"request_id":5273616,"token":"akRCa0R4aXRs4VRCEaWoO56ylrejJNMpkhZ1OLcwLg7TtVUQPRI1hupcnQsMjLmHpv-D2dhYTfv1vbewV7IG1wy7cIjDz5vtgzFDOGta1jPVqyyinYfUmvpbmuiDMhbJWkYM8QhhYJJlrWc3j9NTk1HqtKHJdgus","video":{"type":"ondemand","url":"https://cf.jos.c4assets.com/CH4_44_7_900_18926001001003_001/CH4_44_7_900_18926001001003_001_J01.ism/stream.mpd?c3.ri=13636097515527362794&mpd_segment_template=time&filter=%28type%3D%3D%22video%22%26%26%28%28DisplayHeight%3E%3D288%29%26%26%28systemBitrate%3C4800000%29%29%29%7C%7Ctype%21%3D%22video%22&ts=1650946629&e=600&st=xqC7CpDNg6D0cJwRZt9ks1Lm8HyjUaEj5Zh77ceSP10"},"message":"CAQ="}
payload["message"] = base64.b64encode(r.content).decode()
license_response = post_request(payload)
print(decrypter(license_response))
