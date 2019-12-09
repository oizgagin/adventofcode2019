package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%q: os.Open: %v", filename, err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("%q: ioutil.ReadAll: %v", filename, err)
	}

	return strings.Split(string(content), "\n"), nil
}
