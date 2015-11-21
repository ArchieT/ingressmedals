package parsing

import "os/exec"
import "io/ioutil"
import "log"
import "bytes"
import "regexp"
import "strconv"
import "log"
import "os"

var DebugTrace = log.New(os.Stderr, "DEBUGOCR ", log.LstdFlags)

func Ocradout(filename string) []byte {
	c := exec.Command("png2pnm", filename)
	b := exec.Command("ocrad", "-i", "--scale=-1")
	pipe, err := b.StdinPipe()
	if err != nil {
		DebugTrace.Fatal(err)
	}
	c.Stdout = pipe
	bp, err := b.StdoutPipe()
	if err != nil {
		DebugTrace.Fatal(err)
	}
	err = c.Start()
	if err != nil {
		DebugTrace.Fatal(err)
	}
	bb, err := ioutil.ReadAll(bp)
	if err != nil {
		DebugTrace.Fatal(err)
	}
	return bb
}

func OcradAlterProc(filename string, czasowe bool) map[string]uint32 {
	b := Ocradout(filename)
	c := bytes.Split(b, []byte{byte(10)})
	if czasowe {
		o := sortedbypositionczasowe
	} else {
		o := sortedbypositionnow
	}
	elements := make(map[string][]byte)
	for _, j := range o {
		fullr := regexp.MustCompile(origstrsdictbef[j] + `.+` + origstrsdictaft[j] + `$`)
		for ln := range c {
			if _, ok := elements[j]; ok {
				break
			}
			go func() {
				found := fullr.Find(ln)
				if found != nil {
					elements[j] = found
				}
			}()
		}
	}
	for _, j := range o {
		if _, ok := elements[j]; ok {
			continue
		}
		switch j {
		case "seer", "destr", "destrlink", "destrfield", "neutr":
			DebugTrace.Println("Not found: ", j)
		default:
			DebugTrace.Fatal("NOT FOUND: ", j)
		}
	}
	elementojn := make(map[string]uint32)
	for i := range elements {
		elementojn[i] = Extract(elements[i], i)
	}
	return elementojn
}
