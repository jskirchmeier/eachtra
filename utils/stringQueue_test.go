package utils

import (
	"testing"
)

func TestWordsFromLine(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		words []string
	}{
		{"Empty", "", []string{}},
		{"Empty (multiple spaces)", "     ", []string{}},
		{"one word", "one", []string{"one"}},
		{"one word (multiple spaces)", "one    ", []string{"one"}},
		{"rwo words", "one two", []string{"one", "two"}},
		{"rwo words (multiple spaces)", "one    two   ", []string{"one", "two"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStringQueueFromLine(tt.arg)
			if got.Remaining() != len(tt.words) {
				t.Errorf("NewStringQueueFromLine() = %v, want %v", got, tt.words)
			}
			for i := 0; i < len(tt.words); i++ {
				if !got.More() {
					t.Errorf("NewStringQueueFromLine(), expecting more = %v, want %v", got, tt.words)
				}
				if got.Peek() != tt.words[i] {
					t.Errorf("NewStringQueueFromLine(), peek not what expected = %v, want %v", got.Peek(), tt.words[i])
				}
				if got.Next() != tt.words[i] {
					t.Errorf("NewStringQueueFromLine(), next not what expected = %v, want %v", got.Peek(), tt.words[i])
				}
			}
			if got.More() {
				t.Errorf("NewStringQueueFromLine(), expecting the end (more) = %v, want %v", got, tt.words)
			}
			if got.Remaining() > 0 {
				t.Errorf("NewStringQueueFromLine(), expecting the end (remaining) = %v, want %v", got, tt.words)
			}

		})
	}
}
