package validator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type InitData struct {
	TelegramID string
}

type BotSettings struct {
	Token string
}

func GetTelegramInitData(ctx *gin.Context, botToken string) (*InitData, error) {
	initdata := ctx.Param("initData")
	if initdata == "" {
		return nil, errors.New("invalid telegram initdata")
	}

	if ok, err := ValidateInitData(initdata, botToken); !ok || err != nil {
		return nil, errors.New("invalid telegram initdata")
	}

	params, _ := url.ParseQuery(initdata)
	userData := params.Get("user")

	var user struct{ ID int64 }
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return nil, errors.New("invalid telegram initdata")
	}

	return &InitData{TelegramID: strconv.FormatInt(user.ID, 10)}, nil
}

func ValidateInitData(initData string, botToken string) (bool, error) {
	params, err := url.ParseQuery(initData)
	if err != nil {
		return false, err
	}

	if authDate, err := strconv.ParseInt(params.Get("auth_date"), 10, 64); err == nil {
		if time.Since(time.Unix(authDate, 0)) > 10*time.Minute {
			return false, nil
		}
	}

	receivedHash := params.Get("hash")
	if receivedHash == "" {
		return false, nil
	}

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(botToken))
	secretKey := secret.Sum(nil)

	var checkStr strings.Builder
	params.Del("hash")

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		checkStr.WriteString(k)
		checkStr.WriteByte('=')
		checkStr.WriteString(params.Get(k))
		checkStr.WriteByte('\n')
	}
	dataCheckStr := strings.TrimSuffix(checkStr.String(), "\n")

	hash := hmac.New(sha256.New, secretKey)
	hash.Write([]byte(dataCheckStr))
	computedHash := hex.EncodeToString(hash.Sum(nil))

	return computedHash == receivedHash, nil
}
