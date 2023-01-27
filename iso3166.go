package iso316

type ISO interface {
	ALPHA2() Alpha2CI
	ALPHA3() Alpha3CI
}

type Alpha2CI interface {
	Exist(code string) bool
	GetAlpha3(code string) (string, error)
}

type Alpha3CI interface {
	Exist(code string) bool
	GetAlpha2(code string) (string, error)
}

// data used in alpha2 and alpha3 structs
type data struct {
	A2      string `json:"alpha2"`
	A3      string `json:"alpha3"`
	Numeric string `json:"numeric"`
	Name    string `json:"name"`
}

// I3166 entry point
func I3166() ISO {
	return i
}

// iso main struct used in package init function
type iso struct {
	Alpha2 *alpha2
	Alpha3 *alpha3
}

// ALPHA2 returns ALPHA2 controller interface, use if you have Alpha2 iso code
func (i *iso) ALPHA2() Alpha2CI {
	return i.Alpha2
}

// ALPHA3 returns ALPHA3 controller interface, use if you have Alpha3 iso code
func (i *iso) ALPHA3() Alpha3CI {
	return i.Alpha3
}
