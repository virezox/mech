package goinsta

type Place struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Location *Location `json:"location"`
}
