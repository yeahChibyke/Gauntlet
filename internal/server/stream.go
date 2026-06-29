package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
)

func writeStreamEvent(
	w http.ResponseWriter,
	resp *responses.Response,
) error {

	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(
		w,
		"data: %s\n\n",
		b,
	)

	if err != nil {
		return err
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	return nil
}

func writeStreamDone(
	w http.ResponseWriter,
) error {

	_, err := fmt.Fprint(
		w,
		"data: [DONE]\n\n",
	)

	if err != nil {
		return err
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	return nil
}
