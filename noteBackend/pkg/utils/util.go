package utils

func GetStringOrDefault(str *string, defaultVal string) string {
	if str != nil {
		return *str
	}
	return defaultVal
}
