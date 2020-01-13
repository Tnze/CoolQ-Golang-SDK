package main

import (
	"os/exec"
	"strconv"
	"strings"
)

// 统计当前git代码库的提交次数
func commitCount() (int, error) {
	cmd := exec.Command("git", "rev-list", "--all", "--count")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	seq, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0, err
	}
	return seq, nil
}
