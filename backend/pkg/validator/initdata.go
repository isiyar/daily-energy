package validator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetTelegramUserID(initData string, botToken string) (string, error) {
	if initData == "" {
		return "", errors.New("initData is empty")
	}

	params, err := url.ParseQuery(initData)
	if err != nil || !validateInitData(params, botToken) {
		return "", errors.New("invalid telegram initData")
	}

	var user struct{ ID int64 }
	if err := json.Unmarshal([]byte(params.Get("user")), &user); err != nil {
		return "", errors.New("invalid user data")
	}

	return strconv.FormatInt(user.ID, 10), nil
}

func validateInitData(params url.Values, botToken string) bool {
	if authDate, err := strconv.ParseInt(params.Get("auth_date"), 10, 64); err == nil {
		if time.Since(time.Unix(authDate, 0)) > 10*time.Minute {
			return false
		}
	}

	hash := params.Get("hash")
	if params.Get("hash") == "" {
		return false
	}

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(botToken))

	var dataCheck strings.Builder
	keys := make([]string, 0, len(params)-1)
	
	for k := range params {
		if k != "hash" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		dataCheck.WriteString(k + "=" + params.Get(k) + "\n")
	}
	dataCheckStr := strings.TrimSuffix(dataCheck.String(), "\n")

	h := hmac.New(sha256.New, secret.Sum(nil))
	h.Write([]byte(dataCheckStr))
	return hex.EncodeToString(h.Sum(nil)) == hash
}