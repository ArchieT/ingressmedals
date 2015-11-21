package parsing

import "os/exec"
import "io/ioutil"
import "log"
import "bytes"

func Ocradout(filename string) []byte {
	c := exec.Command("png2pnm", filename)
	b := exec.Command("ocrad", "-i", "--scale=-1")
	pipe, err := b.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	c.Stdout = pipe
	bp, err := b.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Start()
	if err != nil {
		log.Fatal(err)
	}
	bb, err := ioutil.ReadAll(bp)
	if err != nil {
		log.Fatal(err)
	}
	return bb
}

func OcradAlterProc(filename string, czasowe bool) {
	b := Ocradout(filename)
	c := bytes.Split(b, []byte{byte(10)})
	if czasowe {
		o := sortedbypositionczasowe
	} else {
		o := sortedbypositionnow
	}
	//for
}
