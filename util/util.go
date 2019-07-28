package util

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

/*
func StringToDate(unparsed string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, unparsed)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(t.Format("20060102150405"))
	return t
}*/

func StringToDate(unparsed string) time.Time {
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, unparsed)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(t.Format("20060102150405"))
	return t
}

// Returns the string of strconv.Atoi. returns the max int when fails
func BetterAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("failed to parse string to int with error:")
		log.Println(err)
		return ^int(0)
	}
	return i
}
