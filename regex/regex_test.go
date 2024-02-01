package regeex

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	const cookie = "innersign=0; buvid3=BFAF5EB3-D834-F5A5-90BE-21C001704E2D99829infoc; b_nut=1703912299; i-wanna-go-back=-1; b_ut=7; b_lsid=E310F8C25_18CB9165C73; _uuid=C81D194A-21019-9DA3-1C37-57104694DF410E99659infoc; enable_web_push=DISABLE; header_theme_version=undefined; buvid4=59F17ECC-43F5-4640-8F9C-B6FE7D5D347E00770-023123004-UmzC4S68Duu8mys8uw8Zfg%3D%3D; home_feed_column=5; buvid_fp=189647ed78e8cc306076bba7f9aee9c7; SESSDATA=b8ca5887%2C1719464339%2C30cd5%2Ac1CjBB-qwMRq18PjjNwjStZSRdcp0e30c6Yr_uw_5K12Xi6ZiPTQnivQZypI9jarBl3AgSVjRNMzd4RlM0ZnBDYVpXbUItNzVTQVZqUGFwYi00UV9FaUtocGlMbUM1OHRQTFktNm05aDl1dU1UM0tlQ2VqbFhHRW1kZ0Rybk43MGR6Y2NqazBRTDlRIIEC; bili_jct=412ce6ac578e549ee90b1f404cf2521b; DedeUserID=349706324; DedeUserID__ckMd5=07528837d466bb20; sid=5juui2ex; CURRENT_FNVAL=4048; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQxNzE1NTksImlhdCI6MTcwMzkxMjI5OSwicGx0IjotMX0.sjEnDL5GTkRWc_t2ItzoXEXjR-QEQQzWyVgg9lj-TUk; bili_ticket_expires=1704171499; browser_resolution=1920-422"
	re := regexp.MustCompile("_uuid=(.+?);")
	res := re.FindAllStringSubmatch(cookie, -1)
	fmt.Println(res[0][1], "buvid")
}
