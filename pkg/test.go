package pkg

func TestCase() *Grammar {
	g := &Grammar{
		Name: 'G',
		Vn: map[byte]struct{}{
			'A': {},
			'B': {},
		},
		Vt:    make(map[byte]struct{}),
		S:     'N',
		Class: 3,
	}

	g.buildP("abc::=B")
	g.buildP("B::=bc")

	return g
}
