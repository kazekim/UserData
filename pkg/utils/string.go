/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package utils

import "strconv"

func StringToInt64(value string) (int64, error) {
	v, err := strconv.Atoi(value)
	return int64(v), err
}
