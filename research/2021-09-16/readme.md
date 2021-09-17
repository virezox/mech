# September 16 2021

This is what we care about:

~~~java
p8.b(v4_2.toString(), this.a(p8.d()));
~~~

so we need to find `this.a`, which takes a single argument, and returns a
string:

~~~java
private String a(byte[] p2)
{
   return this.a(this.a(), p2);
}
~~~

Now we need to find `this.a`, which takes two arguments, and returns a string:

~~~java
public String a(byte[] p4, byte[] p5)
{
  if (p5 != null) {
      byte[] v1_6 = new byte[(p4.length + p5.length)];
      System.arraycopy(p4, 0, v1_6, 0, p4.length);
      System.arraycopy(p5, 0, v1_6, p4.length, p5.length);
      p4 = v1_6;
  }
  String v5_1 = this.d;
  if (v5_1 <= null) {
      return com.bandcamp.shared.util.i.b(new String(p4), new String(com.bandcamp.android.network.d.j));
  } else {
      String v4_6;
      if (v5_1 <= 30) {
          v4_6 = com.bandcamp.shared.util.i.a(new String(p4), new String(this.c), 0);
      } else {
          v4_6 = com.bandcamp.shared.util.i.a(new String(p4), new String(this.c), 0, 0);
      }
      return v4_6;
  }
}
~~~

- https://androguard.readthedocs.io/en/latest/intro/gettingstarted.html
- https://forensics.spreitzenbarth.de/2015/10/05/androguard-a-simple-step-by-step-guide/
- https://github.com/the-eater/camp-collective/issues/5
