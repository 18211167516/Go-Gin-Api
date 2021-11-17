package response

type Button struct {
	Url         string
	Permissions bool
	Allow       func() string
}

func (b Button) Ok() func() string {
	return func() string {
		if b.Permissions {
			return ""
		}
		return "hide"
	}
}
