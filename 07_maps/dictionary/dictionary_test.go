package dictionary

import (
	"testing"
)

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error")
	}

	if got != want {
		t.Errorf("got err %s want %s", got, want)
	}
}

func assertDefinition(t *testing.T, dic Dictionary, word, want string) {
	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:  ", err)
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		word := "test"
		want := "this is just a test"
		assertDefinition(t, dictionary, word, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete exist word", func(t *testing.T) {
		word := "test"
		definition := "test definition"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %s to be deleted", word)
		}
	})

	t.Run("delete noneixst word", func(t *testing.T) {
		word := "test"
		definition := "test definition"
		dictionary := Dictionary{word: definition}
		word2 := "test2"
		err := dictionary.Delete(word2)
		assertError(t, err, ErrNotFound)
	})
}
