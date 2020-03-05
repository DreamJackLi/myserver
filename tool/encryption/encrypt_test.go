package encryption

import (
	"fmt"
	"testing"
)

func TestEn(t *testing.T) {
	str := "qwertyuiopasdfgh"
	res, tim := AuthCode(str, "", true, 5)

	fmt.Println(res, tim)

	s, _ := AuthCode(res, "", false, 5)

	fmt.Println(s)
	// fmt.Println(md5Str(nil, 32))
}
