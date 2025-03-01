package api

import (
	"fmt"
	"time"
	"timex/types"
)

func validInputSession(inp types.InputSession) error {
	ef := fmt.Errorf
	if inp.Start == 0 {
		return ef("start cannot be empty")
	}
	if inp.End == 0 {
		return ef("end cannot be empty")
	}
	if inp.CategoryID == 0 {
		return ef("category_id cannot be empty")
	}
	if inp.Focus == 0 {
		return ef("focus cannot be empty")
	}
	t2 := time.Unix(inp.End, 0)
	if t2.After(time.Now()) {
		return ef("cannot finish a session in the future")
	}
	delt := t2.Sub(time.Unix(inp.Start, 0)).Minutes()

	if delt < 0 {
		return ef("end time must be bigger than start time")
	}
	if delt < 10 {
		return ef("the delta must be more than 10 minutes. Do you think 10 minutes is considered productive?")
	}
	if inp.Focus < 1 || inp.Focus > 5 {
		return ef("focus must be between 1 and 5")
	}
	return nil
}

func validTimeMode(tm string) bool {
	switch tm {
	case "year", "month", "week", "day":
		return true
	default:
		return false
	}
}
func validTimeHorizon(th string) bool {
	_, errDate := time.Parse("2006-01-02", th)
	switch th {
	case "current", "previous":
		return true
	default:
		return errDate == nil
	}
}

func validVisualMode(vm string) bool {
	return vm == "graph" || vm == "pie"
}
