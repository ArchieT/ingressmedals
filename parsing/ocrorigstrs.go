package parsing

//import "regexp"

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

var origstrsdictbef = map[string]string{
	"ap":              `^[_-]{0,3} ?`,
	"uniqvis":         `Un[il]qu[eP] Por[tf]als [Vv][il]s[il][tl][Pe]d `,
	"seer":            `Por[tf]als [DO][il]s[cI]ov[eP]r[eP]d `, //"Porfals OlsIovPrPd l"
	"xm":              `XM Coll[eP][cI][tI][eP]d `,
	"hack":            `Ha[cI]ks `,
	"depl":            `R[eP]sona[tl]ors [DO][eP]ploy[eP]d `,
	"link":            `L[il]nks Cr[eP]a[tl][eP]d `,
	"field":           `Con[tl]rol F[iI][eP]lds Cr[eP]a[tl][eP]d `,
	"mu":              `M[il]nd Un[il][tl]s Cap[tl]ur[eP]d `,
	"longestlink":     `Long[eP]s[tl] L[il]nk Ev[eP]r Cr[eP]a[tl][eP]d `,
	"largestfield":    `Larg[eP]s[tl] Con[tl]rol F[iI][eP]ld `,
	"rech":            `XM R[eP][cI]harg[eP]d `,
	"capt":            `^Por[tlf]als Cap[tl]ur[eP]d `,
	"uniqcapt":        `Un[il]qu[eP] Por[tlf]als Cap[tl]ur[eP]d `,
	"destr":           `R[eP]sona[tl]ors [DO][eP]s[tl]roy[eP]d `,
	"neutr":           `Por[tlf]als N[eP]u[tl]ral[il][zI_][eP]d `, //'Porlals NPulrall_Pd IBT'
	"destrlink":       `En[eP]my L[il]nks [DO][eP]s[tl]roy[eP]d `,
	"destrfield":      `En[eP]my Con[tl]rol F[iI][eP]lds [DO][eP]s[tl]roy[eP]d `,
	"walk":            `[DO][il]s[tl]a[nm]*[cm]*[eP] Walk[eP]d `, //'OlslamP WalkPd 9km'
	"guard":           `Max T[il]m[eP] Por[tlf]al H[eP]ld `,      //'Max TlmP Porlal HPld 14 day5'
	"guardlink":       `Max T[il]m[eP] L[il]nk Ma[il]n[tl]a[il]n[eP]d `,
	"longxguardlink":  `Max L[il]nk L[eP]ng[tl]h x [DO]ays `,
	"guardfield":      `Max T[il]m[eP] Fi[eP]ld H[eP]ld `,
	"longxguardfield": `Larg[eP]s[tl] F[iI][eP]ld MUs x [DO]ays `,
	"mods":            `Mods [DO][eP]ploy[eP]d `,
	"uniqmis":         `Un[il]qu[eP] m[il][sS5][sS5][il]on[sS5] completed `,
	"glyph":           `Glyph Hack Points `,
	"linksactiv":      `L[il]nks Active `,
	"pwned":           `Por[tlf]als Owned `,
	"fieldsactiv":     `Con[tl]rol F[iI][eP]lds Active `,
	"mucontrol":       `M[il]nd Un[il][tl] Control `,
}

var origstrsdictaft = map[string]string{
	"ap":              `[ ]?A[Pp][ \|Il_]+.+A[Pp]`,
	"xm":              ` XM`,
	"mu":              `MUs`,
	"longestlink":     `km`,
	"largestfield":    `MUs`,
	"rech":            ` XM`,
	"walk":            `km`,
	"guard":           ` day[S5s]`,
	"guardlink":       ` day[S5s]`,
	"longxguardlink":  ` km-day[s5]`,
	"guardfield":      ` days`,
	"longxguardfield": ` MU-day[s5]`,
	"mucontrol":       `MUs`,
}
