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

func ItemsKey(iid string) string {
	return fmt.Sprintf("session#%s", iid)
}

func UsernamesUniqueKey() string {
	return "username:unique"
}

func UserLikesKey(uid string) string {
	return fmt.Sprintf("users:likes#%v", uid)
}
