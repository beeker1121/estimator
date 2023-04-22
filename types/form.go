package types

// Form defines a form.
type Form struct {
	ID      string   `json:"id"`
	Modules []Module `json:"modules"`
}
