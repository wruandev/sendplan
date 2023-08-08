package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func GetUserPlanInputString(reader *bufio.Reader, name string) (string, error) {
	fmt.Printf("    > Enter %s : ", name)

	companyName, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	sanitizedStr := strings.TrimSuffix(companyName, "\n")

	sanitizedStr = strings.TrimSuffix(sanitizedStr, "\r")

	return sanitizedStr, nil
}
