package person

// Person represents the Person model.
type Person struct {
	FirstName string // Exportable
	LastName  string // Exportable
	secret    string // Unexportable (private)
}

// Secret returns private secret field.
func (u Person) Secret() string {
	return u.secret
}

// SetSecret sets private secret value.
func (u *Person) SetSecret(s string) {
	u.secret = s
}
