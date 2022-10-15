package testingExample

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Fatal("The world is crumbling")
	}
}

// func TestGreeting(t *testing.T) {
// 	got := Greeting("George")
// 	want := "Hello George"
// 	if got != want {
// 		t.Fatalf("Expected %q, got %q", want, got)
// 	}
// }

func TestFrTranslation(t *testing.T) {
	got := translate("fr-FR")
	want := "Bonjour"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestUSTranslation(t *testing.T) {
	got := Greeting("George", "en-US")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

type GreetingTest struct {
	name   string
	locale string
	want   string
}

var greetingTest = []GreetingTest{
	{"George", "en-US", "Hello George"},
	{"Chloé", "fr-FR", "Bonjour Chloé"},
	{"Giuseppe", "it-IT", "Ciao Giuseppe"},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTest {
		got := Greeting(test.name, test.locale)
		if got != test.want {
			t.Errorf("Greeting(%s,%s) = %v; want %v", test.name, test.locale, got, test.want)
		}
	}
}

func BenchmarkStringFromAssignement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAssignement(100)
	}
}

func BenchmarkFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(100)
	}
}

func BenchmarkFromStringBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(100)
	}
}
