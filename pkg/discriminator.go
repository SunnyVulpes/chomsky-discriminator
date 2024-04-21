package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grammar struct {
	Name  byte
	Vn    map[byte]struct{}
	Vt    map[byte]struct{}
	P     []Production
	S     byte
	Class int
}

type Production struct {
	Left        []byte
	Right       [][]byte
	OriginalStr string
}

func BuildGrammar() *Grammar {
	g := &Grammar{}
	g.Vn = make(map[byte]struct{})
	g.Vt = make(map[byte]struct{})
	g.P = make([]Production, 0)
	g.Class = 3

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

	at := strings.Index(cleanPStr, "::=")
	if at < 1 {
		exit("产生式未含有::=")
	}
	if len(cleanPStr) < at+4 {
		exit("产生式未含有右边")
	}

	left := []byte(pStr[:at])
	hasVn := false
	for _, c := range left {
		if !g.checkVn(c) {
			g.Vt[c] = struct{}{}
		} else {
			hasVn = true
		}
	}
	if !hasVn {
		exit("产生式左边不含有非终结符号，非法输入")
	}

	rightStr := pStr[at+3:]
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
				if !g.checkVn(rightStr[i]) {
					g.Vt[rightStr[i]] = struct{}{}
				}
				expression = append(expression, rightStr[i])
				i++
			}
			i--

			hasReadBar = false

			right = append(right, expression)
			continue
		}

		exit("意外的输入")
	}

	class := g.classifyP(left, right)
	if class < g.Class {
		g.Class = class
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
	var builder strings.Builder
	builder.WriteString("{")
	flag := false
	for k, _ := range g.Vn {
		if flag {
			builder.WriteString(",")
		}
		flag = true
		builder.WriteString(fmt.Sprintf("%c", k))
	}
	builder.WriteString("}")

	var builder1 strings.Builder
	flag = false
	builder1.WriteString("{")
	for k, _ := range g.Vt {
		if flag {
			builder1.WriteString(",")
		}
		flag = true
		builder1.WriteString(fmt.Sprintf("%c", k))
	}
	builder1.WriteString("}")

	fmt.Printf("文法%c[%c]=(%s, %s, P, %c)", g.Name, g.S, builder.String(), builder1.String(), g.S)
	fmt.Printf("\nP:\n")
	for _, p := range g.P {
		fmt.Printf("%s\n", p.OriginalStr)
	}
	fmt.Printf("该文法是 Chomsky%d型文法", g.Class)
}

func (g *Grammar) classifyP(left []byte, right [][]byte) int {
	flag := false
	for _, c := range left {
		if g.checkVt(c) {
			flag = true
		}
	}

	if flag {
		for _, r := range right {
			if len(r) == 1 && r[0] == 'e' {
				return 0
			}
			if len(left) > len(r) {
				return 0
			}
		}
		return 1
	} else {
		for _, r := range right {
			n := true
			for _, s := range r {
				if !n && !g.checkVn(s) {
					return 2
				}
				if n && !g.checkVt(s) {
					n = false
					continue
				}
				if !n && g.checkVt(s) {
					return 2
				}
			}
		}
		return 3
	}
}
