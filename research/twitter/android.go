package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func search() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{
      strings.Join([]string{
         `OAuth oauth_version=1.0`,
         `oauth_nonce=3631287121008092069727528464482`,
         `oauth_timestamp=1649508643`,
         `oauth_signature=s%2FAtWUq2kmE3Th37knZIsZvxudE%3D`,
         `oauth_consumer_key=3nVuSoBZnx6U4vzUxf5w`,
         `oauth_signature_method=HMAC-SHA1`,
         // not always used:
         `oauth_token=449483305-wcH6DvQDjePDx6LsD4dVtiXvdWxYE8JOfI1KKJjS`,
      }, ","),
   }
   req.URL = new(url.URL)
   //req.URL.Host = "na.glbtls.t.co"
   req.URL.Host = "api.twitter.com"
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
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
