// #18. Evaluate arithmatic expression.
// https://coderun.yandex.ru/problem/the-value-of-an-arithmetic-expression?currentPage=2&pageSize=10&rowNumber=18
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	exp := scanner.Text()
	result, err := Eval(exp)
	if err != nil {
		fmt.Println("WRONG")
		return
	}
	fmt.Println(result)
}

func Eval(exp string) (int, error) {
	tnz := newTokenizer(exp)
	stk := new(stack)
	tok, err := tnz.read()
	if err != nil {
		if err == io.EOF {
			return 0, nil
		}
		return 0, err
	}
	var val int
	switch tok.name {
	case tokenNum:
		stk.push(tok.val)
	case tokenLPar:
		if val, err = evalPar(tnz, stk); err != nil {
			return 0, err
		}
		stk.push(val)
	case tokenPlus:
		if val, err = evalSign(tnz, stk, "+"); err != nil {
			return 0, err
		}
		stk.push(val)
	case tokenMinus:
		if val, err = evalSign(tnz, stk, "-"); err != nil {
			return 0, err
		}
		stk.push(val)
	default:
		return 0, fmt.Errorf("invalid token '%s' at %d", tok, tok.pos)
	}
	for {
		tok, err = tnz.read()
		if err != nil {
			if err == io.EOF {
				goto done
			}
		}
		switch tok.name {
		case tokenLPar:
			if val, err = evalPar(tnz, stk); err != nil {
				return 0, err
			}
			prevVal, ok := stk.pop()
			if !ok {
				panic("must never reach this point")
			}
			val = prevVal * val
		case tokenMult:
			if val, err = evalMult(tnz, stk); err != nil {
				return 0, err
			}
		case tokenPlus:
			if val, err = evalAddSub(tnz, stk, "+"); err != nil {
				return 0, err
			}
		case tokenMinus:
			if val, err = evalAddSub(tnz, stk, "-"); err != nil {
				return 0, err
			}
		default:
			return 0, fmt.Errorf("invalid token '%s' at %d", tok, tok.pos)
		}
		stk.push(val)
	}
done:
	if stk.len() != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	val, _ = stk.pop()
	return val, nil
}

func evalPar(tnz *tokenizer, stk *stack) (int, error) {
	tok, err := tnz.read()
	if err != nil {
		return 0, fmt.Errorf("evalPar: empty")
	}
	var val int
	switch tok.name {
	case tokenNum:
		stk.push(tok.val)
	case tokenLPar:
		if val, err = evalPar(tnz, stk); err != nil {
			return 0, err
		}
		stk.push(val)
	case tokenPlus:
		if val, err = evalSign(tnz, stk, "+"); err != nil {
			return 0, err
		}
		stk.push(val)
	case tokenMinus:
		if val, err = evalSign(tnz, stk, "-"); err != nil {
			return 0, err
		}
		stk.push(val)
	default:
		return 0, fmt.Errorf("evalPar: invalid token '%s' at %d", tok, tok.pos)
	}
	for {
		tok, err = tnz.read()
		if err != nil {
			if err == io.EOF {
				goto done
			}
		}
		switch tok.name {
		case tokenLPar:
			if val, err = evalPar(tnz, stk); err != nil {
				return 0, err
			}
			prevVal, ok := stk.pop()
			if !ok {
				panic("must never reach this point")
			}
			val = prevVal * val
		case tokenRPar:
			tok, err = tnz.read()
			if err != nil {
				if err == io.EOF {
					goto done
				}
				return 0, err
			}
			if tok.name != tokenLPar {
				tnz.backtrack(tok)
				goto done
			}
			if val, err = evalPar(tnz, stk); err != nil {
				return 0, err
			}
			prevVal, ok := stk.pop()
			if !ok {
				panic("must never reach this point")
			}
			val = prevVal * val
		case tokenMult:
			if val, err = evalMult(tnz, stk); err != nil {
				return 0, err
			}
		case tokenPlus:
			if val, err = evalAddSub(tnz, stk, "+"); err != nil {
				return 0, err
			}
		case tokenMinus:
			if val, err = evalAddSub(tnz, stk, "-"); err != nil {
				return 0, err
			}
		default:
			return 0, fmt.Errorf("evalPar: invalid token '%s' at %d", tok, tok.pos)
		}
		stk.push(val)
	}
done:
	val, ok := stk.pop()
	if !ok {
		panic("must never reach this point")
	}
	return val, nil
}

