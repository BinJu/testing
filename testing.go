package testing

import (
	"fmt"
	"testing"
)

type TestSuite interface {
	Execute()
	Before(func()) TestSuite
	After(func()) TestSuite
	Test(name string, test func()) TestSuite
	Context() *testing.T
}

type testItem struct {
	Name     string
	TestFunc func()
}

type testSuite struct {
	before func()
	after  func()
	test   []testItem
	t      *testing.T
}

func NewTestSuite(t *testing.T) TestSuite {
	return &testSuite{t: t}
}

func (t *testSuite) Context() *testing.T {
	return t.t
}

func (t *testSuite) Before(before func()) TestSuite {
	t.before = before
	return t
}

func (t *testSuite) After(after func()) TestSuite {
	t.after = after
	return t
}

func (t *testSuite) Test(name string, test func()) TestSuite {
	t.test = append(t.test, testItem{Name: name, TestFunc: test})
	return t
}

func (t *testSuite) Execute() {
	if nil != t.before {
		t.before()
	}
	if nil != t.after {
		t.after()
	}
	for _, f := range t.test {
		fmt.Printf("=== Testing: %s\n", f.Name)
		f.TestFunc()
	}
}
