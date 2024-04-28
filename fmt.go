package jfmt

import (
	"fmt"
	"strings"
	"time"
)

func FmtCount32(count uint32) string {
	const unit = uint32(1000)
	if count < unit {
		return fmt.Sprintf("%d", count)
	}
	div, exp := unit, 0
	// 1000, 1000*1000, 1000*1000*1000, 1000*1000*1000*1000
	for n := count / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%c", float32(count)/float32(div), "KMGTPE"[exp])
}

func FmtCount64(count uint64) string {
	const unit = uint64(1000)
	if count < unit {
		return fmt.Sprintf("%d", count)
	}
	div, exp := unit, 0
	for n := count / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%c", float64(count)/float64(div), "KMGTPE"[exp])
}

func FmtSize64(size uint64) string {
	const unit = uint64(1024)
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := unit, 0
	// 1024, 1024*1024, 1024*1024*1024, 1024*1024*1024*1024
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%cB",
		float64(size)/float64(div), "KMGTPE"[exp])
}

// FmtDuration formats a time.Duration into a shorter string representation.
func FmtDuration(d time.Duration) string {
	var parts []string

	hours := d / time.Hour
	d -= hours * time.Hour
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}

	minutes := d / time.Minute
	d -= minutes * time.Minute
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	seconds := d / time.Second
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%ds", seconds))
	}

	return strings.Join(parts, "")
}
