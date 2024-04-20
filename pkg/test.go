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

	g.buildP("N::=ND|D")
	g.buildP("D::=0|1|2|3|4|5|6|7|8|9")

	return g
}
