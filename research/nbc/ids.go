package main

var accountIDs = []int{
   2304980677,
   2304981409,
   2304982139,
   2304985974, // .channelsConfig[0]
   2304986603,
   2304987407,
   2304988172,
   2304989560,
   2304990266,
   2304991196,
   2304992029,
   2310481745,
   2410887629, // .globalConfig
   2702430253,
   2708904471,
}

var videos = []int{
   9000246183,
   9000221348,
}

type vodRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID int `json:"accountId"`
   } `json:"mpx"`
}
