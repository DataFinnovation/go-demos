package apidemos

import "os"

// Getenv returns the defaultValue if the variable is not set
func Getenv(envVarName, defaultValue string) string {
	res, present := os.LookupEnv(envVarName)
	if present {
		return res
	}
	return defaultValue
}

// eof
