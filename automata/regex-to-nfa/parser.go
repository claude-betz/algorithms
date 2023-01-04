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

func union(l *Lexer) (*node, error) {	
	n1, err1 := concat(l)
	n2, err2 := unionTail(l)

	// construct union NFA
	return n1, nil
}

func unionTail(l *Lexer) (*node, error) {
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}

	_, err = matchPunctuation(t, "|", EOF)	
	if err != nil {
		return nil, err
	}

	n, err := concat(l) 

	return unionTail(l)
}

func concat(l *Lexer) (*node, error) {
	closure(l)
	concatTail(l)

	// construct concat NFA
}

func concatTail(l *Lexer) (*node, error) {
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}

	// match tag for cocat
	if t.Tag == TagId {
		closure(l)
	}

	concatTail(l)
}

func closure(l *Lexer) (*node, error) {
	value(l)
	closureTail(l)

	// construct closure NFA
}

func closureTail(l *Lexer) (*node, error) {
	t, err := l.ReadToken()
	if err != nil {
		return nil, err
	}
	fmt.Printf("token: %s", t.Value)
	
	_, err = matchPunctuation(t, "*", EOF)
	if err != nil {
		return nil, err
	}

	return closureTail(l)
}

func value(l *Lexer) (*node, error) {
	t, err := l.ReadToken()
	if err != nil {
		// error
	}
	fmt.Printf("token: %s", t.Value)
	
	switch t.Tag {
		case TagId:
			// construct value NFA
			return &Value{t.Value}, nil 
		case TagPunct:	
			matchPunctuation(t, "(")
			if err != nil {
				return nil, fmt.Errorf("wrong token in idlst: %s", err)
			}

			n, err := union(l)

			matchPunctuation(t, ")")
			if err != nil {
				return nil, fmt.Errorf("wrong token in idlst: %s", err)
			}

			// construct value NFA
			return n, err
	}
}

