# Docs

## How to save HAR?

Open browser and open a new tab. Click `Open menu`, `Web Developer`, `Network`.
Then click `Network Settings`, `Persist Logs`. Also check `Disable Cache`. Then
browse to the page you want to capture. Once you are ready, click `Pause`, then
click `Network Settings`, `Save All As HAR`. Rename file to JSON.

## How to add a new site?

1. `--proxy 127.0.0.1:8080 --no-check-certificate`
2. try navigating to the target page from home screen, instead of going directly
   to page
3. check media page for JSON requests
4. MITM Proxy Firefox
5. MITM Proxy APK
6. JADX

## Sites

Amazon Prime Video:

https://github.com/ytdl-org/youtube-dl/issues/1753

Apple TV:

https://github.com/ytdl-org/youtube-dl/issues/30808
