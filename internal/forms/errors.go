package forms

// errors is a map of errors, keyed by field name
type errors map[string][]string

// Add adds an error message for a given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first message for the given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
