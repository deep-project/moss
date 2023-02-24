package entity

type Theme struct {
	Current string `json:"current"`
}

func (*Theme) ConfigID() string {
	return "theme"
}

func NewTheme() *Theme {
	return &Theme{Current: "germ"}
}
