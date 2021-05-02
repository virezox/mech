# YouTube

Download from YouTube

https://pkg.go.dev/github.com/89z/youtube

~~~js
eg: extract decipher from
https://youtube.com/s/player/4fbb4d5b/player_ias.vflset/en_US/base.js

var Mt={
splice:function(a,b){a.splice(0,b)},
reverse:function(a){a.reverse()},
EQ:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c}};

a=a.split("");
Mt.splice(a,3);
Mt.EQ(a,39);
Mt.splice(a,2);
Mt.EQ(a,1);
Mt.splice(a,1);
Mt.EQ(a,35);
Mt.EQ(a,51);
Mt.splice(a,2);
Mt.reverse(a,52);
return a.join("")
~~~

## Thanks

https://tyrrrz.me/blog/reverse-engineering-youtube
