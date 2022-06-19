# Why use OAuth instead of Cookie?

A benefit of OAuth, is it provides an improved user experience. With cookie,
extracting them is an awkward process, requiring the user the install a browser
extension, or by the tool itself digging into the users cookie in the local
machine, which I think is a privacy violation. Contrast with OAuth, which gives
the same, simple login experience that you'd get with just logging into YouTube
with a browser. For example with my tool, first you see this screen:

> ![terminal code](https://user-images.githubusercontent.com/73562167/174496364-0705642b-601a-4300-8061-4e0be6c7e8f7.png)

Then you go to the page and see this:

> ![web code](https://user-images.githubusercontent.com/73562167/174496444-de1b35a0-e2b3-4689-8ea2-a8b1811703a2.png)

Then:

> ![web email](https://user-images.githubusercontent.com/73562167/174496481-f4752190-d524-49c8-ac11-4946e50de230.png)

Then:

> ![web-allow](https://user-images.githubusercontent.com/73562167/174496530-0dacd308-e27e-4b94-acba-1ca54e61534e.png)

and finally the OAuth refresh token can be saved as JSON:

> ![terminal save](https://user-images.githubusercontent.com/73562167/174496568-9d6c1399-4863-4723-b577-479863bc1dd9.png)
