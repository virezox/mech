use {
   oxhttp::Client,
   oxhttp::model::Method,
   oxhttp::model::Request,
   tinyjson::JsonValue
};

#[derive(Debug)]
pub struct JsonWeb(JsonValue);

impl JsonWeb {
   pub fn new() -> Self {
      let next = "https://vimeo.com/_next/jwt".parse().unwrap();
      let req = Request::builder(Method::GET, next).
         with_header("X-Requested-With".parse().unwrap(), "XMLHttpRequest").
         unwrap().build();
      let res = Client::new().request(req).unwrap();
      let body = res.into_body().to_string().unwrap();
      let value: JsonValue = body.parse().unwrap();
      Self(value)
   }
}

#[test]
fn test_web() {
   let web = JsonWeb::new();
   let token = web.0["token"].is_string();
   assert!(token);
}
