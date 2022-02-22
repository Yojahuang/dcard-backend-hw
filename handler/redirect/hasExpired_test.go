package redirectHandler

import (
    "testing"
)

import (
    "dcard-backend-hw/model"
    "time"
)

func getTime(after int) string {
	result := time.Now()

    for ; after != 0; {
        if after > 0 {
            result = result.Add(time.Hour)
            after--
        } else {
            result = result.Add(time.Hour * -1)
            after++
        }
    }

	return result.Format(time.RFC3339)
}


func TestHasExpired(t *testing.T) {
    type Case struct {
        arg model.Url
        expected bool
    }

    cases := [] Case {
        Case{arg: model.Url{ExpireAt: getTime(1)}, expected: false} ,
        Case{arg: model.Url{ExpireAt: getTime(2)}, expected: false} ,
        Case{arg: model.Url{ExpireAt: getTime(3)}, expected: false} ,
        Case{arg: model.Url{ExpireAt: getTime(4)}, expected: false} ,
        Case{arg: model.Url{ExpireAt: getTime(5)}, expected: false} ,
        Case{arg: model.Url{ExpireAt: getTime(-1)}, expected: true} ,
        Case{arg: model.Url{ExpireAt: getTime(-2)}, expected: true} ,
        Case{arg: model.Url{ExpireAt: getTime(-3)}, expected: true} ,
        Case{arg: model.Url{ExpireAt: getTime(-4)}, expected: true} ,
        Case{arg: model.Url{ExpireAt: getTime(-5)}, expected: true} ,
    }

    for i := 0; i < len(cases) ; i = i + 1 {
        if hasExpired(cases[i].arg) != cases[i].expected {
            t.Errorf("%s expect got %v, but got %v", cases[i].arg.ExpireAt, cases[i].expected, !cases[i].expected)
        }
    }
}
