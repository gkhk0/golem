package golem

import (
	"errors"
	"fmt"
	"slices"
)

type tokenizer struct {
  tokens []rune
}

func (t *tokenizer) next() (rune, error) {
  if len(t.tokens) < 1 {
    return 0, errors.New("tokens is empty");
  }
  token := t.tokens[0];
  t.tokens = slices.Delete(t.tokens, 0, 1)
  return token, nil
}

func Print(a string) {
  fmt.Println(a);
}
