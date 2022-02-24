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
    result = result.Add(time.Hour * time.Duration(after))

    return result.Format(time.RFC3339)
}

func TestHasExpired(t *testing.T) {
    type testCase struct {
        arg model.Url
        expected bool
    }

    testCases := [] testCase {
        testCase{arg: model.Url{ExpireAt: getTime(1)}, expected: false} ,
        testCase{arg: model.Url{ExpireAt: getTime(2)}, expected: false} ,
        testCase{arg: model.Url{ExpireAt: getTime(3)}, expected: false} ,
        testCase{arg: model.Url{ExpireAt: getTime(4)}, expected: false} ,
        testCase{arg: model.Url{ExpireAt: getTime(5)}, expected: false} ,
        testCase{arg: model.Url{ExpireAt: getTime(-1)}, expected: true} ,
        testCase{arg: model.Url{ExpireAt: getTime(-2)}, expected: true} ,
        testCase{arg: model.Url{ExpireAt: getTime(-3)}, expected: true} ,
        testCase{arg: model.Url{ExpireAt: getTime(-4)}, expected: true} ,
        testCase{arg: model.Url{ExpireAt: getTime(-5)}, expected: true} ,
    }

    for _, testCase := range testCases {
        if hasExpired(testCase.arg) != testCase.expected {
            t.Errorf("%s expect got %v, but got %v", testCase.arg.ExpireAt, testCase.expected, !testCase.expected)
        }
    }
}
