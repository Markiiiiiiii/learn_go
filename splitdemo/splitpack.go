package splitdemo

import "strings"

func Split(s, b string) []string {
	var strTmp []string
	n := strings.Index(s, b)
	for n >= 0 {
		strTmp = append(strTmp, s[:n])
		s = s[n+1:]
		n = strings.Index(s, b)
	}
	strTmp = append(strTmp, s)
	return strTmp
}
