package main

import (
	"fmt"
	"strings"
)

func peek(l *Lexer) (*Token, error) {
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}

	l.UnreadToken(t)
	return t, nil
}

func matchCharacter(t *Token, opts ...string) (string, error) {
	s := strings.Join(opts, "")
	if t.Tag == TagPunct && strings.Index(s, t.Value) >= 0 {
		return t.Value, nil
	}
	return "", fmt.Errorf("unknown token %s", t)
}

func matchTag(t *Token, tag Tag) bool {
	if t.Tag == tag {
		return true
	}
	return false
}

func union(l *Lexer) (node, error) {
	n, err := concat(l)
	if err != nil {
		return nil, err
	}

	n, err = unionTail(l)
	if err != nil {
		return nil, err
	}

	return n, err
}

func unionTail(l *Lexer) (node, error) {
	t, err := peek(l)
	if err != nil {
		return nil, err
	}
	
	c, err := matchCharacter(t, "|", EOF)
	if c == EOF {
		return nil, fmt.Errorf("end of file")
	}
	if err != nil {	
		return nil, nil
	}

	t, err = l.ReadToken()
	fmt.Printf(t.Value)

	n, err := concat(l)
	if err != nil {
		return nil, err
	}
	
	n, err = unionTail(l)
	if err != nil {
		return nil, err
	}

	return n, err
}

func concat(l *Lexer) (node, error) {
	n, err := closure(l)
	if err != nil {
		return nil, err
	}

	n, err = concatTail(l)
	if err != nil {
		return nil, err
	}

	return n, err
}

func concatTail(l *Lexer) (node, error) {
	t, err := peek(l)
	if err != nil {
		return nil, err
	}
	
	if !matchTag(t, TagId) {
		return nil, nil				
	}

	n, err := closure(l)
	if err != nil {
		return nil, err
	}

	n, err = concatTail(l)
	if err != nil {
		return nil, err
	}

	return n, err
}

func closure(l *Lexer) (node, error) {
	n, err := value(l)
	if err != nil {
		return nil, err
	}

	n, err = closureTail(l)
	if err != nil {
		return nil, err
	}

	return n, err
}

func closureTail(l *Lexer) (node, error) {
	t, err := peek(l)
	if err != nil {
		return nil, err
	}
	
	c, err := matchCharacter(t, "*", EOF)
	if c == EOF {
		return nil, fmt.Errorf("end of file")
	}
	if err != nil {	
		return nil, nil
	}

	t, err = l.ReadToken()
	if err != nil {
		return nil, err
	}
	fmt.Printf(t.Value)
	return nil, nil
}

func value(l *Lexer) (node, error) {
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}
	
	switch t.Tag {
		case TagId:
			// construct value NFA
			fmt.Printf(t.Value)
			return &ValueNode{}, nil 
		case TagPunct:		
			c, err := matchCharacter(t, "(")
			if err != nil {
				return nil, err
			}	
			fmt.Printf(c)
			
			n, err := union(l)
			if err != nil {
				return nil, fmt.Errorf("wrong punctutation %s", err)
			}

			t, err := l.ReadToken()
			if err != nil {
				return nil, err
			}

			c, err = matchCharacter(t, ")")
			if err != nil {
				return nil, err
			}
			fmt.Printf(c)

			return n, err
	}
	return nil, nil
}

