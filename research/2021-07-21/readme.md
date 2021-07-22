# July 21 2021

https://pkg.go.dev/golang.org/x/oauth2/google

## native app approach

Update: here is the process for "native app" [1]. First, program prompts user
to visit page like this:

~~~
https://accounts.google.com/o/oauth2/v2/auth?
client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com&
redirect_uri=urn:ietf:wg:oauth:2.0:oob&
response_type=code&
scope=https://www.googleapis.com/auth/youtube
~~~

As discussed previously, for the `redirect_uri`, you can also use
`http://localhost:999` or similar. Using the "manual copy/paste" method above,
a code will be returned to the user, which they can copy/paste into the program
(PyTube). Then, PyTube can make an internal request like this:

~~~
curl -v `
-d client_id=861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com `
-d client_secret=SboVhoG9s0rNafixCSGGKXAT `
-d code=4/1AX4XfWgGbS8R7Fza-TojWaRuE0QFiN4asvmmc07VKlSjsH0ghn3Sm5... `
-d grant_type=authorization_code `
-d redirect_uri=urn:ietf:wg:oauth:2.0:oob `
https://oauth2.googleapis.com/token
~~~

Response will be `access_token` and `refresh_token`, as before. If PyTube adds
OAuth, will just need to decide to use "device" approach or "native app"
approach, and if using "native app", whether to do "manual copy/paste" or
"loopback ip" method. At this point, I think I am done with this GitHub issue.
Unless someone finds a magic way with the Android keys I put, it seems OAuth is
the best way to handle this situation.

1. https://developers.google.com/identity/protocols/oauth2/native-app
