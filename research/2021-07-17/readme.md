# July 17 2021

~~~
.\step oauth -scope https://www.googleapis.com/auth/youtube

https://accounts.google.com/o/oauth2/v2/auth?
client_id=1087160488420-8qt7bavg3qesdhs6it824mhnfgcfe8il.apps.googleusercontent.com&
code_challenge=mNWBt1cT4WgjO043f-fg92Gzi1u8PYEvFA_TRupZ9cg&
code_challenge_method=S256&
nonce=cc840387e435ebf0c099da93b5690d151cb094a5f0a9b2e4653fc9a4fd35134c&
redirect_uri=http%3A%2F%2F127.0.0.1%3A50342&
response_type=code&
scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube&
state=iyoqUaZhy8z5unIgVULFwllteuPdJfQh
~~~

Exchange:

https://github.com/smallstep/cli/blob/ca448947/command/oauth/cmd.go#L875-L897
