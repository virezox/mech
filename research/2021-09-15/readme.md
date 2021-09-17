# September 16 2021

This is what we care about:

~~~java
p8.b(v4_2.toString(), this.a(p8.d()));
~~~

so we need to find `this.a`, which takes a single argument, and returns a
string:

~~~java
private String a(byte[] body)
{
   return this.a(this.a(), body);
}
~~~

Now we need to find `this.a`, which takes no arguments and return byte slice:

~~~java
private byte[] a()
{
  byte[] v0_0 = this.a;
  if (v0_0 == null) {
      v0_0 = new byte[0];
  }
  return v0_0;
}
~~~

Now we need to find `this.a`, which takes two arguments, and returns a string:

~~~java
public String a(byte[] p4, byte[] body)
{
  if (body != null) {
      byte[] v1_6 = new byte[(p4.length + body.length)];
      System.arraycopy(p4, 0, v1_6, 0, p4.length);
      System.arraycopy(body, 0, v1_6, p4.length, body.length);
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