func evalSign(tnz *tokenizer, stk *stack, sign string) (int, error) {
	tok, err := tnz.read()
	if err != nil {
		return 0, fmt.Errorf("evalSign: operand missing at %d", tnz.pos())
	}
	var val int
	switch tok.name {
	case tokenNum:
		val = tok.val
	case tokenLPar:
		if val, err = evalPar(tnz, stk); err != nil {
			return 0, err
		}
	case tokenPlus:
		if val, err = evalSign(tnz, stk, "+"); err != nil {
			return 0, err
		}
	case tokenMinus:
		if val, err = evalSign(tnz, stk, "-"); err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("evalSign: invalid token '%s' at %d", tok, tok.pos)
	}

	if sign == "-" {
		return -val, nil
	}
	return val, nil
}

func evalMult(tnz *tokenizer, stk *stack) (int, error) {
	leftOp, ok := stk.pop()
	if !ok {
		return 0, fmt.Errorf("evalMult: left operand missing at %d", tnz.pos())
	}
	tok, err := tnz.read()
	if err != nil {
		return 0, fmt.Errorf("evalMult: right operand missing at %d", tnz.pos())
	}
	var rightOp int
	switch tok.name {
	case tokenNum:
		rightOp = tok.val
	case tokenLPar:
		if rightOp, err = evalPar(tnz, stk); err != nil {
			return 0, err
		}
	case tokenPlus:
		if rightOp, err = evalSign(tnz, stk, "+"); err != nil {
			return 0, err
		}
	case tokenMinus:
		if rightOp, err = evalSign(tnz, stk, "-"); err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("evalMult: invalid token '%s' at %d", tok, tok.pos)
	}

	return leftOp * rightOp, nil
}

func evalAddSub(tnz *tokenizer, stk *stack, sign string) (int, error) {
	leftOp, ok := stk.pop()
	if !ok {
		return 0, fmt.Errorf("evalAddSub: left operand missing at %d", tnz.pos())
	}
	tok, err := tnz.read()
	if err != nil {
		return 0, fmt.Errorf("evalAddSub: right operand missing at %d", tnz.pos())
	}
	var val int
	switch tok.name {
	case tokenNum:
		stk.push(tok.val)
	case tokenLPar:
		if val, err = evalPar(tnz, stk); err != nil {
			return 0, err
		}
		stk.push(val)
	case tokenPlus:
		if val, err = evalSign(tnz, stk, "+"); err != nil {
			return 0, err
		}
		stk.push(val)
	case tokenMinus:
		if val, err = evalSign(tnz, stk, "-"); err != nil {
			return 0, err
		}
		stk.push(val)
	default:
		return 0, fmt.Errorf("evalAddSub: invalid token '%s' at %d", tok, tok.pos)
	}
	for {
		tok, err = tnz.read()
		if err != nil {
			if err == io.EOF {
				goto done
			}
		}
		switch tok.name {
		case tokenLPar:
			if val, err = evalPar(tnz, stk); err != nil {
				return 0, err
			}
			prevVal, ok := stk.pop()
			if !ok {
				panic("must never reach this point")
			}
			val = prevVal * val
		case tokenRPar:
			tnz.backtrack(tok)
			goto done
		case tokenMult:
			if val, err = evalMult(tnz, stk); err != nil {
				return 0, err
			}
		case tokenPlus:
			tnz.backtrack(tok)
			goto done
		case tokenMinus:
			tnz.backtrack(tok)
			goto done
		default:
			return 0, fmt.Errorf("evalAddSub: invalid token '%s' at %d", tok, tok.pos)
		}
		stk.push(val)
	}
done:
	rightOp, ok := stk.pop()
	if !ok {
		return 0, fmt.Errorf("evalAddSub: right operand missing at %d", tnz.pos())
	}

	if sign == "-" {
		return leftOp - rightOp, nil
	}
	return leftOp + rightOp, nil
}

func newTokenizer(in string) *tokenizer {
	return &tokenizer{in: in}
}

type tokenizer struct {
	in          string
	readPos     int
	prevTok     token
	backtracked []token
}

