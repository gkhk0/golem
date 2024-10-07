package golem;

import "testing";

func PrintTest(t *testing.T) {
  Print("test");
  t.Errorf("error");
}
