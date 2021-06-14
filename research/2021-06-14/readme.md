# June 14 2021

~~~
curl -o index.html -A iPad `
https://m.youtube.com/results?search_query=Nelly+Furtado+Afraid
~~~

Result:

~~~
>var ytInitialData = '\x7b\x22responseContext\x22:\x7b\x22serviceTrackingPara...
...ms\x22:\x22CAIQsV4iEwiw7tu6tJfxAhUBcIMKHSk_B-A\x3d\x22\x7d\x7d\x7d\x7d\x7d';<
~~~

What is the size difference?

~~~
curl -I https://www.youtube.com/results?search_query=Nelly+Furtado+Afraid
content-length: 420,577

curl -I -A iPad https://m.youtube.com/results?search_query=Nelly+Furtado+Afraid
content-length: 184,228
~~~
