package parsing

import "regexp"

var sortedbypositionczasowe = []string{
	`ap`,
	`uniqvis`,
	`seer`,
	`xm`,
	`walk`,
	`depl`,
	`link`,
	`field`,
	`mu`,
	`longestlink`,
	`largestfield`,
	`rech`,
	`capt`,
	`uniqcapt`,
	`mods`,
	`destr`,
	`neutr`,
	`destrlink`,
	`destrfield`,
	`guard`,
	`guardlink`,
	`longxguardlink`,
	`guardfield`,
	`longxguardfield`,
	`uniqmis`,
	`hack`,
	`glyph`,
}

var sortedbypositionnow = []string{"ap", "linksactiv", "pwned", "fieldsactiv", "mucontrol"}

var origstrsdictbef = map[string]*regexp.Regexp{
	"ap":                     regexp.MustCompile(`^[_-]{0,3} ?`),
	"uniqvis":                regexp.MustCompile(`Un[il]qu[eP] Por[tf]als [Vv][il]s[il][tl][Pe]d `),
	"seer":                   regexp.MustCompile(`Por[tf]als [DO][il]s[cI]ov[eP]r[eP]d `), //"Porfals OlsIovPrPd l"
	"xm":                     regexp.MustCompile(`XM Coll[eP][cI][tI][eP]d `),
	"hack":                   regexp.MustCompile(`Ha[cI]ks `),
	"depl":                   regexp.MustCompile(`R[eP]sona[tl]ors [DO][eP]ploy[eP]d `),
	"link":                   regexp.MustCompile(`L[il]nks Cr[eP]a[tl][eP]d `),
	"field":                  regexp.MustCompile(`Con[tl]rol F[iI][eP]lds Cr[eP]a[tl][eP]d `),
	"mu":                     regexp.MustCompile(`M[il]nd Un[il][tl]s Cap[tl]ur[eP]d `),
	"longestlink":            regexp.MustCompile(`Long[eP]s[tl] L[il]nk Ev[eP]r Cr[eP]a[tl][eP]d `),
	"largestfield":           regexp.MustCompile(`Larg[eP]s[tl] Con[tl]rol F[iI][eP]ld `),
	"rech":                   regexp.MustCompile(`XM R[eP][cI]harg[eP]d `),
	"capt":                   regexp.MustCompile(`^Por[tlf]als Cap[tl]ur[eP]d `),
	"uniqcapt":               regexp.MustCompile(`Un[il]qu[eP] Por[tlf]als Cap[tl]ur[eP]d `),
	"destregexp.MustCompile": regexp.MustCompile(`R[eP]sona[tl]ors [DO][eP]s[tl]roy[eP]d `),
	"neutr":                  regexp.MustCompile(`Por[tlf]als N[eP]u[tl]ral[il][zI_][eP]d `), //'Porlals NPulrall_Pd IBT'
	"destrlink":              regexp.MustCompile(`En[eP]my L[il]nks [DO][eP]s[tl]roy[eP]d `),
	"destrfield":             regexp.MustCompile(`En[eP]my Con[tl]rol F[iI][eP]lds [DO][eP]s[tl]roy[eP]d `),
	"walk":                   regexp.MustCompile(`[DO][il]s[tl]a[nm]*[cm]*[eP] Walk[eP]d `), //'OlslamP WalkPd 9km'
	"guard":                  regexp.MustCompile(`Max T[il]m[eP] Por[tlf]al H[eP]ld `),      //'Max TlmP Porlal HPld 14 day5'
	"guardlink":              regexp.MustCompile(`Max T[il]m[eP] L[il]nk Ma[il]n[tl]a[il]n[eP]d `),
	"longxguardlink":         regexp.MustCompile(`Max L[il]nk L[eP]ng[tl]h x [DO]ays `),
	"guardfield":             regexp.MustCompile(`Max T[il]m[eP] Fi[eP]ld H[eP]ld `),
	"longxguardfield":        regexp.MustCompile(`Larg[eP]s[tl] F[iI][eP]ld MUs x [DO]ays `),
	"mods":                   regexp.MustCompile(`Mods [DO][eP]ploy[eP]d `),
	"uniqmis":                regexp.MustCompile(`Un[il]qu[eP] m[il][sS5][sS5][il]on[sS5] completed `),
	"glyph":                  regexp.MustCompile(`Glyph Hack Points `),
	"linksactiv":             regexp.MustCompile(`L[il]nks Active `),
	"pwned":                  regexp.MustCompile(`Por[tlf]als Owned `),
	"fieldsactiv":            regexp.MustCompile(`Con[tl]rol F[iI][eP]lds Active `),
	"mucontrol":              regexp.MustCompile(`M[il]nd Un[il][tl] Control `),
}

var origstrsdictaft = map[string]*regexp.Regexp{
	"ap":              regexp.MustCompile(`[ ]?A[Pp][ \|Il_]+.+A[Pp]`),
	"xm":              regexp.MustCompile(` XM`),
	"mu":              regexp.MustCompile(`MUs`),
	"longestlink":     regexp.MustCompile(`km`),
	"largestfield":    regexp.MustCompile(`MUs`),
	"rech":            regexp.MustCompile(` XM`),
	"walk":            regexp.MustCompile(`km`),
	"guard":           regexp.MustCompile(` day[S5s]`),
	"guardlink":       regexp.MustCompile(` day[S5s]`),
	"longxguardlink":  regexp.MustCompile(` km-day[s5]`),
	"guardfield":      regexp.MustCompile(` days`),
	"longxguardfield": regexp.MustCompile(` MU-day[s5]`),
	"mucontrol":       regexp.MustCompile(`MUs`),
}
