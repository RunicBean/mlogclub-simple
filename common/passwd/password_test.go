package passwd_test

import (
	"fmt"
	"testing"

	"github.com/RunicBean/mlogclub-simple/common/passwd"
)

func TestPassword(t *testing.T) {
	fmt.Println(passwd.GenerateRandomPassword(16))
}
