// #21. Broken XML
// https://coderun.yandex.ru/problem/corrupted-xml?currentPage=3&pageSize=10&rowNumber=21
package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"unsafe"
)

var allowedLetters []byte

func init() {
	allowedLetters = make([]byte, 0, ('z'-'a')+1)
	for c := 'a'; c <= 'z'; c++ {
		allowedLetters = append(allowedLetters, byte(c))
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadBytes('\n')
	s = bytes.TrimSpace(s)
	fixPos, fixChar, err := parseFix(s)
	if err != nil {
		panic(err)
	}
	if fixPos < 0 {
		panic("fix expected")
	}
	s[fixPos] = fixChar
	fmt.Println(unsafeToString(s))
}

func parseFix(s []byte) (int, byte, error) {
	tags, tokenizeErr := tokenize(s)
	if tokenizeErr != nil {
		var possibleFixes []fix
		if err, ok := tokenizeErr.(*tokenizeError); ok {
			switch err.state {
			case tokenizeStateLeftChevron:
				possibleFixes = appendFixes(possibleFixes, err.pos, '<')
				if err.pos > 0 {
					possibleFixes = appendFixes(possibleFixes, err.pos-1, allowedLetters...)
				}
			case tokenizeStateSlashOrFirstLetter:
				possibleFixes = appendFixes(possibleFixes, err.pos, '/')
				possibleFixes = appendFixes(possibleFixes, err.pos, allowedLetters...)
			case tokenizeStateLetter:
				possibleFixes = appendFixes(possibleFixes, err.pos, allowedLetters...)
				possibleFixes = appendFixes(possibleFixes, err.pos-1, allowedLetters...)
			case tokenizeStateLetterOrRightChevron:
				possibleFixes = appendFixes(possibleFixes, err.pos, '>')
				possibleFixes = appendFixes(possibleFixes, err.pos-1, '>')
				possibleFixes = appendFixes(possibleFixes, err.pos, allowedLetters...)
			}
		}
		for _, pf := range possibleFixes {
			backup := s[pf.pos]
			s[pf.pos] = pf.char
			tags, err := tokenize(s)
			if err != nil {
				s[pf.pos] = backup
				continue
			}
			err = reduce(tags)
			s[pf.pos] = backup
			if err == nil {
				return pf.pos, pf.char, nil
			}
		}
		return -1, 0, tokenizeErr
	}
	reduceErr := reduce(tags)
	if reduceErr != nil {
		possibleFixes, err := analyzeReduceError(tags)
		if err != nil {
			return -1, 0, err
		}
		for _, pf := range possibleFixes {
			backup := s[pf.pos]
			s[pf.pos] = pf.char
			err = reduce(tags)
			s[pf.pos] = backup
			if err == nil {
				return pf.pos, pf.char, nil
			}
		}
		return -1, 0, reduceErr
	}
	return -1, 0, nil
}

func tokenize(s []byte) ([]tag, error) {
	state := tokenizeStateLeftChevron
	tags := make([]tag, 0, len(s)/3)
	begin := 0
	for i, char := range s {
		switch state {
		case tokenizeStateLeftChevron:
			if char != '<' {
				return tags, &tokenizeError{state, i, char}
			}
			begin = i
			state = tokenizeStateSlashOrFirstLetter
		case tokenizeStateSlashOrFirstLetter:
			if char == '<' || char == '>' {
				return tags, &tokenizeError{state, i, char}
			}
			if char == '/' {
				state = tokenizeStateLetter
				continue
			}
			state = tokenizeStateLetterOrRightChevron
		case tokenizeStateLetter:
			if char == '<' || char == '>' || char == '/' {
				return tags, &tokenizeError{state, i, char}
			}
			state = tokenizeStateLetterOrRightChevron
		case tokenizeStateLetterOrRightChevron:
			if char == '<' || char == '/' {
				return tags, &tokenizeError{state, i, char}
			}
			if char == '>' {
				t := tag{s[begin : i+1], begin}
				tags = append(tags, t)
				state = tokenizeStateLeftChevron
				continue
			}
		}
	}
	return tags, nil
}

type tag struct {
	content []byte
	pos     int
}

func (t *tag) isClosing() bool {
	return t.content[1] == '/'
}

func (t *tag) name() string {
	if t.isClosing() {
		nameBytes := t.content[2 : len(t.content)-1]
		return *(*string)(unsafe.Pointer(&nameBytes))
	}
	nameBytes := t.content[1 : len(t.content)-1]
	return unsafeToString(nameBytes)
}

func unsafeToString(p []byte) string {
	return *(*string)(unsafe.Pointer(&p))
}

type tokenizeState int

const (
	tokenizeStateLeftChevron tokenizeState = iota
	tokenizeStateSlashOrFirstLetter
	tokenizeStateLetter
	tokenizeStateLetterOrRightChevron
)

func (ts tokenizeState) want() string {
	return map[tokenizeState]string{
		tokenizeStateLeftChevron:          "<",
		tokenizeStateSlashOrFirstLetter:   "/|[a-z]",
		tokenizeStateLetter:               "[a-z]",
		tokenizeStateLetterOrRightChevron: "[a-z]|>",
	}[ts]
}

type tokenizeError struct {
	state tokenizeState
	pos   int
	char  byte
}

func (e *tokenizeError) Error() string {
	return fmt.Sprintf("tokenize error at %d: want=%s, got=%c", e.pos, e.state.want(), e.char)
}

func reduce(tags []tag) error {
	stack := make([]int, 0, len(tags)/2)
	for i, t := range tags {
		// Put opening tags into the stack.
		if !t.isClosing() {
			stack = append(stack, i)
			continue
		}
		// If there are no opening tags available, then return an error.
		if len(stack) == 0 {
			return &reduceError{-1, nil, i, &tags[i]}
		}
		// If the opening tag on the top of the stack does not match the
		// closing tag name, then return an error.
		stackHead := len(stack) - 1
		openingTag := &tags[stack[stackHead]]
		if openingTag.name() != t.name() {
			return &reduceError{stackHead, openingTag, i, &tags[i]}
		}
		// Remove the opening tag from the stack.
		stack = stack[:len(stack)-1]
	}
	if len(stack) > 0 {
		return errors.New("stack not empty")
	}
	return nil
}

type reduceError struct {
	openingTagIdx int
	openingTag    *tag
	closingTagIdx int
	closingTag    *tag
}

func (e *reduceError) Error() string {
	if e.openingTagIdx == -1 {
		return fmt.Sprintf("reduce error: closing=[%d]%s", e.closingTagIdx, e.closingTag.content)
	}
	return fmt.Sprintf("reduce error: opening=[%d]%s, closing=[%d]%s",
		e.openingTagIdx, e.openingTag.content, e.closingTagIdx, e.closingTag.content)
}

func generateXML(alphabet string, depth, maxSequenceLen, maxNameLen int) []byte {
	var buf bytes.Buffer
	sequenceLen := rand.Intn(maxSequenceLen) + 1
	for i := 0; i < sequenceLen; i++ {
		generateXMLTag(&buf, alphabet, depth, maxSequenceLen, maxNameLen)
	}
	return buf.Bytes()
}

func generateXMLTag(buf *bytes.Buffer, alphabet string, depth, maxSequenceLen, maxNameLen int) {
	// Write an opening tag with random name
	nameLen := rand.Intn(maxNameLen) + 1
	buf.WriteByte('<')
	nameBegin := buf.Len()
	for i := 0; i < nameLen; i++ {
		buf.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	nameEnd := buf.Len()
	name := buf.Bytes()[nameBegin:nameEnd]
	buf.WriteByte('>')

	// Unless the recursion bottom has been reached write random number of tags.
	if depth > 0 {
		sequenceLen := rand.Intn(maxSequenceLen) + 1
		for i := 0; i < sequenceLen; i++ {
			generateXMLTag(buf, alphabet, depth-1, maxSequenceLen, maxNameLen)
		}
	}

	// Write a closing tag
	buf.WriteByte('<')
	buf.WriteByte('/')
	buf.Write(name)
	buf.WriteByte('>')
}

func analyzeReduceError(tags []tag) ([]fix, error) {
	// Group tags with the same name. Collect opening and closing tags separately.
	openingTagsByName := make(map[string][]*tag, len(tags)/2)
	closingTagsByName := make(map[string][]*tag, len(tags)/2)
	tagNames := make(map[string]struct{}, len(tags)/2)
	for i := range tags {
		t := &tags[i]
		tagNames[t.name()] = struct{}{}
		tagsByName := openingTagsByName
		if t.isClosing() {
			tagsByName = closingTagsByName
		}
		nameGroup := tagsByName[t.name()]
		nameGroup = append(nameGroup, t)
		tagsByName[t.name()] = nameGroup
	}

	// Select tag names that have uneven number of opening and closing tags.
	mismatched := make([]mismatchTag, 0, 2)
	for tagName := range tagNames {
		openingTags := openingTagsByName[tagName]
		closingTags := closingTagsByName[tagName]
		diff := len(openingTags) - len(closingTags)
		if diff != 0 {
			mismatched = append(mismatched, mismatchTag{tagName, diff, openingTags, closingTags})
		}
	}

	if len(mismatched) != 2 {
		return nil, fmt.Errorf("tag mismatch cannot be fixed: %v", mismatched)
	}

	// Consider all possible fixes.
	possibleFixes := make([]fix, 0, 100)

	if len(mismatched[0].name) == len(mismatched[1].name) {
		pos := distinctChar(mismatched[0].name, mismatched[1].name)
		// If flip opening[1]->opening[0]
		if mismatched[0].diff == 1 && mismatched[1].diff == -1 {
			// Replace distinct char of opening[0] with distinct char of mismatched[1]
			for _, t := range mismatched[0].opening {
				possibleFixes = appendFixes(possibleFixes, t.pos+1+pos, mismatched[1].name[pos])
			}
			return possibleFixes, nil
		}
		// If flip opening[0]->opening[1]
		if mismatched[0].diff == -1 && mismatched[1].diff == 1 {
			// Replace distinct char of opening[1] with distinct char of mismatched[0]
			for _, t := range mismatched[1].opening {
				possibleFixes = appendFixes(possibleFixes, t.pos+1+pos, mismatched[0].name[pos])
			}
			return possibleFixes, nil
		}
		return nil, fmt.Errorf("tag mismatch cannot be fixed: %v", mismatched)
	}

	if len(mismatched[0].name)+1 == len(mismatched[1].name) {
		// If flip opening[1]->closing[0] ...
		if mismatched[0].diff == -1 && mismatched[1].diff == -1 {
			// Replace `/` of closing[0] to the first char of mismatched[0].name
			for _, t := range mismatched[0].closing {
				possibleFixes = appendFixes(possibleFixes, t.pos+1, mismatched[1].name[0])
			}
			return possibleFixes, nil
		}
		// If flip closing[0]->opening[1]
		if mismatched[0].diff == 1 && mismatched[1].diff == 1 {
			// Replace first char of opening[1] to '/'
			for _, t := range mismatched[1].opening {
				possibleFixes = appendFixes(possibleFixes, t.pos+1, '/')
			}
			return possibleFixes, nil
		}
		return nil, fmt.Errorf("tag mismatch cannot be fixed: %v", mismatched)
	}

	if len(mismatched[0].name) == len(mismatched[1].name)+1 {
		// If flip opening[0]->closing[1]
		if mismatched[0].diff == -1 && mismatched[1].diff == -1 {
			// Replace `/` of closing[1] to the first char of mismatched[0].name
			for _, t := range mismatched[1].closing {
				possibleFixes = appendFixes(possibleFixes, t.pos+1, mismatched[0].name[0])
			}
			return possibleFixes, nil
		}
		// If flip closing[1]->opening[0]
		if mismatched[0].diff == 1 && mismatched[1].diff == 1 {
			// Replace first char of opening[0] to '/'
			for _, t := range mismatched[0].opening {
				possibleFixes = appendFixes(possibleFixes, t.pos+1, '/')
			}
			return possibleFixes, nil
		}
		return nil, fmt.Errorf("tag mismatch cannot be fixed: %v", mismatched)
	}

	// Alas, there is nothing we can do to fix that damn thing :(
	return nil, fmt.Errorf("tag mismatch cannot be fixed: %v", mismatched)
}

type mismatchTag struct {
	name    string
	diff    int
	opening []*tag
	closing []*tag
}

type fix struct {
	pos  int
	char byte
}

func distinctChar(s1, s2 string) int {
	for i := range s1 {
		c1 := s1[i]
		c2 := s2[i]
		if c1 != c2 {
			return i
		}
	}
	return -1
}

func appendFixes(fixes []fix, pos int, chars ...byte) []fix {
	for _, c := range chars {
		fixes = append(fixes, fix{pos, c})
	}
	return fixes
}
