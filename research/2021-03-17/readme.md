# March 17 2021

## PHP

~~~php
php > var_export(openssl_get_cipher_methods());
array (
  8 => 'aes-128-ecb',
  19 => 'aes-192-ecb',
  31 => 'aes-256-ecb',
)
~~~

https://php.net/function.openssl-encrypt

## Ruby

~~~ruby
irb(main):009:0> pp OpenSSL::Cipher.ciphers
[
 "aes-128-ecb",
 "aes-192-ecb",
 "aes-256-ecb",
]
~~~
