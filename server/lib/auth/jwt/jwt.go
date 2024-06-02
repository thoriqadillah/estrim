package jwt

import (
	"fcompressor/env"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/patrickmn/go-cache"
)

type fastJwt struct {
	method *jwt.SigningMethodHMAC
	key    []byte
	cache  *cache.Cache
}

var jwtmethod = jwt.SigningMethodHS256
var exp = 10 * time.Minute
var memcache = cache.New(exp, exp*2)

func Encode(claims jwt.MapClaims) (token string, err error) {
	key := env.Get("JWT_SECRET").Bytes([]byte("secret"))
	encoder := jwt.NewWithClaims(jwtmethod, claims)

	token, err = encoder.SignedString(key)
	if err != nil {
		return "", err
	}

	memcache.Set(token, claims, exp)
	return token, nil
}

func Decode(token string) (jwt.MapClaims, error) {
	if claims, ok := memcache.Get(token); ok {
		return claims.(jwt.MapClaims), nil
	}

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("signing method invalid")
		}

		if method != jwtmethod {
			return nil, fmt.Errorf("signing method invalid")
		}

		key := env.Get("JWT_SECRET").Bytes([]byte("secret"))
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok {
		memcache.Set(token, claims, exp)
		return claims, nil
	}

	return nil, fmt.Errorf("token is not valid")
}
