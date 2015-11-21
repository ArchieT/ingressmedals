package parsing

import "regexp"
import "strconv"

type univfilter struct {
	*regexp.Regexp
	string
}

var ufs = []univfilter{
	{regexp.MustCompile(`[.]|,|\s/`), ""},
	{regexp.MustCompile(`o|O`), "0"},
	{regexp.MustCompile(`[\|]|l|i|I|\|`), "1"},
	{regexp.MustCompile(`q`), "4"},
	{regexp.MustCompile(`t|T`), "7"},
	{regexp.MustCompile(`a|e|B`), "8"},
	{regexp.MustCompile(`g`), "9"},
	{regexp.MustCompile(`S|s`), "6"},
	{regexp.MustCompile(`z|Z`), "2"},
	{regexp.MustCompile(`n`), "77"},
	{regexp.MustCompile(`\D`), ""},
}

func UnivFilter(b []byte) []byte {
	c := b
	for i := range ufs {
		c := ufs[i].Regexp.ReplaceAll(c, ufs[i].string)
	}
	return c
}

func Filter(b []byte, elem string) []byte {
	c := regexp.MustCompile(origstrsdictbef[elem]).ReplaceAll(b, "")
	c = regexp.MustCompile(origstrsdictaft[elem]).ReplaceAll(c, "")
	return UnivFilter(c)
}

func IntExtraction(b []byte) uint32 {
	return uint32(strconv.ParseUint(string(b), 10, 32))
}

func Extract(b []byte, elem string) uint32 {
	return IntExtraction(Filter(b, elem))
}
