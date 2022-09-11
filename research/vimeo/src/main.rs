use vimeo::JsonWeb;

fn main() -> Result<(), vimeo::Error> {
   let web = JsonWeb::new()?;
   let token = web.token();
   println!("{}", token);
   Ok(())
}
