package serverutils

import (
	"fmt"
	"time"
)

func StringToDate(unparsed string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, unparsed)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(t.Format("20060102150405"))
	return t
}
