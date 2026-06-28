package canonical

// Message represents a single conversational message.
//
// The canonical protocol intentionally stores only plain text.
// Provider-specific concepts (images, audio, tool calls, etc.)
// will be introduced later when Gauntlet supports them.
type Message struct {
	Role Role

	Content string
}
