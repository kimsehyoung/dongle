package google

// https://pkg.go.dev/cloud.google.com/go/texttospeech/apiv1/texttospeechpb#SsmlVoiceGender
// https://pkg.go.dev/cloud.google.com/go/texttospeech/apiv1/texttospeechpb#AudioEncoding
// https://cloud.google.com/text-to-speech/docs/voices?hl=ko

type TtsRequest struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

type TtsResponse struct {
	Audio []byte `json:"audio"`
}

func ConvertLang(lang string) string {
	switch lang {
	case "ko":
		return "ko-KR"
	case "en":
		return "en-US"
	case "ja":
		return "ja-JP"
	case "zh":
		return "cmn-CN"
	default:
		return ""
	}
}
