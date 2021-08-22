from math import ceil
import base64

def string(field_number, data):
    data = as_bytes(data)
    return _proto_field(2, field_number, varint_encode(len(data)) + data)

def as_bytes(value):
    if isinstance(value, str):
        return value.encode('utf-8')
    return value

def uint(field_number, value):
    return _proto_field(0, field_number, varint_encode(value))

def _proto_field(wire_type, field_number, data):
    return varint_encode( (field_number << 3) | wire_type) + data

def varint_encode(offset):
    needed_bytes = ceil(offset.bit_length()/7) or 1 # (0).bit_length() returns
    encoded_bytes = bytearray(needed_bytes)
    for i in range(0, needed_bytes - 1):
        encoded_bytes[i] = (offset & 127) | 128  # 7 least significant bits
        offset = offset >> 7
    encoded_bytes[-1] = offset & 127 # leave first bit as zero for
    return bytes(encoded_bytes)

offset = 0
sort = 0
video_id = b'q5UnT4Ik6KU'
page_info = string(4,video_id) + uint(6, sort)
page_params = string(2, video_id)
offset_information = string(4, page_info) + uint(5, offset)
result = string(2, page_params) + uint(3,6) + string(6, offset_information)
ctoken = base64.urlsafe_b64encode(result).decode('ascii')
print(ctoken)
