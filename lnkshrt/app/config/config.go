package config

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const portDefault = 8080
const idLengthDefault = 6
const persistentDefault = false
const dbfileDefault = "data/lnkshrt.json"
const indexRedirectDefault = ""

var originsDefault = []string{}

func Executable() string {
	path, _ := os.Executable()
	return filepath.Base(path)
}

func Port() int {
	port, found := os.LookupEnv("LNKSHRT_PORT")

	if !found {
		return portDefault
	}

	num, err := strconv.Atoi(port)

	if err != nil || num < 1 || num > 65535 {
		return portDefault
	}

	return num
}

func IdLength() int {
	idsize, found := os.LookupEnv("LNKSHRT_IDSIZE")

	if !found {
		return idLengthDefault
	}

	num, err := strconv.Atoi(idsize)

	if err != nil || num < 1 || num > 128 {
		return idLengthDefault
	}

	return num
}

func Persistent() bool {
	if persistent, found := os.LookupEnv("LNKSHRT_PERSISTENT"); found {
		val, err := strconv.ParseBool(persistent)
		if err != nil {
			val = true
		}
		return val
	}

	return persistentDefault
}

func DBFile() string {
	dbfile, found := os.LookupEnv("LNKSHRT_DBFILE")

	if !found || dbfile == "" {
		return dbfileDefault
	}

	return dbfile
}

func Origins() []string {
	if origins, found := os.LookupEnv("LNKSHRT_ORIGINS"); found {
		return strings.Split(origins, ",")
	}

	return originsDefault
}

func IndexRedirect() string {
	if indexRedirect, found := os.LookupEnv("LNKSHRT_INDEX_REDIRECT"); found {
		return indexRedirect
	}

	return indexRedirectDefault
}
