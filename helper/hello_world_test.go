package helper_test

import (
	"fmt"
	helper "learning-golang-unit-test/helper"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
commands:
go test
-v => verbose
-run=TestFunc/SubTest => run test that match with the given name
-bench=BenchmarkFunc/SubBenchmark => run benchmark that match with the given name
./... => to test all tests/benchmarks deep inside

notes:
panic() 										=> stop the app
Fail(), Error(), assert 		=> still run the func if failed
FailNow(), Fatal(), require => stop the func if failed

*/

// benchmark also has features such as sub benchmark and benchmark table
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.HelloWorld()
	}
}

func TestTable(t *testing.T) {
	tests := []struct {
		Name     string
		Expected string
	}{
		{
			Name:     "Dita",
			Expected: "Hello Dita",
		}, {
			Name:     "Laras",
			Expected: "Hello Laras",
		}, {
			Name:     "Adi",
			Expected: "Hello Adi",
		}, {
			Name:     "Tri",
			Expected: "Hello Tri",
		},
	}

	for _, value := range tests {
		result := helper.HelloWorld()

		t.Run(value.Name, func(t *testing.T) {
			require.Equal(t, value.Expected, result, "Test table "+value.Name+" is failed")
		})
	}

}

func TestSubtest(t *testing.T) {
	result := helper.HelloWorld()
	t.Run("Subtest1", func(t *testing.T) {
		if result != "Hello World" {
			t.Error("Subtest1 is failed")
		}

		fmt.Println("Subtest1 done")
	})

	t.Run("Subtest2", func(t *testing.T) {
		if result != "Hello World" {
			t.Error("Subtest2 is failed")
		}

		fmt.Println("Subtest2 done")
	})

}

// hook before after the unit test of this package be ran
func TestMain(m *testing.M) {
	fmt.Println("BEFORE TEST RUN")

	m.Run()

	fmt.Println("AFTER TEST RUN")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("This unit test is skipped for GOOS darwin")
	}
}

//////////////////////////////////////////////////

// use panic
/*
func TestHelloWorld(t *testing.T) {
	result := helper.HelloWorld()

	if result != "Hello World" {
		panic("TestHelloWorld is failed")
		// remember, panic stop the application if there is no recovery
	}

	// not printed if failed
	fmt.Println("TestHelloWorld() done")
}
*/

// use Fail()
func TestHelloWorldFail(t *testing.T) {
	result := helper.HelloWorld()

	if result != "Hello World" {
		t.Fail()
	}

	// printed if failed
	fmt.Println("TestHelloWorldFail() done")

}

// use FailNow()
func TestHelloWorldFailNow(t *testing.T) {
	result := helper.HelloWorld()

	if result != "Hello World" {
		t.FailNow()
	}

	// not printed if failed
	fmt.Println("TestHelloWorldFailNow() done")

}

// use Error()
func TestHelloWorldError(t *testing.T) {
	result := helper.HelloWorld()

	if result != "Hello World" {
		t.Error("TestHelloWorldError is failed")
	}

	// printed if failed
	fmt.Println("TestHelloWorldError() done")

}

// use Fatal()
func TestHelloWorldFatal(t *testing.T) {
	result := helper.HelloWorld()

	if result != "Hello World" {
		t.Fatal("TestHelloWorldFatal() is failed")
	}

	// not printed if failed
	fmt.Println("TestHelloWorldFatal() done")

}

// use testify/assert
func TestHelloWorldAssert(t *testing.T) {
	result := helper.HelloWorld()

	assert.Equal(t, "Hello World", result, "TestHelloWorldAssert() is failed")

	// printed if failed
	fmt.Println("TestHelloWorldAssert() done")

}

// use testify/require
func TestHelloWorldRequire(t *testing.T) {
	result := helper.HelloWorld()

	require.Equal(t, "Hello World", result, "TestHelloWorldRequire() is failed")

	// not printed if failed
	fmt.Println("TestHelloWorldRequire() done")

}
