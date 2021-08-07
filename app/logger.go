package app

import (
	"errors"
	"log"
	"strings"
	"time"
)

func Logger(data ...interface{}) {
	log.Println(flagInfo(), "-:-", data)
}

func Success(data ...interface{}) {
	log.Println(flagInfo(), "-->", data)
}

func Error(e error) error {
	log.Println(flagErr(), "<->", e.Error())
	return e
}

func NewError(ds ...string) (e error) {
	d := ""
	for _, s := range ds {
		d += s + " "
	}
	e = errors.New(strings.Trim(d, " "))
	log.Println(flagErr(), d)
	return
}

func NewReturnError(ds ...string) (error, interface{}) {
	d := ""
	for _, s := range ds {
		d += s + " "
	}
	e := errors.New(strings.Trim(d, " "))
	log.Println(flagErr(), d)
	return e, nil
}

func NewErrors(d ...string) (es []error) {
	for _, s := range d {
		e := errors.New(s)
		es = append(es, e)
		log.Println(flagErr(), s)
	}
	return
}

func flagErr() string {
	return "[ERR-" + time.Now().Format("20060102-150405.0000") + "]"
}

func flagInfo() string {
	return "[INF-" + time.Now().Format("20060102-150405.0000") + "]"
}

func flagSuccess() string {
	return "[SUC-" + time.Now().Format("20060102-150405.0000") + "]"
}
