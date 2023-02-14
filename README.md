# iso3166

## Installation
```
go get github.com/DarkJian/iso3166
```

## Important
All ISO3166 country codes have to passed as UPPERCASE and returns UPPERCASE

## Usage

```golang
package main

import (
	"fmt"
	iso "github.com/DarkJian/iso3166"
)

func main() {
	var exist bool

  // if you have Alpha2 ISO3116 code

	exist = iso.I3166().ALPHA2().Exist("US")
	fmt.Println(exist)

	exist = iso.I3166().ALPHA2().Exist("us")
	fmt.Println(exist)

	alpha3Code, err := iso.I3166().ALPHA2().GetAlpha3("US")
	fmt.Println(alpha3Code, err)

	alpha3Code, err = iso.I3166().ALPHA2().GetAlpha3("us")
	fmt.Println(alpha3Code, err)

  // if you have Alpha3 ISO3116 code
  
	exist = iso.I3166().ALPHA3().Exist("USA")
	fmt.Println(exist)

	exist = iso.I3166().ALPHA3().Exist("usa")
	fmt.Println(exist)

	alpha2Code, err := iso.I3166().ALPHA3().GetAlpha2("USA")
	fmt.Println(alpha2Code, err)

	alpha2Code, err = iso.I3166().ALPHA3().GetAlpha2("usa")
	fmt.Println(alpha2Code, err)
}
```

Outputs:
```
true
false
USA, nil
"", error description

true
false
US, nil
"", error description
```
