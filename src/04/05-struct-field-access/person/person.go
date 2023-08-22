package person

// Person represents the Person model.
type Person struct {
	FirstName string // Exportable
	LastName  string // Exportable
	secret    string // nolint Unexportable (private)
}
