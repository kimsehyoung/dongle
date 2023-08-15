package whisper

const (
	TRANSCRIPTION_PATH = "/speech/transcription"
	TRANSLATION_PATH   = "/speech/translation"
)

type SttRequest struct {
	Audio    string `json:"audio"`
	Language string `json:"language"`
}

type SttResponse struct {
	Text string `json:"text"`
}
