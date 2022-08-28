pub struct Clip {
   pub id: u8,
   pub unlisted_hash: String
}

impl Clip {
   // https://player.vimeo.com/video/412573977?h=f7f2d6fcb7
   // https://player.vimeo.com/video/412573977?unlisted_hash=f7f2d6fcb7
   // https://vimeo.com/477957994/2282452868
   // https://vimeo.com/477957994?unlisted_hash=2282452868
   // https://vimeo.com/534685752
   pub fn new(address: String) -> Self {
      Clip{
         id: 0,
         unlisted_hash: address
      }
   }
}
