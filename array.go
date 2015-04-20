package pqarrays

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
)

var (
	ErrUnexpectedValueType = errors.New("expected value to be a string or []byte")
)

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	output := make([]string, 0, len(s))
	for _, item := range s {
		item = strconv.Quote(item)
		item = strings.Replace(item, "'", "\\'", -1)
		output = append(output, item)
	}
	return []byte(`{` + strings.Join(output, ",") + `}`), nil
}

func (s *StringArray) Scan(value interface{}) error {
	*s = (*s)[:0]
	var input string
	if _, ok := value.(string); ok {
		input = value.(string)
	} else if _, ok := value.([]byte); ok {
		input = string(value.([]byte))
	} else {
		return ErrUnexpectedValueType
	}
	l := lex(input)
	parsed, err := parse(l)
	if err != nil {
		return err
	}
	for _, item := range parsed {
		if item == nil {
			continue
		}
		*s = append(*s, *item)
	}
	return nil
}
