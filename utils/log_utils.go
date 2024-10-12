package utils

import "strings"

func isImportant(log string) bool {
	keywords := []string{"error", "failed", "panic", "critical", "warning"}
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(log), keyword) {
			return true
		}
	}
	return false
}

func FilterImportantLog(logs []string) string {
	importantLogs := []string{}
	var output string

	for _, log := range logs {
		if isImportant(log) {
			importantLogs = append(importantLogs, log)
		}
	}

	if len(importantLogs) == 0 {
		return "No important logs found"
	} else {
		for _, impLog := range importantLogs {
			output = output + impLog + "\n"
		}
	}

	return output
}
