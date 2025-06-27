package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BypassList(list string, hideFails bool) {
	file, err := os.Open(list)
	if err != nil {
		fmt.Println("[-] Error opening file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fmt.Println(strings.Repeat("-", 40))
		BypassUrl(line, hideFails)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[-] Error reading file:", err)
	}
}
