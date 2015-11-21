package parsing

import "regexp"
import "strconv"

type univfilter struct {
	*regexp.Regexp
	bytes []byte
}

var ufs = []univfilter{
	{regexp.MustCompile(`[.]|,|\s/`), []byte{}},
	{regexp.MustCompile(`o|O`), []byte("0")},
	{regexp.MustCompile(`[\|]|l|i|I|\|`), []byte("1")},
	{regexp.MustCompile(`q`), []byte("4")},
	{regexp.MustCompile(`t|T`), []byte("7")},
	{regexp.MustCompile(`a|e|B`), []byte("8")},
	{regexp.MustCompile(`g`), []byte("9")},
	{regexp.MustCompile(`S|s`), []byte("6")},
	{regexp.MustCompile(`z|Z`), []byte("2")},
	{regexp.MustCompile(`n`), []byte("77")},
	{regexp.MustCompile(`\D`), []byte{}},
}

//UnivFilter fixes most common ocrad digit reading errors
func UnivFilter(b []byte) []byte {
	c := b
	for i := range ufs {
		c = ufs[i].Regexp.ReplaceAll(c, ufs[i].bytes)
	}
	return c
}

func filter(b []byte, elem string) []byte {
	c := regexp.MustCompile(origstrsdictbef[elem]).ReplaceAll(b, []byte{})
	c = regexp.MustCompile(origstrsdictaft[elem]).ReplaceAll(c, []byte{})
	return UnivFilter(c)
}

func intExtraction(b []byte) uint32 {
	a, err := strconv.ParseUint(string(b), 10, 32)
	if err != nil {
		DebugTrace.Fatal(err)
	}
	return uint32(a)
}

//Extract gives you the uint32
func Extract(b []byte, elem string) uint32 {
	return intExtraction(filter(b, elem))
}
