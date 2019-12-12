package run

import "strings"

type Tree map[string]string

func parseTree(in chan string) Tree {
	tree := make(Tree, 10)
	for i := range in {
		v := strings.Split(i, ")")
		tree[v[1]] = v[0]
	}
	return tree
}

func countParents(p string, tree Tree) int {
	if val, ok := tree[p]; ok {
		return countParents(val, tree) + 1
	} else {
		return 0
	}
}

func listParents(p string, tree Tree) []string {
	if val, ok := tree[p]; ok {
		return append(listParents(val, tree), val)
	} else {
		return make([]string, 0)
	}
}

func Six(in chan string) interface{} {
	tree := parseTree(in)
	total := 0
	for node, _ := range tree {
		total += countParents(node, tree)
	}
	return total
}

func Six2(in chan string) interface{} {
	tree := parseTree(in)
	left := listParents("SAN", tree)
	right := listParents("YOU", tree)
	c := 0
	for _, l := range left {
		for _, r := range right {
			if l == r {
				c++
			}
		}
	}
	return len(left) + len(right) - (2*c)
}