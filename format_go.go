package gostyle

import (
	"go/format"
	"io/ioutil"
)

// GoFormat 原地格式化go代码
func GoFormat(fpath string) error {

	in, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	out, err := format.Source(in)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fpath, out, 0644)
	if err != nil {
		return err
	}

	return nil
}
