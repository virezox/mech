use tinyjson::{
   JsonParseError,
   JsonValue
};

struct JSON_Web(JsonValue);

impl JSON_Web {
   fn token(&self) -> Option<String> {
      self.0["token"].get()
   }
   fn new() -> Result<Self, JsonParseError> {
      let v: JsonValue = s.parse()?;
      Ok(Self(v))
   }
}
