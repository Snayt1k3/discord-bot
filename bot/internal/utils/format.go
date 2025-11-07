package utils

import "time"

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
