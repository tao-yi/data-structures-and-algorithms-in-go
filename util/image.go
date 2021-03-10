package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/bradleyjkemp/memviz"
)

func MakeImg(name string, data interface{}) {
	buf := &bytes.Buffer{}
	memviz.Map(buf, data)
	err := ioutil.WriteFile(name, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("dot", "-Tpng", name, "-o", name+".png")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\nreason:%s", err, string(out))
	}
	fmt.Println(string(out))
	_ = os.Remove(name)
}
