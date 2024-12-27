package main

import "sort"

type website struct {
	page   string
	visits int
}

func sortPages(pages map[string]int) []website {

	reversed := make([]website, 0)

	for k, v := range pages {
		reversed = append(reversed, website{k, v})

	}
	sort.Slice(reversed, func(a, b int) bool { return reversed[a].visits >= reversed[b].visits })

	return reversed

}
