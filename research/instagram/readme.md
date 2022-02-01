# Instagram

https://github.com/instaloader/instaloader/issues/1022

Using this post:

https://instagram.com/p/CLHoAQpCI2i

here is what we get with `/graphql/query/`:

~~~
720 x 540
239 KB
~~~

here is what we get with `/api/v1/media/`:

~~~
720 x 540
606 KB

720 x 540
239 KB
~~~

here is what we get with Web client:

~~~
720 x 540
1.29 MB
~~~

here is what we get with `__a=1`:

~~~
720 x 540
1.29 MB

720 x 540
606 KB

720 x 540
557 KB

720 x 540
239 KB

720 x 540
217 KB

490 x 368
93.2 KB
~~~
