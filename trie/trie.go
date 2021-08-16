package trie

func New() *Trie {
	return &Trie{root: &node{children: map[rune]*node{}}}
}

type Trie struct {
	root *node
}

func (t *Trie) Add(key, value string) {
	t.root.add([]rune(key), value)
}

func (t *Trie) Get(key string) string {
	return t.root.get([]rune(key))
}

func (t *Trie) Optimise() {
	t.root.optimise()
}

func (t *Trie) Size() int {
	return t.root.nodes()
}

func (t *Trie) AsMap() map[string]string {
	m := map[string]string{}
	t.root.untrie("", m)
	return m
}

type node struct {
	children map[rune]*node
	value    string
}

func (curr *node) add(parts []rune, value string) {
	part := parts[0]
	parts = parts[1:]

	child, ok := curr.children[part]
	if !ok {
		child = &node{children: map[rune]*node{}}
		curr.children[part] = child
	}

	if len(parts) > 0 {
		child.add(parts, value)
	} else {
		child.value = value
	}
}

func (curr *node) get(parts []rune) string {
	if len(parts) == 0 {
		return curr.value
	}

	child, ok := curr.children[parts[0]]
	if !ok {
		return curr.value
	}

	return child.get(parts[1:])
}

func (curr *node) nodes() int {
	total := 1

	for _, child := range curr.children {
		total += child.nodes()
	}

	return total
}

func (curr *node) optimise() {
	for _, child := range curr.children {
		child.optimise()
	}

	if val, ok := curr.same(); ok {
		curr.children = map[rune]*node{}
		curr.value = val
	}
}

func (curr *node) same() (string, bool) {
	if curr.value != "" {
		return curr.value, true
	}

	var val string
	for _, child := range curr.children {
		childVal, ok := child.same()
		if !ok {
			return "", false
		}

		if val == "" || val == childVal {
			val = childVal
		} else {
			return "", false
		}
	}

	return val, true
}

func (curr *node) untrie(root string, o map[string]string) {
	for key, child := range curr.children {
		next := root + string(key)
		if child.value != "" {
			o[next] = child.value
		} else {
			child.untrie(next, o)
		}
	}
}
