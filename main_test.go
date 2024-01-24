package main

import (
	"testing"

	"gihtub.com/task4233/protoc/generated"
)

func TestGetOption(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		argColor      generated.Color
		wantNil       bool
		wantColorCode string
	}{
		"correctly got the option for BLACK": {
			argColor:      generated.Color_BLACK,
			wantNil:       false,
			wantColorCode: "#000000",
		},
		"correctly got the option for WHITE": {
			argColor:      generated.Color_WHITE,
			wantNil:       false,
			wantColorCode: "#FFFFFF",
		},
		"failed to get invalid color": {
			argColor:      generated.Color(-1),
			wantNil:       true,
			wantColorCode: "",
		},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			opts := GetOption(tt.argColor)
			if (opts == nil) != tt.wantNil {
				t.Fatalf("unexpected opts: %v", opts)
			}
			if opts == nil {
				return
			}

			if opts.ColorCode != tt.wantColorCode {
				t.Fatalf("unexpected color code, want: %s, got: %s", tt.wantColorCode, opts.ColorCode)
			}
		})
	}

}
