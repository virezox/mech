# March 15 2021

## song.getData

~~~
const gatewayWWW = "http://www.deezer.com/ajax/gw-light.php"

type song struct {
   Results struct {
      MD5_Origin string
      Track_Token string
   }
}

func newSong(apiToken, sid string, sngId int) (*song, error) {
   in, out := map[string]int{"SNG_ID": sngId}, new(bytes.Buffer)
   json.NewEncoder(out).Encode(in)
   req, err := mech.NewRequest("POST", gatewayWWW, out)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("api_version", "1.0")
   val.Set("input", "3")
   val.Set("method", "song.getData")
   val.Set("api_token", apiToken)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Cookie", "sid=" + sid)
}
~~~

## Fail

~~~
deezer.ping -> deezer.getUserData -> song.getData -> get_url
                     |                                  ^
                     |__________________________________|
~~~

## Pass cookie

~~~
deezer.getUserData -> deezer.pageTrack
~~~

~~~
deezer.getUserData -> song.getData
~~~

## Pass no cookie

~~~
deezer.ping -> deezer.getUserData -> song.getListData -> get_url
                     |                                      ^
                     |______________________________________|
~~~

~~~
deezer.ping -> deezer.getUserData -> deezer.pageTrack -> get_url
                     |                                      ^
                     |______________________________________|
~~~

Now what about this:

~~~
http://api-v3.deezer.com/1.0/gateway.php
http://api.deezer.com/1.0/gateway.php
deezer gateway api_key
~~~

https://github.com/svbnet/diezel

Here are some methods used:

~~~
shouldn't need to call this normally
api_checkToken

shouldn't need to call this normally
mobile_auth

Signs in with an email and password
mobile_userAuth

Restores a signed-in session
mobile_userAutoLog

Gets a song by ID
song_getData
~~~

Lets look at the APK again:

https://apkmirror.com/apk/deezer-music

~~~
assets/icon2.png
~~~

Result:

~~~
MOBILE_GW_KEY
VBK1FSUEXHTSDBJJ

MOBILE_API_KEY
4VCYIJUCDLOUELGD1V8WBVYBNVDYOXEWSLLZDONGBBDFVXTZJRXPR29JRLQFO6ZE
~~~

So I downloaded the Deezer app from here:

https://store.rg-adguard.net

by choosing `ProductId` and entering `9nblggh6j7vv`. The `Fast`, `RP` and
`Retail` options all seem to produce the same result. I downloaded these:

~~~
Deezer.62021768415AF_4.32.30.0_neutral_~_q7m17pa7q8kj0.BlockMap
Deezer.62021768415AF_4.32.30.0_neutral_~_q7m17pa7q8kj0.appxbundle
~~~

and searched but did not find a key. I also extracted `app/Deezer.exe` and
searched it with no luck. I should post the AdGuard site here:

https://superuser.com/questions/501699

Here is the Android app:

https://play.google.com/store/apps/details?id=deezer.android.app

This is the current Android key:

~~~
4VCYIJUCDLOUELGD1V8WBVYBNVDYOXEWSLLZDONGBBDFVXTZJRXPR29JRLQFO6ZE
~~~

Extract the APK, then look inside `classes.dex`. Here is the Apple app:

https://apps.apple.com/app/deezer-music-podcast-player/id292738169

This is the Apple key:

~~~
ZAIVAHCEISOHWAICUQUEXAEPICENGUAFAEZAIPHAELEEVAHPHUCUFONGUAPASUAY
~~~

APK is 27.2 MB. What about Apple? Here is link to Apple iOS IPA file:

https://appdb.to/app/ios/292738169

which is 138 MB.

- https://apple.stackexchange.com/questions/415575
- https://tor.stackexchange.com/questions/9230
