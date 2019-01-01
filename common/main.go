package common

import (
	"fmt"
	"time"
)

func DatePrint(message string) {

	fmt.Println(fmt.Sprintf("%v %s", time.Now(), message))
}
