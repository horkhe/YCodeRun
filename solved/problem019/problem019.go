// #19. Evaluate logical expression
// https://coderun.yandex.ru/problem/the-value-of-a-boolean-expression?currentPage=2&pageSize=10&rowNumber=19
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	exp := scanner.Text()
	res, err := Eval(exp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Eval(exp string) (int, error) {
	stack := new(stack)
	tokenizer := newTokenizer(exp)
	return evalExp(stack, tokenizer)
}

func evalExp(stack *stack, tokenizer *tokenizer) (int, error) {
	token, tokenPos, err := tokenizer.read()
	if err != nil {
		return -1, fmt.Errorf("invalid expression: %w", err)
	}

	var val int
	switch token {
	case "0":
		val = 0
	case "1":
		val = 1
	case "(":
		if val, err = evalPar(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	case "!":
		if val, err = evalNot(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	default:
		return -1, fmt.Errorf("invalid token '%s' at %d", token, tokenPos)
	}
	stack.push(val)

	for {
		token, tokenPos, err := tokenizer.read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return -1, fmt.Errorf("invalid expression: %w", err)
		}
		switch token {
		case "&":
			if val, err = evalAnd(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case "|":
			if val, err = evalOr(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case "^":
			if val, err = evalXor(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		default:
			return -1, fmt.Errorf("invalid token '%s' at %d", token, tokenPos)
		}
		stack.push(val)
	}

	if stack.len() != 1 {
		return -1, errors.New("invalid expression")
	}
	res, _ := stack.pop()
	return res, nil
}

func evalPar(stack *stack, tokenizer *tokenizer, pos int) (int, error) {
	token, tokenPos, err := tokenizer.read()
	if err != nil {
		return -1, fmt.Errorf("evalPar: invalid expression at %d: %w", pos, err)
	}
	var val int
	switch token {
	case "0":
		val = 0
	case "1":
		val = 1
	case "(":
		if val, err = evalPar(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	case "!":
		if val, err = evalNot(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	default:
		return -1, fmt.Errorf("evalPar: invalid token '%s' at %d", token, tokenPos)
	}
	stack.push(val)

	for {
		token, tokenPos, err := tokenizer.read()
		if err != nil {
			return -1, fmt.Errorf("evalPar: invalid expression at %d: %w", pos, err)
		}
		switch token {
		case "&":
			if val, err = evalAnd(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case "|":
			if val, err = evalOr(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case "^":
			if val, err = evalXor(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case ")":
			goto done
		default:
			return -1, fmt.Errorf("evalPar: invalid token '%s' at %d", token, tokenPos)
		}
		stack.push(val)
	}
done:
	res, ok := stack.pop()
	if !ok {
		return -1, fmt.Errorf("evalPar: invalid expression at %d", pos)
	}
	return res, nil
}

func evalNot(stack *stack, tokenizer *tokenizer, pos int) (int, error) {
	token, tokenPos, err := tokenizer.read()
	if err != nil {
		return -1, fmt.Errorf("evalNot: operand missing at %d: %w", pos, err)
	}
	var op int
	switch token {
	case "0":
		op = 0
	case "1":
		op = 1
	case "!":
		if op, err = evalNot(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	case "(":
		if op, err = evalPar(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	default:
		return -1, fmt.Errorf("evalNot: invalid token '%s' at %d", token, tokenPos)
	}
	res := op ^ 1
	return res, nil
}

func evalAnd(stack *stack, tokenizer *tokenizer, pos int) (int, error) {
	leftOp, ok := stack.pop()
	if !ok {
		return -1, fmt.Errorf("evalAnd: left operand missing at %d", pos)
	}
	token, tokenPos, err := tokenizer.read()
	if err != nil {
		return -1, fmt.Errorf("evalAnd: right operand missing at %d: %w", pos, err)
	}
	var rightOp int
	switch token {
	case "0":
		rightOp = 0
	case "1":
		rightOp = 1
	case "!":
		if rightOp, err = evalNot(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	case "(":
		if rightOp, err = evalPar(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	default:
		return -1, fmt.Errorf("evalAnd: invalid token '%s' at %d", token, tokenPos)
	}
	res := leftOp & rightOp
	return res, nil
}

func evalOr(stack *stack, tokenizer *tokenizer, pos int) (int, error) {
	leftOp, ok := stack.pop()
	if !ok {
		return -1, fmt.Errorf("evalOr: left operand missing at %d", pos)
	}

	token, tokenPos, err := tokenizer.read()
	if err != nil {
		return -1, fmt.Errorf("evalOr: right operand missing at %d: %w", pos, err)
	}
	var val int
	switch token {
	case "0":
		val = 0
	case "1":
		val = 1
	case "(":
		if val, err = evalPar(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	case "!":
		if val, err = evalNot(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	default:
		return -1, fmt.Errorf("evalOr: invalid right operand '%s' at %d", token, tokenPos)
	}
	stack.push(val)

	for {
		token, tokenPos, err := tokenizer.read()
		if err != nil {
			if err == io.EOF {
				goto done
			}
			return -1, fmt.Errorf("evalOr: invalid right operand at %d: %w", pos, err)
		}
		switch token {
		case "&":
			if val, err = evalAnd(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case "|":
			tokenizer.backtrack()
			goto done
		case "^":
			tokenizer.backtrack()
			goto done
		case ")":
			tokenizer.backtrack()
			goto done
		default:
			return -1, fmt.Errorf("evalOr: invalid right operand '%s' at %d", token, tokenPos)
		}
		stack.push(val)
	}
done:
	rightOp, ok := stack.pop()
	if !ok {
		return -1, fmt.Errorf("evalOr: invalid right operand at %d", pos)
	}

	res := leftOp | rightOp
	return res, nil
}

func evalXor(stack *stack, tokenizer *tokenizer, pos int) (int, error) {
	leftOp, ok := stack.pop()
	if !ok {
		return -1, fmt.Errorf("evalXor: left operand missing at %d", pos)
	}

	token, tokenPos, err := tokenizer.read()
	if err != nil {
		return -1, fmt.Errorf("evalXor: right operand missing at %d: %w", pos, err)
	}
	var val int
	switch token {
	case "0":
		val = 0
	case "1":
		val = 1
	case "(":
		if val, err = evalPar(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	case "!":
		if val, err = evalNot(stack, tokenizer, tokenPos); err != nil {
			return -1, err
		}
	default:
		return -1, fmt.Errorf("evalXor: invalid right operand '%s' at %d", token, tokenPos)
	}
	stack.push(val)

	for {
		token, tokenPos, err := tokenizer.read()
		if err != nil {
			if err == io.EOF {
				goto done
			}
			return -1, fmt.Errorf("evalXor: invalid right operand at %d: %w", pos, err)
		}
		switch token {
		case "&":
			if val, err = evalAnd(stack, tokenizer, tokenPos); err != nil {
				return -1, err
			}
		case "|":
			tokenizer.backtrack()
			goto done
		case "^":
			tokenizer.backtrack()
			goto done
		case ")":
			tokenizer.backtrack()
			goto done
		default:
			return -1, fmt.Errorf("evalXor: invalid right operand '%s' at %d", token, tokenPos)
		}
		stack.push(val)
	}
done:
	rightOp, ok := stack.pop()
	if !ok {
		return -1, fmt.Errorf("evalXor: invalid right operand at %d", pos)
	}

	res := leftOp ^ rightOp
	return res, nil
}

func newTokenizer(exp string) *tokenizer {
	return &tokenizer{exp: exp}
}

type tokenizer struct {
	exp string
	pos int
}

func (t *tokenizer) read() (string, int, error) {
	if t.pos >= len(t.exp) {
		return "", -1, io.EOF
	}
	tokenPos := t.pos
	t.pos = t.pos + 1
	token := t.exp[tokenPos:t.pos]
	return token, tokenPos, nil
}

func (t *tokenizer) backtrack() {
	if t.pos > 0 {
		t.pos--
	}
}

type stack struct {
	buf []int
}

func (s *stack) push(val int) {
	s.buf = append(s.buf, val)
}

func (s *stack) pop() (int, bool) {
	if len(s.buf) == 0 {
		return -1, false
	}
	lastIdx := len(s.buf) - 1
	val := s.buf[lastIdx]
	s.buf = s.buf[:lastIdx]
	return val, true
}

func (s *stack) len() int {
	return len(s.buf)
}
