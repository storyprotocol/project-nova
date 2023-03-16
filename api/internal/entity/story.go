package entity

import "time"

type StoryChapterResp struct {
	ChapterNum int    `json:"chapterNum"`
	Title      string `json:"title"`
	Heading    string `json:"heading"`
	CoverUrl   string `json:"coverUrl"`
}

// StoryChapterModel represents the story chapter's model in data storage
type StoryChapterModel struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	StoryId   string    `json:"storyId"`
	SeqNum    int       `json:"seqNum"`
	Title     *string   `json:"title"`
	Heading   *string   `json:"heading"`
	CoverUrl  *string   `json:"coverUrl"`
	ReleaseAt time.Time `json:"releaseAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (StoryChapterModel) TableName() string {
	return "story_chapter"
}

func (s *StoryChapterModel) ToStoryChapterResp() *StoryChapterResp {
	resp := &StoryChapterResp{
		ChapterNum: s.SeqNum,
	}

	if s.Heading != nil {
		resp.Heading = *s.Heading
	}
	if s.Title != nil {
		resp.Title = *s.Title
	}
	if s.CoverUrl != nil {
		resp.CoverUrl = *s.CoverUrl
	}

	return resp
}
