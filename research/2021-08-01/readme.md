# August 1 2021

Duration then images?

Images then duration?

---

Can we just look at duration? No:

~~~
0s B3szYRzZqp4 oneohtrix point never - describing bodies
0s ZXNscpJIzQs Oneohtrix Point Never - Returnal (2010) - Describing Bodies
1s 4FnsdJkUBhk Oneohtrix Point Never - Describing Bodies
8s 11Bvzknjo2Q Oneohtrix Point Never - Describing Bodies
~~~

Can we just look at the images? No.

~~~
PS D:\Desktop> curl -I i.ytimg.com/vi/ZXNscpJIzQs/hq1.jpg
Content-Length: 10508
PS D:\Desktop> curl -I i.ytimg.com/vi/ZXNscpJIzQs/hq2.jpg
Content-Length: 10500
8

PS D:\Desktop> curl -I i.ytimg.com/vi/B3szYRzZqp4/hq1.jpg
Content-Length: 17532
PS D:\Desktop> curl -I i.ytimg.com/vi/B3szYRzZqp4/hq2.jpg
Content-Length: 17242
290

PS D:\Desktop> curl -I i.ytimg.com/vi/4FnsdJkUBhk/hq1.jpg
Content-Length: 27879
PS D:\Desktop> curl -I i.ytimg.com/vi/4FnsdJkUBhk/hq2.jpg
Content-Length: 27263
616

PS D:\Desktop> curl -I i.ytimg.com/vi/11Bvzknjo2Q/hq1.jpg
Content-Length: 28382
PS D:\Desktop> curl -I i.ytimg.com/vi/11Bvzknjo2Q/hq2.jpg
Content-Length: 31122
2740
~~~
