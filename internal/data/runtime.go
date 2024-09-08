package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("Invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// Unmarshal method that satisfies the json.Unmashaler interface
// Use a pointer because UnmarshalJSON() needs to modify the receiver or we will be
// modifying a copy which gets discarded when the method returns
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// Convert int32 to runtime type and assign it to receiver
	// Use * to dereference the receiver in order to set underlying value of the pointer
	*r = Runtime(i)

	return nil
}
