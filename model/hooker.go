package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Hooker struct {
	endpoint string
	message  string
}

func NewHooker(endpoint, message string) *Hooker {
	return &Hooker{
		endpoint,
		message,
	}
}

type Message struct {
	Text string `json:"text"`
}

func (h *Hooker) body() (io.Reader, error) {
	msg := Message{
		Text: h.message,
	}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(msgJson), nil
}

func (h *Hooker) Notify() error {
	body, err := h.body()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", h.endpoint, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("status code is not 200")
	}
	return nil
}
