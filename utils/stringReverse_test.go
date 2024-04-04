package utils

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test for empty string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "Test for single character",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "Test for single word string",
			args: args{
				s: "Hello",
			},
			want: "olleH",
		},
		{
			name: "Test for multiple word string",
			args: args{
				s: "Hello World!",
			},
			want: "!dlroW olleH",
		},
		{
			name: "Test for non ASCII characters",
			args: args{
				s: "こんにちは",
			},
			want: "はちにんこ",
		},
		{
			name: "Test for repeating string",
			args: args{
				s: "LALLAL",
			},
			want: "LALLAL",
		},
		{
			name: "Test for numbers string",
			args: args{
				s: "1234567890",
			},
			want: "0987654321",
		},
		{
			name: "Test for multiple lines string",
			args: args{
				s: "Test\nTest line 2\nTest line 3",
			},
			want: "3 enil tseT\n2 enil tseT\ntseT",
		},
		{
			name: "Test for long string",
			args: args{
				s: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			},
			want: ".murobal tse di mina tillom tnuresed aiciffo iuq apluc ni tnus ,tnediorp non tatadipuc taceacco tnis ruetpecxE .rutairap allun taiguf ue erolod mullic esse tilev etatpulov ni tiredneherper ni rolod eruri etua siuD .tauqesnoc odommoc ae xe piuqila tu isin sirobal ocmallu noitaticrexe durtson siuq ,mainev minim da mine tU .auqila angam erolod te erobal tu tnudidicni ropmet domsuie od des ,tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
