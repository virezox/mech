package paramount

import (
   "bytes"
   "encoding/base64"
   "encoding/xml"
   "errors"
   "github.com/89z/format"
   "io"
   "net/http"
   "os"
)

type mpd struct {
   Period                    struct {
      AdaptationSet []struct {
         ContentProtection []struct {
            SchemeIdUri string `xml:"schemeIdUri,attr"`
            Pssh        string `xml:"pssh"`
         } `xml:"ContentProtection"`
         Representation []struct {
            ContentProtection []struct {
               SchemeIdUri string `xml:"schemeIdUri,attr"`
               Pssh        struct {
                  Text string `xml:",chardata"`
               } `xml:"pssh"`
            } `xml:"ContentProtection"`
         } `xml:"Representation"`
      } `xml:"AdaptationSet"`
   } `xml:"Period"`
}

var LogLevel format.LogLevel

func newKeys(contentID, bearer string) ([]licenseKey, error) {
   file, err := os.Open("ignore/stream.mpd")
   if err != nil {
      return nil, err
   }
   defer file.Close()
   initData, err := initDataFromMPD(file)
   if err != nil {
      return nil, err
   }
   privateKey, err := os.ReadFile("ignore/device_private_key")
   if err != nil {
      return nil, err
   }
   clientID, err := os.ReadFile("ignore/device_client_id_blob")
   if err != nil {
      return nil, err
   }
   cdm, err := newCDM(privateKey, clientID, initData)
   if err != nil {
      return nil, err
   }
   licenseRequest, err := cdm.getLicenseRequest()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST",
      "https://cbsi.live.ott.irdeto.com/widevine/getlicense?AccountId=cbsi&ContentId=" + contentID,
      bytes.NewReader(licenseRequest),
   )
   if err != nil {
      return nil, err
   }
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   licenseResponse, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return cdm.getLicenseKeys(licenseRequest, licenseResponse)
}

// This function retrieves the PSSH/Init Data from a given MPD file reader.
// Example file: https://bitmovin-a.akamaihd.net/content/art-of-motion_drm/mpds/11331.mpd
func initDataFromMPD(r io.Reader) ([]byte, error) {
   var mpdPlaylist mpd
   if err := xml.NewDecoder(r).Decode(&mpdPlaylist); err != nil {
      return nil, err
   }
   const widevineSchemeIdURI = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
   for _, adaptionSet := range mpdPlaylist.Period.AdaptationSet {
      for _, protection := range adaptionSet.ContentProtection {
         if protection.SchemeIdUri == widevineSchemeIdURI && len(protection.Pssh) > 0 {
            return base64.StdEncoding.DecodeString(protection.Pssh)
         }
      }
   }
   for _, adaptionSet := range mpdPlaylist.Period.AdaptationSet {
      for _, representation := range adaptionSet.Representation {
         for _, protection := range representation.ContentProtection {
            if protection.SchemeIdUri == widevineSchemeIdURI && len(protection.Pssh.Text) > 0 {
               return base64.StdEncoding.DecodeString(protection.Pssh.Text)
            }
         }
      }
   }
   return nil, errors.New("no init data found")
}
