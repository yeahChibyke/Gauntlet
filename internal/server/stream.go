package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
)

func writeStreamEvent(
	w http.ResponseWriter,
	event string,
	data any,
) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		w,
		"event: %s\ndata: %s\n\n",
		event,
		b,
	); err != nil {
		return err
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	return nil
}

func writeResponseCreated(w http.ResponseWriter) error {
	return writeStreamEvent(
		w,
		"response.created",
		map[string]any{
			"type": "response.created",
		},
	)
}

func writeResponseDelta(
	w http.ResponseWriter,
	resp *responses.Response,
) error {
	return writeStreamEvent(
		w,
		"response.output_text.delta",
		resp,
	)
}

func writeResponseOutputDone(w http.ResponseWriter) error {
	return writeStreamEvent(
		w,
		"response.output_text.done",
		map[string]any{
			"type": "response.output_text.done",
		},
	)
}

func writeResponseCompleted(w http.ResponseWriter) error {
	return writeStreamEvent(
		w,
		"response.completed",
		map[string]any{
			"type": "response.completed",
		},
	)
}

func writeStreamDone(w http.ResponseWriter) error {

	if err := writeResponseOutputDone(w); err != nil {
		return err
	}

	if err := writeResponseCompleted(w); err != nil {
		return err
	}

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
