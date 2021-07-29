package main

import (
	"bytes"
	"testing"
)

func Test_run(t *testing.T) {
	buf := &bytes.Buffer{}
	w = buf
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test 1",
			want: "3 + 7 = 10\n",
		},
	}
	for _, tt := range tests {
		buf.Reset()
		t.Run(tt.name, func(t *testing.T) {
			run()
			if buf.String() != tt.want {
				t.Errorf("wanted is %s, got is %s\n", tt.want, buf.String())
			}
		})
	}
}
