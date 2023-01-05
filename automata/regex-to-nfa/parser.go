package main

import (
	"fmt"
)

func matchCharacter(t *Token, c string) (string, error) {
	if t.Value == c {
		return t.Value, nil
	}
	return "", fmt.Errorf("character doesn't match")
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
	
	c, err := matchCharacter(t, "*")
	if err != nil {
		l.UnreadToken(t)
		return nil, nil
	}

	fmt.Printf(c)
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
			
			n, err := closure(l)
			if err != nil {
				return nil, fmt.Errorf("wrong punctutation %s", err)
			}

			t, err := l.ReadToken()
			if err != nil {
				return nil, err
			}

			c, err = matchCharacter(t, ")")
			fmt.Printf(c)

			return n, err
	}
	return nil, nil
}

