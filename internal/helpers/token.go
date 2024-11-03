package helpers

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

type PayloadRefreshToken struct {
	UserId    uint
	Email     string
	ExpiredAt time.Time
}

func CreateRefreshToken(payload PayloadRefreshToken) (string, error) {
	encodePayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encodePayload), nil
}

func DecodeRefreshToken(src string) (PayloadRefreshToken, error) {
	var payload PayloadRefreshToken

	decode, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return payload, err
	}

	if err := json.Unmarshal(decode, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
