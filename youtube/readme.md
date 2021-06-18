# YouTube

Download from YouTube

## Search

So, we are looking for this:

~~~
/watch?v=XFkzRNyygfk
~~~

First result:

~~~
<script nonce="TCh7gubawSzSBgq1Zg3rSA">var ytInitialData = {"responseContext"...
...ead creep cover","radiohead fake plastic trees","radiohead kid a"]};</script>
~~~

Everything after `var ytInitialData =` and before `;` is valid JSON. The search
results are here:

~~~
contents	
   twoColumnSearchResultsRenderer	
      primaryContents	
         sectionListRenderer	
            contents	
               0	
                  itemSectionRenderer	
                     contents
~~~

careful, first result might be an advertisement.

## Free proxy list

https://proxy.webshare.io/register

## Links

- https://github.com/iawia002/annie/issues/839
- https://github.com/kkdai/youtube/issues/186
- https://golang.org/pkg/net/http#Header.WriteSubset
- https://superuser.com/questions/773719/how-do-all-of-these-save-video
