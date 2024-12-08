// Package todostatus represents the status of a todo in the system.
package todoapp

import "fmt"

// The set of status types that can be used.
var (
	Complete   = newStatus("COMPLETE")
	Incomplete = newStatus("INCOMPLETE")
)

// =============================================================================

// Set of known status types.
var statusTypes = make(map[string]Status)

// Status represents a status type in the system.
type Status struct {
	value string
}

func newStatus(status string) Status {
	st := Status{status}
	statusTypes[status] = st
	return st
}

// String returns the name of the status.
func (st Status) String() string {
	return st.value
}

// Equal provides support for comparisons and testing.
func (st Status) Equal(st2 Status) bool {
	return st.value == st2.value
}

// MarshalText provides support for logging and any marshal needs.
func (st Status) MarshalText() ([]byte, error) {
	return []byte(st.value), nil
}

// =============================================================================

// Parse parses the string value and returns a status type if one exists.
func Parse(value string) (Status, error) {
	typ, exists := statusTypes[value]
	if !exists {
		return Status{}, fmt.Errorf("invalid status type %q", value)
	}
	return typ, nil
}

// MustParse parses the string value and returns a status type if one exists. If
// an error occurs, the function panics.
func MustParse(value string) Status {
	typ, err := Parse(value)
	if err != nil {
		panic(err)
	}
	return typ
}