func (t *tokenizer) read() (token, error) {
	if len(t.backtracked) > 0 {
		lastIdx := len(t.backtracked) - 1
		tok := t.backtracked[lastIdx]
		t.backtracked = t.backtracked[:lastIdx]
		return tok, nil
	}
	for t.readPos < len(t.in) {
		ch := t.in[t.readPos]
		switch {
		case ch == '(':
			t.prevTok = token{name: tokenLPar, pos: t.readPos}
			t.readPos++
			return t.prevTok, nil
		case ch == ')':
			t.prevTok = token{name: tokenRPar, pos: t.readPos}
			t.readPos++
			return t.prevTok, nil
		case ch == '*':
			t.prevTok = token{name: tokenMult, pos: t.readPos}
			t.readPos++
			return t.prevTok, nil
		case ch == '+':
			t.prevTok = token{name: tokenPlus, pos: t.readPos}
			t.readPos++
			return t.prevTok, nil
		case ch == '-':
			t.prevTok = token{name: tokenMinus, pos: t.readPos}
			t.readPos++
			return t.prevTok, nil
		case ch == ' ':
			t.readPos++
			continue
		case ch >= '0' || ch <= '9':
			return t.readNum()
		default:
			return token{}, fmt.Errorf("invalid symbol '%c' at %d", ch, t.readPos)
		}
	}
	return token{}, io.EOF
}

func (t *tokenizer) readPlusOrSignedNum() (token, error) {
	tok := token{name: tokenPlus, pos: t.readPos}
	t.readPos++
	if t.readPos >= len(t.in) {
		return tok, nil
	}
	ch := t.in[t.readPos]
	if ch < '0' || ch > '9' {
		return tok, nil
	}
	if t.prevTok.name == tokenRPar {
		return tok, nil
	}
	if t.prevTok.name == tokenNum {
		return tok, nil
	}
	tok, err := t.readNum()
	if err != nil {
		return tok, err
	}
	tok.pos--
	t.prevTok = tok
	return tok, nil
}

func (t *tokenizer) readMinusOrSignedNum() (token, error) {
	tok := token{name: tokenMinus, pos: t.readPos}
	t.readPos++
	if t.readPos >= len(t.in) {
		return tok, nil
	}
	ch := t.in[t.readPos]
	if ch < '0' || ch > '9' {
		return tok, nil
	}
	if t.prevTok.name == tokenRPar {
		return tok, nil
	}
	if t.prevTok.name == tokenNum {
		return tok, nil
	}
	tok, err := t.readNum()
	if err != nil {
		return tok, err
	}
	tok.pos--
	tok.val = -tok.val
	t.prevTok = tok
	return tok, err
}

func (t *tokenizer) readNum() (token, error) {
	tok := token{name: tokenNum, pos: t.readPos}
	for t.readPos++; t.readPos < len(t.in); t.readPos++ {
		ch := t.in[t.readPos]
		if ch < '0' || ch > '9' {
			break
		}
	}
	var err error
	if tok.val, err = strconv.Atoi(t.in[tok.pos:t.readPos]); err != nil {
		return token{}, fmt.Errorf("invalid number at %d: %w", tok.pos, err)
	}
	t.prevTok = tok
	return tok, nil
}

func (t *tokenizer) backtrack(tok token) {
	t.backtracked = append(t.backtracked, tok)
}

func (t *tokenizer) pos() int {
	return t.prevTok.pos
}

type tokenName string

const (
	tokenLPar  tokenName = "("
	tokenRPar  tokenName = ")"
	tokenMult  tokenName = "*"
	tokenPlus  tokenName = "+"
	tokenMinus tokenName = "-"
	tokenNum   tokenName = "{num}"
)

type token struct {
	name tokenName
	pos  int
	val  int
}

func (t token) String() string {
	if t.name == tokenNum {
		return strconv.Itoa(t.val)
	}
	return string(t.name)
}

type stack struct {
	buf []int
}

func (s *stack) push(val int) {
	s.buf = append(s.buf, val)
}

func (s *stack) pop() (int, bool) {
	if len(s.buf) == 0 {
		return 0, false
	}
	lastIdx := len(s.buf) - 1
	val := s.buf[lastIdx]
	s.buf = s.buf[:lastIdx]
	return val, true
}

func (s *stack) len() int {
	return len(s.buf)
}
