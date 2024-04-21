package pkg

func TestCase() *Grammar {
	g := &Grammar{
		Name: 'S',
		Vn: map[byte]struct{}{
			'N': {},
			'D': {},
		},
		Vt:    make(map[byte]struct{}),
		S:     'N',
		Class: 3,
	}

	g.buildP("ab::=B")
	g.buildP("B::=bc")

	return g
}
