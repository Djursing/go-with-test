package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		exp := "this is just a test"

		assertStrings(t, got, exp)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dict.Add(word, definition)

		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dict := Dictionary{word: definition}
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		newDef := "new definition"
		dict := Dictionary{word: def}

		err := dict.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, def)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}

		dict.Delete(word)

		_, err := dict.Search(word)
		if err != ErrNotFound {
			t.Errorf("\nExpected %q to be deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, exp string) {
	t.Helper()

	if got != exp {
		t.Errorf("\ngot: %q\nexp: %q", got, exp)
	}
}

func assertError(t testing.TB, got, exp error) {
	t.Helper()

	if got != exp {
		t.Errorf("\ngot: %q\nexp: %q", got, exp)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("\ngot: %q\nexp: %q", got, definition)
	}
}
