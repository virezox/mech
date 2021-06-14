# Rotten Tomatoes

~~~
curl -o index.html https://www.rottentomatoes.com/m/one_night_in_miami
~~~

Looking for:

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
