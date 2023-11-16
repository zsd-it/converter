package main

import (
	"fmt"

	"github.com/zsd-it/converter"
)

func main() {
	t2t := converter.NewTable2Struct()

	err := t2t.
		SavePath("/Users/shoudongzhao/data/wwwroot/go/tool/converter/dao/a.go").
		Table("activity").
		EnableJsonTag(true).
		Run()

	fmt.Println(err)
}
