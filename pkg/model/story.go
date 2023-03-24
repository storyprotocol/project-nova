package model

// StoryContentModel represents the top level data model for story content
type StoryContentModel struct {
	Title      string               `json:"title"`
	Heading    string               `json:"heading"`
	CoverImage string               `json:"coverImage"`
	Content    []*StorySectionModel `json:"content"`
}

// StorySectionModel represents the data model for a story section
type StorySectionModel struct {
	Type string             `json:"type"`
	Data []*StoryMediaModel `json:"data"`
}

// StoryMediaModel represents the data model for media content
type StoryMediaModel struct {
	Type        string `json:"type"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
}

var StorySectionType = struct {
	Paragraph string
	ImageText string
	TextImage string
}{
	Paragraph: "paragraph",
	ImageText: "imageText",
	TextImage: "textImage",
}

var StoryMediaType = struct {
	Text  string
	Image string
}{
	Text:  "text",
	Image: "image",
}
