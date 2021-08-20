# Vimeo

load request:

~~~
GET /66531465?action=load_download_config HTTP/1.1
Host: vimeo.com
X-Requested-With: XMLHttpRequest
~~~

response:

~~~
{
  "files": [
    {
      "file_name": "gnarls_barkley_-_who's_gonna_save_my_soul_-_from_the_basement.mp4",
      "public_name": "MP4 SD",
      "base_file_name": "gnarls_barkley_-_who's_gonna_save_my_soul_-_from_the_basement",
      "extension": "MP4",
      "download_name": "gnarls_barkley_-_who's_gonna_save_my_soul_-_from_the_basement_640x360.mp4",
      "download_url": "https://player.vimeo.com/play/165631714?s=66531465_1629479284_e36ab2fd96e53bccfa296e64df8ee05e&loc=external&context=Vimeo%5CController%5CClipController.main&download=1",
      "is_cold": false,
      "is_defrosting": false,
      "is_source": false,
      "size": "18.829MB",
      "size_short": "18.83MB",
      "height": 360,
      "width": 640,
      "range": "SDR"
    }
  ],
  "is_owner": false,
  "allow_downloads": true,
  "is_spatial": false
}
~~~

config request:

~~~
GET /video/66531465/config HTTP/1.1
Host: player.vimeo.com
~~~

response:

~~~json
{
  "video": {
    "title": "Gnarls Barkley - Who's Gonna Save My Soul - From the Basement"
  }
}
~~~

looking for:

~~~
uploaded_on
~~~
