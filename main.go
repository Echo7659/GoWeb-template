package main

import "github.com/Echo7659/elog"

// init
func init(){

}

func main() {
	elog.Say.Info("hello world")
	elog.Say.Errorf("%s", "hello world")
	elog.Say.Debug("hello world")
}
