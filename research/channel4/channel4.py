import base64
import json
import requests

version = "0.1.0"


class channel4_main:
    def __init__(self, Kid) -> None:
        self.baseurl = "http://getwvkeys.cc/"
         self.buildinfo = "Xiaomi/nitrogen/nitrogen:10/QKQ1.190910.002/V12.0.1.0.QEDMIXM:user/release-keys"
        self.json_payloads = {"request_id": 5322675,
                              "token": "SUZYT1pzR2Lwpe5zgnaLGRmLia8ssXPI2ctf...",
                              "video": {"type": "ondemand",
                                        "url": "https://ak-jos-c4assets-com..."},
                              "message": "CAQ="}
        self.headers = self.header()
        self.pssh = self.createpsshfromkid(Kid)
        self.license_url = "https://c4.eme.lp.aws.redbeemedia.com/wvlicenceproxy-service/widevine/acquire"
        self.generate_request_api = self.baseurl + "/pssh"
        self.decrypter_api = self.baseurl + "/decrypter"
        self.cache_api = self.baseurl + "/findpssh"

    @staticmethod
    def createpsshfromkid(kid_value):
        kid_value = kid_value.replace('-', '')
        assert len(kid_value) == 32 and not isinstance(kid_value, bytes), "wrong KID length"
        array_of_bytes = bytearray(b'\x00\x00\x002pssh\x00\x00\x00\x00')
        array_of_bytes.extend(bytes.fromhex("edef8ba979d64acea3c827dcd51d21ed"))
        array_of_bytes.extend(b'\x00\x00\x00\x12\x12\x10')
        array_of_bytes.extend(bytes.fromhex(kid_value.replace("-", "")))
        return base64.b64encode(bytes.fromhex(array_of_bytes.hex())).decode('utf-8')

    def match_pssh(self, find_pssh):
        r = requests.post(self.cache_api, data=find_pssh)
        if r.text != "":
            print("Cached keys:\n", r.text, "\n")

    @staticmethod
    def header():
        headers = {
            "Accept-Language": "en-US,en;q=0.9",
            "Sec-Fetch-Site": "cross-site",
            "Sec-Fetch-Mode": "cors",
            "Connection": "keep-alive",
            "sec-ch-ua-platform": "Windows",
            "sec-ch-ua-mobile": "?0",
            "Sec-Fetch-Dest": "empty",
            "Origin": "https://www.channel4.com",
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) "
                          "Chrome/98.0.4758.80 Safari/537.36",
            "Referer": "https://www.channel4.com/",
            "Pragma": "no-cache",
            "Cache-Control": "no-cache",
            "Content-Type": "application/json"
        }
        return headers

    def generate_request(self):
        data = {"pssh": self.pssh, "buildInfo": self.buildinfo}
        r = requests.post(self.generate_request_api, json=data)
        r.raise_for_status()
        return r.content

    def post_request(self, license_request):
        self.json_payloads["message"] = base64.b64encode(license_request).decode()
        r = requests.post(self.license_url, json=self.json_payloads, headers=self.headers, timeout=10)
        r.raise_for_status()
        output_json = json.loads(r.text)
        return output_json['license']

    def decrypter(self, license_url, license_response):
        data = {
            "pssh": self.pssh,
            "license_response": license_response,
            "license": license_url,
            "headers": self.headers,
            "buildInfo": self.buildinfo
        }
        headers = {
            'user-agent': 'Mozilla/5.0 (Linux; Android 11; M2007J20CG) AppleWebKit/537.36 (KHTML, like Gecko) '
                          'Chrome/98.0.4758.60 Mobile Safari/537.36',
            'accept': '*/*',
            'content-type': 'application/json',
        }
        r = requests.post(self.decrypter_api, headers=headers, json=data)
        r.raise_for_status()
        return r.text

    def main(self):
        self.match_pssh(self.pssh)
        license_request = self.generate_request()
        license_response = self.post_request(license_request)
        print("\n" + self.decrypter(self.license_url, license_response).replace("<br>", "\n"))

if __name__ == "__main__":
   kid = input("KID:")
   start = channel4_main(kid)
   start.main()
