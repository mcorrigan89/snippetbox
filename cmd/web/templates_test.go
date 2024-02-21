package main

import (
	"testing"
	"time"

	"snippetbox.corrigan.io/internal/assert"
)

// func TestHumanDate(t *testing.T) {

// 	tm := time.Date(2024, 2, 21, 10, 15, 0, 0, time.UTC)
// 	hd := humanDate(tm)

// 	if hd != "21 Feb 2024 at 10:15" {
// 		t.Errorf("got %q; want %q", hd, "21 Feb 2024 at 10:15")
// 	}
// }

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 2, 21, 10, 15, 0, 0, time.UTC),
			want: "21 Feb 2024 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 2, 21, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "21 Feb 2024 at 09:15",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
