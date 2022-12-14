package handler

import "encoding/json"

func jsonError(msg string) []byte {
	errorS := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	r, err := json.Marshal(errorS)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
