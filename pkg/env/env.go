/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package env

import "os"

func GetValue(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
