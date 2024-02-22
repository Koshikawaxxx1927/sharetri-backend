package utils

import (
    // "fmt"
    "time"
)

var layout = "2006-01-02 15:04:05"

func TimeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func StringToTime(str string) time.Time {
    t, _ := time.Parse(layout, str)
    return t
}