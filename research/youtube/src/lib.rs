use tinyjson::JsonValue;

#[derive(Debug)]
pub struct JsonWeb(JsonValue);

impl JsonWeb {
   pub fn token(self) -> String {
      self.0["token"].clone().try_into().unwrap_or_default()
   }
}

#[derive(Debug)]
pub enum Error {
   HTTP(attohttpc::Error),
   JSON(tinyjson::JsonParseError)
}

impl From<attohttpc::Error> for Error {
   fn from(err: attohttpc::Error) -> Self {
      Self::HTTP(err)
   }
}

impl From<tinyjson::JsonParseError> for Error {
   fn from(err: tinyjson::JsonParseError) -> Self {
      Self::JSON(err)
   }
}

impl JsonWeb {
   pub fn new() -> Result<Self, Error> {
      let res = attohttpc::get("https://vimeo.com/_next/jwt").
         header("X-Requested-With", "XMLHttpRequest").
         send()?;
      let body = res.text()?;
      let value: JsonValue = body.parse()?;
      Ok(Self(value))
   }
}
