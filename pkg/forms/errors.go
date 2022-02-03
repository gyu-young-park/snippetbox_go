package forms

type errors map[string][]string

func (e errors) Add(filed, message string) {
	e[filed] = append(e[filed], message)
}

func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
