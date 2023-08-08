package dictionary

type Dictionary map[string]string

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

const (
	ErrNotFound         = ErrDictionary("could not find the word you were looking for") //nolint: lll
	ErrWordExist        = ErrDictionary("cannot add word because it already exists")    //nolint: lll
	ErrWordDoesNotExist = ErrDictionary("cannot update word because it does not exist") //nolint: lll
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
