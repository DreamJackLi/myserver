package uid

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetStringUUID(extendInfo string) string {

	uid := fmt.Sprintf("%d%s", time.Now().UnixNano(), extendInfo)
	h := md5.Sum([]byte(uid))
	s := hex.EncodeToString(h[:])
	u, err := uuid.FromString(s)

	if err != nil {
		fmt.Println("GetStringUUID err is ", err)
		return "GetStringUUID"
	}
	return u.String()
}
