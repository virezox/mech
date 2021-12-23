const TikTok = require("./tiktok")

;(async () => {
   const client = new TikTok();
   await client.generateDevice();
   await client.registerDevice();
})()
