package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_GetLevel(t *testing.T) {
	tests := []struct {
		name      string
		giveExp   int
		wantLevel int
	}{
		{
			name:      "0 exp means level 1",
			giveExp:   0,
			wantLevel: 1,
		},
		{
			name:      "1 exp means level 1",
			giveExp:   1,
			wantLevel: 1,
		},
		{
			name:      "99 exp means level 1",
			giveExp:   99,
			wantLevel: 1,
		},
		{
			name:      "100 exp means level 2",
			giveExp:   100,
			wantLevel: 2,
		},
		{
			name:      "1000 exp means level 10",
			giveExp:   1000,
			wantLevel: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			level := (&State{Exp: tt.giveExp}).GetLevel()
			assert.Equal(t, tt.wantLevel, level)
		})
	}
}
