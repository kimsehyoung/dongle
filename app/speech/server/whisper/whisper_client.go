package whisper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WhisperClient struct {
	serverAddr string
	*http.Client
}

var (
	once          sync.Once
	whisperClient *WhisperClient
)

func GetWhisperClient(serverAddr string) *WhisperClient {
	once.Do(func() {
		whisperClient = &WhisperClient{
			serverAddr: serverAddr,
			Client:     &http.Client{Timeout: 10 * time.Second},
		}
	})
	return whisperClient
}

func (c *WhisperClient) SpeechToText(endpoint string, body SttRequest) (*SttResponse, error) {
	// Encode request body to json string.
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "marshalling request body: %v", err)
	}
	// Create HTTP Request
	fmt.Println(c.serverAddr + endpoint)
	req, err := http.NewRequest("POST", c.serverAddr+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request to Server
	resp, err := c.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "response from Whisper server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Internal, "received status code %d", resp.StatusCode)
	}
	// Decode HTTP response
	var sttResponse SttResponse
	err = json.NewDecoder(resp.Body).Decode(&sttResponse)
	if err != nil {
		return nil, err
	}
	return &sttResponse, nil
}
