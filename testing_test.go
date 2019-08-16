package testing

import (
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	test := NewTestSuite(t)
	test.Before(func() {
		fmt.Println("setting up the test case")
	})

	test.After(func() {
		fmt.Println("clean up the test case")
	})

	test.Test("step 1", func() {
		test.Context().Log("step 1 is in progress")
	}).Test("step 2", func() {
		test.Context().Error("step 2 is failed")
	}).Execute()

}
