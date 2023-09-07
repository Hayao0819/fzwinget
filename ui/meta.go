package ui

type MetaItem struct{
	Name string
	Func func()error
}

func (m MetaItem) Label() string {
	return m.Name
}
func (m MetaItem) Detail() string {
	return ""
}


