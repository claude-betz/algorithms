package main

import (
	"fmt"
	"strings"
)

func matchPunctuation(t *Token, options ...string) (string, error) {
	s := strings.Join(options, "")
	if t.Tag == TagPunct && strings.Index(s, t.Value) >= 0 {
		return t.Value, nil
	}
	return "", fmt.Errorf("unknown token: %s", t)
}

func matchTokenTag(t *Token, tag Tag) (bool, error) {
	if t.Tag == tag {
		return true, nil 
	}
	return false, fmt.Errorf("unknown tag: %s", tag)
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
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}
	fmt.Printf(t.Value)
	
	_, err = matchPunctuation(t, "*")
	if err != nil {
		return nil, err
	}

	return closureTail(l)
}

func value(l *Lexer) (node, error) {
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}
	fmt.Printf(t.Value)
	
	switch t.Tag {
		case TagId:
			// construct value NFA
			return &ValueNode{}, nil 
		case TagPunct:	
			_, err := matchPunctuation(t, "(")
			if err != nil {
				return nil, fmt.Errorf("wrong punctuation %s", err)
			}

			n, err := closure(l)
			if err != nil {
				return nil, fmt.Errorf("wrong punctutation %s", err)
			}

			_, err = matchPunctuation(t, ")")
			if err != nil {
				return nil, fmt.Errorf("wrong punctuation %s", err)
			}

			// construct value NFA
			return n, err
	}
	return nil, nil
}

