package ntfy

type NtfyMessage struct {
	ID         string   `json:"id"`
	Time       int      `json:"time"`
	Expires    int      `json:"expires"`
	Event      string   `json:"event"`
	Topic      string   `json:"topic"`
	Priority   int      `json:"priority"`
	Tags       []string `json:"tags"`
	Click      string   `json:"click"`
	Attachment struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Size    int    `json:"size"`
		Expires int    `json:"expires"`
		URL     string `json:"url"`
	} `json:"attachment"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
