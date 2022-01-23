# gist.github.com/sclark39/9daf13eea9c0b381667b61e3d2e7bc11
import base64

def shortcode_to_id(shortcode):
     code = ('A' * (12-len(shortcode)))+shortcode
     return int.from_bytes(base64.b64decode(code.encode(), b'-_'), 'big')

id = shortcode_to_id('CY7zg-ulZEZ')
# 2755022163816059161
print(id)
