package main

type Transcript struct {
	JobName   string  `json:"jobName"`
	AccountID string  `json:"accountId"`
	Results   Results `json:"results"`
	Status    string  `json:"status"`
}

type Results struct {
	Transcripts   []TranscriptElement `json:"transcripts"`
	SpeakerLabels SpeakerLabels       `json:"speaker_labels"`
	Items         []Item              `json:"items"`
}

type Item struct {
	StartTime    *string       `json:"start_time,omitempty"`
	EndTime      *string       `json:"end_time,omitempty"`
	Alternatives []Alternative `json:"alternatives"`
	Type         Type          `json:"type"`
	LanguageCode *LanguageCode `json:"language_code,omitempty"`
}

type Alternative struct {
	Confidence string `json:"confidence"`
	Content    string `json:"content"`
}

type SpeakerLabels struct {
	ChannelLabel string    `json:"channel_label"`
	Speakers     int64     `json:"speakers"`
	Segments     []Segment `json:"segments"`
}

type Segment struct {
	StartTime    string       `json:"start_time"`
	SpeakerLabel SpeakerLabel `json:"speaker_label"`
	EndTime      string       `json:"end_time"`
	Items        []Segment    `json:"items,omitempty"`
}

type TranscriptElement struct {
	Transcript string `json:"transcript"`
}

type LanguageCode string

const (
	EnUS LanguageCode = "en-US"
)

type Type string

const (
	Pronunciation Type = "pronunciation"
	Punctuation   Type = "punctuation"
)

type SpeakerLabel string

const (
	Spk0 SpeakerLabel = "spk_0"
	Spk1 SpeakerLabel = "spk_1"
)
