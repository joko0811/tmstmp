package pkg

import (
	"fmt"
	"time"
)

func Generate_hhmm() string {
	t := time.Now()
	hhmm := "[" + fmt.Sprintf("%02d", t.Hour()) + ":" + fmt.Sprintf("%02d", t.Minute()) + "]"
	return hhmm
}
