package parsing

import "os/exec"
import "io/ioutil"
import "log"

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
