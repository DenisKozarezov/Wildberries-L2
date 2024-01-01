package main

import "testing"

func TestUnpackString(t *testing.T) {
	type args struct {
		zippedStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a4bc2d5e",
			args: args{zippedStr: "a4bc2d5e"},
			want: "aaaabccddddde",
		},
		{
			name: "abcd",
			args: args{zippedStr: "abcd"},
			want: "abcd",
		},
		{
			name: "45",
			args: args{zippedStr: "45"},
			want: "",
		},
		{
			name: "Empty string",
			args: args{zippedStr: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnpackString(tt.args.zippedStr)
			t.Logf("Arg: %s. Got: %s, want: %s", tt.args.zippedStr, got, tt.want)
			if got != tt.want {
				t.Errorf("UnpackString() = %v, want %v", got, tt.want)
			}
		})
	}
}
