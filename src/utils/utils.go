package utils

import "os"

func GetEnv(key string,defval string)string{
	keyVal := os.Getenv("S3_BUCKET")
	if keyVal == ""{
		return defval
	}
	return keyVal
}
