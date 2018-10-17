package models

type Comment struct {
	Type    string    `json:"type,omitempty"`
	Content []Content `json:"content"`
}

type Content struct {
	Type    string    `json:"type,omitempty"`
	Text    string    `json:"text,omitempty"`
	Content []Content `json:"content,omitempty"`
}
