package main

// Product is app-defined model
type Product struct {
	Model
	ID          int    `json:"id"`
	Flags       int    `json:"flags"`
	Name        string `json:"name" f0x:"req lenmax:50 lenmin:3"`
	Description string `json:"description" f0x:"lenmax:255"`
}
