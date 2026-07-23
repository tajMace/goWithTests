package maps

type Dictionary map[string]string

const (
	ErrNotFound  = DictionaryErr("word not in dictionary")
	ErrOverwrite = DictionaryErr("word already in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrOverwrite
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	delete(d, word)
	return nil
}
