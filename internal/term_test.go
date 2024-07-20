package internal

import (
	"bytes"
	"testing"
)

func TestSetWriterCapturesOutput(t *testing.T) {
	var buf bytes.Buffer
	SetWriter(&buf)
	ch := make(chan string)
	go Show(ch)

	message := "Hello, World!"
	ch <- message
	close(ch)

	if got := buf.String(); got != message {
		t.Errorf("Expected output %q, got %q", message, got)
	}
}

func TestSetWriterSupportsMultipleMessages(t *testing.T) {
	var buf bytes.Buffer
	SetWriter(&buf)
	ch := make(chan string)
	go Show(ch)

	messages := []string{"Hello", "World", "!"}
	expected := "HelloWorld!"
	for _, msg := range messages {
		ch <- msg
	}
	close(ch)

	if got := buf.String(); got != expected {
		t.Errorf("Expected concatenated output %q, got %q", expected, got)
	}
}

func TestShowHandlesConcurrentWrites(t *testing.T) {
	var buf bytes.Buffer
	SetWriter(&buf)
	ch := make(chan string, 3) // Buffered channel to prevent blocking

	go Show(ch)

	messages := []string{"Concurrent", "Writes", "Test"}
	for _, msg := range messages {
		ch <- msg
	}
	close(ch)

	// Note: This test assumes Show handles concurrent writes correctly.
	// It does not guarantee order due to the nature of concurrent execution.
	expected := "ConcurrentWritesTest" // Assuming order is preserved, which might not always be the case
	if got := buf.String(); got != expected {
		t.Errorf("Expected concatenated output %q, got %q", expected, got)
	}
}
