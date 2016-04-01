package main

import (
	"encoding/xml"
	"errors"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func createV2Notice(c *gin.Context) {
	var (
		err, readErr error
		body         []byte
		app          App
	)

	notice := v2Notice{}
	if body, readErr = ioutil.ReadAll(c.Request.Body); readErr != nil {
		returnError(c, errors.New("Couldn't read body: "+err.Error()))
		return
	}
	if err = xml.Unmarshal(body, &notice); err != nil {
		returnError(c, errors.New("Couldn't unmarshal xml: "+err.Error()))
		return
	}

	for _, app = range apps {
		if notice.ApiKey == app.Key {
			notice.App = app.Name
		}
	}

	if app.Key == "" {
		returnError(c, errors.New("Not authorized"))
		return
	}

	t := time.Now()
	var dir = config.Dir + "/" + strconv.Itoa(t.Year()) + "/" + strconv.Itoa(int(t.Month())) + "/" + strconv.Itoa(t.Day())
	if err = checkAndCreateDir(dir); err != nil {
		returnError(c, err)
		return
	}

	y, err := yaml.Marshal(notice)
	if err != nil {
		returnError(c, err)
		return
	}

	fi, err := os.OpenFile(dir+"/"+app.Name+".log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		returnError(c, err)
		return
	}
	defer fi.Close()

	rec := "================" + time.Now().Format("2006-01-02T15:04:05Z07:00") + "==================\n" + string(y) + "\n"
	io.WriteString(fi, rec)

	return
}
