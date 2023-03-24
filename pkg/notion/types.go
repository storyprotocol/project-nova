package notion

type PageResponse struct {
	Id         string    `json:"id"`
	Cover      *Image    `json:"cover"`
	Properties *Property `json:"properties"`
}

type BlockChildrenResponse struct {
	Object     string   `json:"object"`
	Results    []*Block `json:"results"`
	NextCursor *string  `json:"next_cursor"`
	HasMore    bool     `json:"has_more"`
}

type Property struct {
	Title *Title `json:"title"`
}

type Title struct {
	Title []*RichText `json:"title"`
}

type Block struct {
	Id          string     `json:"id"`
	HasChildren bool       `json:"has_children"`
	Type        BlockType  `json:"type"`
	Paragraph   *Paragraph `json:"paragraph"`
	Image       *Image     `json:"image"`
}

type BlockType string

var BlockTypes = struct {
	Paragraph  BlockType
	Image      BlockType
	Column     BlockType
	ColumnList BlockType
}{
	Paragraph:  "paragraph",
	Image:      "image",
	Column:     "column",
	ColumnList: "column_list",
}

type Paragraph struct {
	RichTexts []*RichText `json:"rich_text"`
}

type RichText struct {
	Type        string     `json:"type"`
	Annotations Annotation `json:"annotations"`
	PlainText   string     `json:"plain_text"`
}

type Annotation struct {
	Bold   bool `json:"bold"`
	Italic bool `json:"italic"`
}

type Image struct {
	Type string     `json:"type"`
	File *ImageFile `json:"file"`
}

type ImageFile struct {
	Url string `json:"url"`
}
