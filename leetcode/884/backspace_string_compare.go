package main

func backspaceCompare(S string, T string) bool {
	if S == "" && T == "" {
		return true
	}

	if S == "" || T == "" {
		return false
	}

	ns, nt := "", ""
	for _, ss := range S {
		if ss == '#' {
			if len(ns) > 0 {
				ns = ns[:len(ns)-1]
			}
		} else {
			ns += string(ss)
		}
	}

	for _, tt := range T {
		if tt == '#' {
			if len(nt) > 0 {
				nt = nt[:len(nt)-1]
			}
		} else {
			nt += string(tt)
		}
	}

	return ns == nt
}
