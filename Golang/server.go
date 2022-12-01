package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"
	"unsafe"

	"context"

	"github.com/go-redis/redis/v9"

	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

func main() {
	rand.Seed(time.Now().UnixNano())
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/html")
		return c.String(http.StatusOK, "<!DOCTYPE html><html lang=\"en\"><head> <meta charset=\"UTF-8\"> <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"> <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"> <title>FCUT</title></head><body><div style=\"text-align: center;\"> <form method=\"post\" action=\"/shorten\"> <label for=\"url\">URL:</label> <br/> <input id=\"url\" name=\"url\" type=\"url\" required/> <br/> <br/> <br/> <input type=\"submit\" value=\"SHORTEN\"/> </form></div></body></html>")
	})
	e.GET("/:short", func(c echo.Context) error {
		val, err := rdb.Get(ctx, c.Param("short")).Result()
		if err != nil {
			return c.String(http.StatusNotFound, "404: Not found")
		}

		return c.Redirect(http.StatusTemporaryRedirect, val)
	})
	e.POST("/shorten", func(c echo.Context) error {
		url := c.FormValue("url")
		generated := RandStringBytesMaskImprSrcUnsafe(8)

		err := rdb.Set(ctx, generated, url, 0).Err()
		if err != nil {
			generated = "error"
		}

		url = c.Request().Header.Get("Origin") + "/" + generated
		out := "<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>FCUT</title></head><div style=\"text-align: center\"><h1>GENERATED!</h1> <a href=\"" + url + "\">" + url + "</a></div>"

		c.Response().Header().Set("Content-Type", "text/html")
		return c.String(http.StatusOK, out)
	})

	e.Logger.Fatal(e.Start(":8003"))
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
