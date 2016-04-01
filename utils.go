package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sys/unix"
	"log"
	"net/http"
	"os"
)

func checkMime(mime string) gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.ContentType()

		if contentType != mime {
			returnError(c, errors.New("Incorrect mime-type in request: "+contentType))
		}
	}
}

func checkPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength == 0 {
			returnError(c, errors.New("Empty body from "+c.Request.RemoteAddr))
		}
	}
}

func returnError(c *gin.Context, err error) {
	log.Println(err.Error())
	c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
	c.Abort()
}

func checkAndCreateDir(path string) error {
	fi, err := os.Stat(path)

	switch {
	case err != nil:
		if os.IsNotExist(err) {
			if err = os.MkdirAll(path, 0755); err == nil {
				log.Printf("Created dir: %v", path)
				return nil
			}
		}
	case fi.IsDir():
		if unix.Access(path, unix.W_OK) == nil {
			return nil
		}

	}
	return errors.New("File error: " + err.Error())
}
