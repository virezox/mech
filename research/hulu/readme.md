# Hulu

Start here:

~~~go
for _, key := range keys {
   if key.Type == widevine.License_KeyContainer_CONTENT {
      command += " --key " + hex.EncodeToString(key.ID) + ":" + hex.EncodeToString(key.Value)
   }
}
~~~

https://github.com/chris124567/hulu/blob/8c095fc9/main.go#L176-L180

Then:

~~~go
keys, err := cdm.GetLicenseKeys(licenseRequest, licenseResponse)
~~~
