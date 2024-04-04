package utils

import "testing"

func TestCountSymbols(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test for empty string",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Test for space string",
			args: args{
				s: "          ",
			},
			want: 10,
		},
		{
			name: "Test for single string",
			args: args{
				s: "Hello",
			},
			want: 5,
		},
		{
			name: "Test for multiple strings",
			args: args{
				s: "Hello World!",
			},
			want: 12,
		},
		{
			name: "Test for non ASCII string",
			args: args{
				s: "こんにちは!",
			},
			want: 6,
		},
		{
			name: "Test for punctuation string",
			args: args{
				s: ".,<>{}[]?!",
			},
			want: 10,
		},
		{
			name: "Test for numbers string",
			args: args{
				s: "1234567890",
			},
			want: 10,
		},
		{
			name: "Test for multiple lines string",
			args: args{
				s: "Test\nTest line 2\nTest line 3",
			},
			want: 28,
		},
		{
			name: "Test for long string",
			args: args{
				s: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			},
			want: 445,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSymbols(tt.args.s); got != tt.want {
				t.Errorf("CountSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
