package api

import (
	"testing"
	"timex/types"
)

func Test_validInputSession(t *testing.T) {
	type TestCase struct {
		inp     types.InputSession
		wantErr bool
	}

	tests := []TestCase{
		{types.InputSession{
			Start:      "2025-02-17 10:52:00-05",
			End:        "2025-02-17 11:15:00-05",
			Focus:      2,
			CategoryID: 1,
		}, false},
		{types.InputSession{
			Start:      "2025-02-17 10:52:00-05",
			End:        "2025-02-17 11:01:00-05",
			Focus:      2,
			CategoryID: 1,
		}, true},
	}

	for _, test := range tests {
		err := validInputSession(test.inp)
		if !test.wantErr && err != nil {
			t.Errorf("expected no error, got %+v", err)
		}
	}
}
