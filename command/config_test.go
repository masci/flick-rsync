package command

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestGetConfigFilePath(t *testing.T) {
	p := getConfigFilePath()
	if p == "" {
		t.Error("Config file path is empty")
	}

	dir, file := path.Split(p)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		t.Error("Config file dir does not exists:", dir)
	}

	if file != ".flick-rsync.cfg" {
		t.Error("Unexpected config file name:", file)
	}
}

func TestParseConfigFile(t *testing.T) {
	jsonData := `
{
	"api_key":"foo",
	"api_secret":"bar"
}`

	bogusJsonData := `
{
	"api_key":"foo",
	"api_secret":"bar"`

	file, err := ioutil.TempFile("", "flickrsynctst")
	if err != nil {
		t.Error("Unable to write temp file")
	}

	_, err = file.WriteString(jsonData)
	if err != nil {
		t.Error("Unable to write test data to file", file.Name())
	}

	out, err := parseConfigFile(file.Name())
	if err != nil {
		t.Error("Unable to parse config file", file.Name())
	}

	if out.ApiKey != "foo" || out.ApiSecret != "bar" {
		t.Error("Error parsing config file, expected foo bar, found", out.ApiKey, out.ApiSecret)
	}

	out, err = parseConfigFile("")
	if err == nil {
		t.Error("Invoking with empty param should return an error")
	}

	if out != nil {
		t.Error("Invoking with empty param should return nil, found", out)
	}

	file, err = ioutil.TempFile("", "flickrsynctst")
	if err != nil {
		t.Error("Unable to write temp file")
	}

	_, err = file.WriteString(bogusJsonData)
	if err != nil {
		t.Error("Unable to write test data to file", file.Name())
	}

	out, err = parseConfigFile(file.Name())
	if err == nil {
		t.Error("Should not parse bogus config file", file.Name())
	}
}
