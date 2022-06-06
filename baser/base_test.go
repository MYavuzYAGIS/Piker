package baser

import (
	"testing"
)

func TestBaseFileDecode(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseFileDecode(); got != tt.want {
				t.Errorf("BaseFileDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseFileEncode(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseFileEncode(); got != tt.want {
				t.Errorf("BaseFileEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseTextDecode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseTextDecode(tt.args.text); got != tt.want {
				t.Errorf("BaseTextDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseTextEncode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BaseTextEncode(tt.args.text); got != tt.want {
				t.Errorf("BaseTextEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
