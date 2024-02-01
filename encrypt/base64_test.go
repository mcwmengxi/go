package encrypt

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	// aGVsbG8=
	encodeString := base64.StdEncoding.EncodeToString([]byte("hello"))
	fmt.Println(encodeString)
	bt := make([]byte, 10)
	base64.StdEncoding.Encode(bt, []byte("hello"))
	fmt.Println(bt)
	decodeStr, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Println(string(decodeStr))
}
