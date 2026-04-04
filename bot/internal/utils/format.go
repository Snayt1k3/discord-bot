package utils

import (
	"fmt"
	"time"
)

func FormatLastMessage(ts string) string {
	if ts == "" {
		return "No messages yet"
	}
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return ts // fallback — показать как есть
	}
	return t.Format("02 Jan 2006 15:04")
}


func FormatDuration(d time.Duration) string {
    if d <= 0 {
        return "unknown"
    }
    d = d.Round(time.Second)
    h := int(d.Hours())
    m := int(d.Minutes()) % 60
    s := int(d.Seconds()) % 60

    switch {
    case h > 0:
        return fmt.Sprintf("%dh %dm %ds", h, m, s)
    case m > 0:
        return fmt.Sprintf("%dm %ds", m, s)
    default:
        return fmt.Sprintf("%ds", s)
    }
}