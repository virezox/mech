# March 17 2021

~~~js
getRandomInt(max) {
   return Math.floor(Math.random() * Math.floor(max));
},

generateNonce(max) {
   const alphabet = "012345689abdef";
   let nonce = '';
   for (let i = 0; i < max; i++) {
      nonce += alphabet[util.getRandomInt(alphabet.length)];
   }
   return nonce;
}
~~~

Derive `uniqId` from `generateNonce(32)`.

Derive `mobileAuthResponse.TOKEN` from `mobile_auth(uniqId)`.

Derive `encrypted` from `mobileAuthResponse.TOKEN`.

Derive `decrypted` from `MOBILE_GW_KEY` and `encrypted`.

Derive `tokenKey` and `token` from `decrypted`.

Derive `gatewayToken` from `tokenKey` and `token`.

Request `api_checkToken(gatewayToken)`. Response `sid`.

~~~js
function createCipher(key) {
   return crypto.createCipheriv('aes-128-ecb', key, null).setAutoPadding(false);
}

paddedLength(data, multiple = 16) {
   const rem = data.length % multiple;
   if (rem > 0) { // If data isn't a multiple pad it out
      return data.length + (multiple - rem);
   } else { // Otherwise just return a buffer around the data
      return data.length;
   }
}

encryptParam(data) {
   const buf = Buffer.alloc(paddedLength(data));
   buf.write(data);
   return createCipher(this.#userKey).update(buf).toString('hex');
}
~~~

Derive `encryptedPassword` from `encryptParam(password)`.

Request `mobile_userAuth(sid, encryptedPassword)`. Response `ARL`.
