package jfmt

import "fmt"

func FmtCount(count uint32) string {
	const unit = 1000
	if count < unit {
		return fmt.Sprintf("%d", count)
	}
	div, exp := uint32(unit), 0
	// 1000, 1000*1000, 1000*1000*1000, 1000*1000*1000*1000
	for n := count / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c",
		float32(count)/float32(div), "KMGTPE"[exp])
}

func FmtFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := uint32(unit), 0
	// 1024, 1024*1024, 1024*1024*1024, 1024*1024*1024*1024
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB",
		float64(size)/float64(div), "KMGTPE"[exp])
}
