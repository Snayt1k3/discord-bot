package utils

import (
	"fmt"
	"strings"
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

func ProgressBlocks(current, max, length int) string {
	if max <= 0 {
		max = 1
	}

	ratio := float64(current) / float64(max)
	filled := int(ratio * float64(length))

	if filled > length {
		filled = length
	}

	return strings.Repeat("▰", filled) + strings.Repeat("▱", length-filled)
}

func FormatVoiceTime(seconds int64) string {
	if seconds == 0 {
		return "None"
	}

	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	months := days / 30
	years := months / 12

	months %= 12
	days %= 30
	hours %= 24
	minutes %= 60

	parts := []string{}

	if years > 0 {
		parts = append(parts, fmt.Sprintf("%dy", years))
	}
	if months > 0 {
		parts = append(parts, fmt.Sprintf("%dmo", months))
	}
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 && years == 0 && months == 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	return strings.Join(parts, " ")
}