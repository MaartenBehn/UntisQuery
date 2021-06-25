package Untis

import (
	"log"
	"strings"
	"time"
)

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return removeEmptiStrings(strings.FieldsFunc(s, splitter)...)
}

func removeEmptiStrings(strings ...string) []string {
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}

func checkError(err error) {
	if err != nil {
		log.Printf("An Error Occured %v\n", err)
	}
}

func TimeZone() *time.Location {
	return time.FixedZone("UTC+2", 2*60*60)
}
