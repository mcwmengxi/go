package encrypt

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {

	// 字符串进行转义
	// https%3A%2F%2Fwww.bilibili.com%2Fvideo%2FBV1cV41197NK%2F%3Fspm_id_from%3D333.1007.tianma.4-1-11.click
	queryEncode := url.QueryEscape("https://www.bilibili.com/video/BV1cV41197NK/?spm_id_from=333.1007.tianma.4-1-11.click")
	log.Println(queryEncode)
}

func TestUrlDecode(t *testing.T) {
	encode := url.QueryEscape("https://www.bilibili.com/video/BV1cV41197NK/?spm_id_from=333.1007.tianma.4-1-11.click")
	fmt.Println("encode:  " + encode)
	decode, _ := url.QueryUnescape(encode)
	fmt.Println("decode:  " + decode)
}
