package utils

import "os"

func GetEnv(key string,defval string)string{
	keyVal := os.Getenv(key)
	if keyVal == ""{
		return defval
	}
	return keyVal
}
