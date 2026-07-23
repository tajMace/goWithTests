package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		searchCompare(t, "test", "this is a test")
	})

	t.Run("key not in map", func(t *testing.T) {
		searchEmpty(t, "badKey", ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		dict := Dictionary{}
		addDefinition(t, dict, "test", "this is a test")
	})

	t.Run("fail case: redefinition", func(t *testing.T) {
		dict := Dictionary{}
		addDefinition(t, dict, "test", "this is a test")
		err := dict.Add("test", "overwrite")
		assertError(t, err, ErrOverwrite)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		dict := Dictionary{}
		addDefinition(t, dict, "test", "this is a test")
		dict.Update("test", "this is STILL a test")

		result, err := dict.Search("test")
		if err != nil {
			t.Fatal("didn't expect an error, but got one anyway")
		}

		assertStrings(t, "test", result, "this is STILL a test")
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("newWord", "no definition yet")
		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		dict := Dictionary{}
		addDefinition(t, dict, "test", "this is a test")
		err := dict.Delete("test")
		if err != nil {
			t.Fatal("did not expect an error, but got one")
		}

		_, err = dict.Search("test")
		assertError(t, err, ErrNotFound)
	})

	t.Run("word doesn't exist", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("test")
		assertError(t, err, ErrNotFound)
	})
}

// ========== HELPERS ==========

func assertStrings(t testing.TB, given, got, want string) {
	if got != want {
		t.Errorf("given %q, got %q, but wanted %q", given, got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error, but didn't get one")
	}
	if err.Error() != want.Error() {
		t.Errorf("got error %q, want %q", err.Error(), want.Error())
	}
}

func searchCompare(t testing.TB, given, want string) {
	t.Helper()

	dictionary := Dictionary{given: want}
	got, _ := dictionary.Search(given)

	assertStrings(t, given, got, want)
}

func searchEmpty(t testing.TB, given, want string) {
	t.Helper()

	dictionary := Dictionary{"test": "this is a test"}
	_, err := dictionary.Search("bad")

	if err == nil {
		t.Fatal("given a bad key, expected an error, but didn't get one.")
	}

	assertStrings(t, given, err.Error(), want)
}

func addDefinition(t testing.TB, dictionary Dictionary, given, want string) {
	dictionary.Add(given, want)
	got, err := dictionary.Search(given)

	if err != nil {
		t.Fatal("didn't expect error:", err)
	}

	assertStrings(t, given, got, want)
}
