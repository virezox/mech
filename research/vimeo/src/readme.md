# Vimeo

~~~rust
use minreq;

fn main() -> Result<(), minreq::Error> {
   let res = minreq::get("https://speedtest.lax.hivelocity.net").send()?;
   let body = res.as_str()?;
   print!("{}", body);
   Ok(())
}
~~~
