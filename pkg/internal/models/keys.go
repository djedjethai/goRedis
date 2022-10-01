package models

import "fmt"

// type Keys struct{}

func PageCacheKey(k string) string {
	return fmt.Sprintf("pagecache#%s", k)
}

func UserIDKey(uid string) string {
	return fmt.Sprintf("users#%s", uid)
}

func SessionKey(sid string) string {
	return fmt.Sprintf("sessions#%s", sid)
}
