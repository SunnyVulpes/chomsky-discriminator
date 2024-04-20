package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grammar struct {
	Name byte
	Vn   map[byte]struct{}
	Vt   map[byte]struct{}
	P    []Production
	S    byte
}

type Production struct {
	Left        byte
	Right       [][]byte
	OriginalStr string
}

func BuildGrammar() *Grammar {
	g := &Grammar{}
	g.Vn = make(map[byte]struct{})
	g.Vt = make(map[byte]struct{})
	g.P = make([]Production, 0)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("请输入文法，例如:'G[N]'")
	_, err := fmt.Scanf("%c[%c]\n", &g.Name, &g.S)
	if err != nil {
		exit(err.Error())
	}

	var vnStr string
	fmt.Println("请输入非终止符，例如:'N, D'(必须为大写字母)")
	vnStr, err = reader.ReadString('\n')
	vnStr = strings.TrimSpace(vnStr)
	g.buildVn(vnStr)

	var pStr string
	fmt.Println("请依次输入产生式规则，例如:'N::=ND|D'，输入 'end' 结束")
	for {
		pStr, err = reader.ReadString('\n')
		if err != nil {
			exit(err.Error())
		}
		pStr = strings.TrimSpace(pStr)

		if pStr == "end" {
			break
		}

		g.buildP(pStr)
	}

	return g
}

func (g *Grammar) buildVn(vnStr string) {
	hasReadComma := true
	for i := 0; i < len(vnStr); i++ {
		c := vnStr[i]
		if c == ' ' {
			continue
		}
		if c == ',' {
			if hasReadComma {
				exit("非终止符格式错误")
			} else {
				hasReadComma = true
			}
			continue
		}
		if isUpperCase(c) {
			hasReadComma = false
			if g.checkVn(c) {
				exit("重复输入非终止符")
			}
			g.Vn[c] = struct{}{}
			continue
		}

		exit("意外的输入")
	}
}

func (g *Grammar) buildP(pStr string) {
	cleanPStr := strings.ReplaceAll(pStr, " ", "")

	if len(cleanPStr) < 5 {
		exit("错误的产生式规则格式")
	}

	left := pStr[0]
	if !g.checkVn(left) {
		exit("产生式第一个字符应该是非终结符中的一个")
	}

	if pStr[1:4] != "::=" {
		exit("错误的产生式规则格式")
	}

	rightStr := pStr[4:]
	right := make([][]byte, 0)

	hasReadBar := true
	for i := 0; i < len(rightStr); i++ {
		if rightStr[i] == '|' {
			if hasReadBar {
				exit("错误的产生式规则格式")
			} else {
				hasReadBar = true
			}
			continue
		}

		if isCase(rightStr[i]) {
			expression := make([]byte, 0)
			for i < len(rightStr) && isCase(rightStr[i]) {
				expression = append(expression, rightStr[i])
				i++
			}
			i--

			hasReadBar = false
			if !g.checkVt(rightStr[i]) {
				g.Vt[rightStr[i]] = struct{}{}
			}

			right = append(right, expression)
			continue
		}

		exit("意外的输入")
	}

	g.P = append(g.P, Production{
		Left:        left,
		Right:       right,
		OriginalStr: cleanPStr,
	})
}

func (g *Grammar) checkVn(c byte) bool {
	if _, ok := g.Vn[c]; !ok {
		return false
	} else {
		return true
	}
}

func (g *Grammar) checkVt(c byte) bool {
	if _, ok := g.Vt[c]; !ok {
		return false
	} else {
		return true
	}
}

func (g *Grammar) Print() {
	fmt.Printf("文法%c[%c]=()", g.Name, g.S)
}
