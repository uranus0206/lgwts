package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	// expected := "Chris"
	// var got []string

	// x := struct {
	// 	Name string
	// }{expected}

	// walk(x, func(input string) {
	// 	got = append(got, input)
	// })

	// if len(got) != 1 {
	// 	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	// }

	// if got[0] != expected {
	// 	t.Errorf("got '%s', want '%s'", got[0], expected)
	// }

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reysjshdj"},
			},
			[]string{"London", "Reysjshdj"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "R"},
			},
			[]string{"London", "R"},
		},
		// {
		// 	"Maps",
		// 	map[string]string{
		// 		"Foo": "Bar",
		// 		"Baz": "Boz",
		// 	},
		// 	[]string{"Bar", "Boz"},
		// },
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}

		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContain(t, got, "Bar")
		assertContain(t, got, "Boz")

	})
}

func assertContain(t *testing.T, got []string, need string) {
	contains := false

	for _, x := range got {
		if x == need {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v contain %s, but it didn't.", got, need)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
