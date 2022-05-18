# Paramount

> Theatricality and deception, powerful agents to the uninitiated.
>
> But we are initiated, aren’t we, Bruce?
>
> The Dark Knight Rises (2012)

https://github.com/ytdl-org/youtube-dl/issues/29038

## Where did proto file come from?

https://github.com/TDenisM/widevinedump/tree/main/pywidevine/cdm/formats

## Where to get CDM?

<https://github.com/Jnzzi/4464_L3-CDM>

## Android client

~~~
com.cbs.app
~~~

Install system certificate.

## Bearer

~~~
com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("6c70b33080758409");

com\cbs\app\androiddata\retrofit\util\RetrofitUtil.java
SecretKeySpec secretKeySpec = new SecretKeySpec(b("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"), "AES");
~~~

https://github.com/matthuisman/slyguy.addons/issues/136
