# Rotten Tomatoes

## Item

Start scanning:

~~~
.Scan()
.Text() == "Wonder"
.Scan()
.Text() == "Woman"
~~~

This could be ReleaseYear, or it could be part of title:

~~~
.Scan()
.Text() == "1984"
~~~

Only way to know is to scan again:

~~~
.Scan()
.Text() == "2020"
~~~

but that could be part of title too, so scan again:

~~~
.Scan()
.Text() == "IMAX"
~~~

now we know that ReleaseYear is 2020.

## Search

~~~
curl -o index.html https://www.rottentomatoes.com/search?search=inception
~~~

Looking for:

~~~html
<script id="movies-json" type="application/json">{"items":[{"imageUrl":"https:...
...Page":"true","hasPreviousPage":"false","startCursor":""},"count":136}</script>
~~~

## Rating

~~~
curl -o index.html https://www.rottentomatoes.com/m/one_night_in_miami
~~~

Looking for:

~~~html
<script id="mps-page-integration">
window.mpscall = {"adunits":"Multi Logo|Box Ad|Marquee Banner|Top Banner",
"cag[score]":"98","cag[certified_fresh]":"1","cag[fresh_rotten]":"fresh",
"cag[rating]":"R","cag[release]":"2020","cag[movieshow]":"One Night in Miami"
,"cag[genre]":"Drama","cag[urlid]":"one_night_in_miami","cat":"movie|movie_page",
"field[env]":"production","field[rtid]":"771534628",
"path":"/m/one_night_in_miami","site":"rottentomatoes-web",
"title":"One Night in Miami","type":"movie_page"};
</script>
~~~

Pretty:

~~~json
{
   "adunits": "Multi Logo|Box Ad|Marquee Banner|Top Banner",
   "cag[certified_fresh]": "1",
   "cag[fresh_rotten]": "fresh",
   "cag[genre]": "Drama",
   "cag[movieshow]": "One Night in Miami",
   "cag[rating]": "R",
   "cag[release]": "2020",
   "cag[score]": "98",
   "cag[urlid]": "one_night_in_miami",
   "cat": "movie|movie_page",
   "field[env]": "production",
   "field[rtid]": "771534628",
   "path": "/m/one_night_in_miami",
   "site": "rottentomatoes-web",
   "title": "One Night in Miami",
   "type": "movie_page"
}
~~~

Alternative:

~~~html
<script type="application/ld+json">
{
   "aggregateRating": {
      "ratingValue": "98"
   }
}
</script>
<script id="score-details-json" type="application/json">
{
   "scoreboard": {
      "audienceScore": "79"
   }
}
</script>
~~~

- <https://developer.fandango.com/rotten_tomatoes>
- <https://github.com/search?q=scoreboard+audienceScore+language:go>
- <https://www.rottentomatoes.com/m/1058966-red>
- <https://www.rottentomatoes.com/m/one_night_in_miami>
