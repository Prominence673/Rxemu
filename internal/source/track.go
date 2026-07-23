package source

type Track struct{
	ID string  `json:"id"`
	Title string  `json:"title"`
	Artist string  `json:"artist"`
	Duration float64 `json:"duration"`
	URL string  `json:"url"`
}

type searchDocument struct{
	Entries []searchEntry `json:"entries"`
}
	
type searchEntry struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Channel string `json:"channel"`
	Uploader string `json:"uploader"`
	Duration float64 `json:"duration"`
	WebpageURL string `json:"webpage_url"`
}