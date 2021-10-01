package goinsta

type reelShare struct {
	Text        string `json:"text"`
	IsPersisted bool   `json:"is_reel_persisted"`
	OwnderID    int64  `json:"reel_owner_id"`
	Type        string `json:"type"`
	ReelType    string `json:"reel_type"`
	Media       Item   `json:"media"`
}

type actionLog struct {
	Description string `json:"description"`
}
