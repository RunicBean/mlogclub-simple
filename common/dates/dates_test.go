package dates_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RunicBean/mlogclub-simple/common/dates"
)

func TestWithTimeAsEndOfDay(t *testing.T) {
	fmt.Println(dates.Timestamp(dates.WithTimeAsEndOfDay(time.Now())))
}
