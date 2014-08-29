package main

import (
	"fmt"
	"io/ioutil"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Args struct {
		Dest string
	} `positional-args:"yes" required:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}
	dest := opts.Args.Dest
	defer fmt.Print(dest)

	src := dest + ".env"

	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	conf := os.ExpandEnv(string(bytes))

	err = ioutil.WriteFile(dest, []byte(conf), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
