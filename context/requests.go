package context

import "fmt"

type Authorized struct {
	Uid int `json:"uid"`
}

func (r Authorized) GetUid() (int, error) {
	if r.Uid == 0 {
		return 0, fmt.Errorf("No uid")
	}
	return r.Uid, nil
}
