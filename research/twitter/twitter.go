package twitter

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/klaidas/go-oauth1"
   "net/http"
   "net/url"
   "strings"
)

type xauth struct {
   OAuth_Token string
   OAuth_Token_Secret string
}

func (x xauth) auth(method, addr string, param map[string]string) string {
   auth := go_oauth1.OAuth1{
      AccessSecret: x.OAuth_Token_Secret,
      AccessToken: x.OAuth_Token,
      ConsumerKey: "3nVuSoBZnx6U4vzUxf5w",
      ConsumerSecret: "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
   }
   return auth.BuildOAuth1Header(method, addr, param)
}

func search() (*http.Response, error) {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{
      strings.Join([]string{
         `OAuth oauth_version=1.0`,
         `oauth_nonce=3631287121008092069727528464482`,
         `oauth_timestamp=1649508643`,
         `oauth_signature=s%2FAtWUq2kmE3Th37knZIsZvxudE%3D`,
         `oauth_consumer_key=3nVuSoBZnx6U4vzUxf5w`,
         `oauth_signature_method=HMAC-SHA1`,
         `oauth_token=449483305-wcH6DvQDjePDx6LsD4dVtiXvdWxYE8JOfI1KKJjS`,
      }, ","),
   }
   req.URL = new(url.URL)
   req.URL.Host = "na.glbtls.t.co"
   //req.URL.Host = "api.twitter.com"
   req.URL.Path = "/2/search/adaptive.json"
   val := make(url.Values)
   val["cards_platform"] = []string{"Android-12"}
   val["earned"] = []string{"true"}
   val["ext"] = []string{"mediaRestrictions,altText,mediaStats,mediaColor,info360,highlightedLabel,superFollowMetadata,hasNftAvatar,unmentionInfo"}
   val["include_blocked_by"] = []string{"true"}
   val["include_blocking"] = []string{"true"}
   val["include_cards"] = []string{"true"}
   val["include_carousels"] = []string{"true"}
   val["include_composer_source"] = []string{"true"}
   val["include_entities"] = []string{"true"}
   val["include_ext_enrichments"] = []string{"true"}
   val["include_ext_has_nft_avatar"] = []string{"true"}
   val["include_ext_media_availability"] = []string{"true"}
   val["include_ext_professional"] = []string{"true"}
   val["include_ext_replyvoting_downvote_perspective"] = []string{"true"}
   val["include_ext_sensitive_media_warning"] = []string{"true"}
   val["include_media_features"] = []string{"true"}
   val["include_profile_interstitial_type"] = []string{"true"}
   val["include_quote_count"] = []string{"true"}
   val["include_reply_count"] = []string{"true"}
   val["include_user_entities"] = []string{"true"}
   val["include_viewer_quick_promote_eligibility"] = []string{"true"}
   val["q"] = []string{"cats"}
   val["query_source"] = []string{"typed_query"}
   val["simple_quoted_tweet"] = []string{"true"}
   val["spelling_corrections"] = []string{"true"}
   val["tweet_mode"] = []string{"extended"}
   val["tweet_search_mode"] = []string{"top"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

func NewGuest() (*Guest, error) {
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/1.1/guest/activate.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + bearer)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   guest := new(Guest)
   if err := json.NewDecoder(res.Body).Decode(guest); err != nil {
      return nil, err
   }
   return guest, nil
}

type Guest struct {
   Guest_Token string
}

var LogLevel format.LogLevel

func (g Guest) xauth(identifier, password string) (*xauth, error) {
   body := url.Values{
      "x_auth_identifier": {identifier},
      "x_auth_password": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/auth/1/xauth_password.json",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "Content-Type": {"application/x-www-form-urlencoded"},
      "X-Guest-Token": {g.Guest_Token},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(xauth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

