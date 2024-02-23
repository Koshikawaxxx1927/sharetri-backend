package utils

import (
    "fmt"
    "time"
)

var layout = "2006-01-02T15:04:05+09:00"

func TimeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func StringToTime(str string) time.Time {
    fmt.Println(str)
    t, _ := time.Parse(layout, str)
    fmt.Println(t)
    return t
}