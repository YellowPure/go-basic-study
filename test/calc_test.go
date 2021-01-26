package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1+2 expected be 3, but %d got", ans)
	}
}

type calcCase struct {
	A, B, Expected int
	Name           string
}

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("test1")
}
func Test2(t *testing.T) {
	fmt.Println("test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
		}
	})
}

// func TestMul(t *testing.T) {
// 	createMulTestCase(t, &calcCase{2, 3, 6, "pos"})
// 	createMulTestCase(t, &calcCase{2, 0, 1, "zero"})
// if ans := Mul(-10, -20); ans != 200 {
// 	t.Errorf("-10*-20 expected be 200, but %d got", ans)
// }
// cases := []struct {
// 	Name           string
// 	A, B, Expected int
// }{
// 	{"pos", 2, 3, 6},
// 	{"neg", 2, -3, -6},
// 	{"zero", 2, 0, 0},
// }

// for _, c := range cases {
// 	t.Run(c.Name, func(t *testing.T) {
// 		if ans := Mul(c.A, c.B); ans != c.Expected {
// 			t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
// 		}
// 	})
// }
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)

	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
	fmt.Println("test conn")
	// ln, err := net.Listen("tcp", "127.0.0.1:0")
	// handleError(t, err)
	// defer ln.Close()

	// http.HandleFunc("/hello", helloHandler)
	// go http.Serve(ln, nil)

	// resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	// handleError(t, err)
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// handleError(t, err)

	// if string(body) != "hello world" {
	// 	t.Fatal("expected hello world, but got", string(body))
	// }
}

// func BenchmarkHello(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Sprintf("hello")
// 	}
// }

func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Helloo, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "world")
		}
	})
}
