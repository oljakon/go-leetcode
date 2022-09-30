package main

func generateParenthesis(n int) []string {
	res := []string{}
	pars := ""

	if n == 0 {
		return res
	} else if n == 1 {
		return []string{"()"}
	}

	var recursion func(opened, closed int)
	recursion = func(opened, closed int) {
		if n == opened && n == closed {
			res = append(res, pars)
			return
		}

		if opened < n {
			pars = pars + "("
			recursion(opened+1, closed)
			pars = pars[:len(pars)-1]
		}

		if closed < opened {
			pars = pars + ")"
			recursion(opened, closed+1)
			pars = pars[:len(pars)-1]
		}
	}

	recursion(0, 0)
	return res
}
