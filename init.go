package iso316

import "sync"

var i *iso

func init() {
	i = &iso{
		Alpha2: &alpha2{
			mu: &sync.RWMutex{},
			data: map[string]*data{
				"AF": {
					A2:      "AF",
					A3:      "AFG",
					Numeric: "004",
					Name:    "Afghanistan",
				}, "AL": {
					A2:      "AL",
					A3:      "ALB",
					Numeric: "008",
					Name:    "Albania, People's Socialist Republic of",
				}, "DZ": {
					A2:      "DZ",
					A3:      "DZA",
					Numeric: "012",
					Name:    "Algeria, People's Democratic Republic of",
				}, "AS": {
					A2:      "AS",
					A3:      "ASM",
					Numeric: "016",
					Name:    "American Samoa",
				}, "AD": {
					A2:      "AD",
					A3:      "AND",
					Numeric: "020",
					Name:    "Andorra, Principality of",
				}, "AO": {
					A2:      "AO",
					A3:      "AGO",
					Numeric: "024",
					Name:    "Angola, Republic of",
				}, "AI": {
					A2:      "AI",
					A3:      "AIA",
					Numeric: "660",
					Name:    "Anguilla",
				}, "AQ": {
					A2:      "AQ",
					A3:      "ATA",
					Numeric: "010",
					Name:    "Antarctica (the territory South of 60 deg S)",
				}, "AG": {
					A2:      "AG",
					A3:      "ATG",
					Numeric: "028",
					Name:    "Antigua and Barbuda",
				}, "AR": {
					A2:      "AR",
					A3:      "ARG",
					Numeric: "032",
					Name:    "Argentina, Argentine Republic",
				}, "AM": {
					A2:      "AM",
					A3:      "ARM",
					Numeric: "051",
					Name:    "Armenia",
				}, "AW": {
					A2:      "AW",
					A3:      "ABW",
					Numeric: "533",
					Name:    "Aruba",
				}, "AU": {
					A2:      "AU",
					A3:      "AUS",
					Numeric: "036",
					Name:    "Australia, Commonwealth of",
				}, "AT": {
					A2:      "AT",
					A3:      "AUT",
					Numeric: "040",
					Name:    "Austria, Republic of",
				}, "AZ": {
					A2:      "AZ",
					A3:      "AZE",
					Numeric: "031",
					Name:    "Azerbaijan, Republic of",
				}, "BS": {
					A2:      "BS",
					A3:      "BHS",
					Numeric: "044",
					Name:    "Bahamas, Commonwealth of the",
				}, "BH": {
					A2:      "BH",
					A3:      "BHR",
					Numeric: "048",
					Name:    "Bahrain, Kingdom of",
				}, "BD": {
					A2:      "BD",
					A3:      "BGD",
					Numeric: "050",
					Name:    "Bangladesh, People's Republic of",
				}, "BB": {
					A2:      "BB",
					A3:      "BRB",
					Numeric: "052",
					Name:    "Barbados",
				}, "BY": {
					A2:      "BY",
					A3:      "BLR",
					Numeric: "112",
					Name:    "Belarus",
				}, "BE": {
					A2:      "BE",
					A3:      "BEL",
					Numeric: "056",
					Name:    "Belgium, Kingdom of",
				}, "BZ": {
					A2:      "BZ",
					A3:      "BLZ",
					Numeric: "084",
					Name:    "Belize",
				}, "BJ": {
					A2:      "BJ",
					A3:      "BEN",
					Numeric: "204",
					Name:    "Benin, People's Republic of",
				}, "BM": {
					A2:      "BM",
					A3:      "BMU",
					Numeric: "060",
					Name:    "Bermuda",
				}, "BT": {
					A2:      "BT",
					A3:      "BTN",
					Numeric: "064",
					Name:    "Bhutan, Kingdom of",
				}, "BO": {
					A2:      "BO",
					A3:      "BOL",
					Numeric: "068",
					Name:    "Bolivia, Republic of",
				}, "BA": {
					A2:      "BA",
					A3:      "BIH",
					Numeric: "070",
					Name:    "Bosnia and Herzegovina",
				}, "BW": {
					A2:      "BW",
					A3:      "BWA",
					Numeric: "072",
					Name:    "Botswana, Republic of",
				}, "BV": {
					A2:      "BV",
					A3:      "BVT",
					Numeric: "074",
					Name:    "Bouvet Island (Bouvetoya)",
				}, "BR": {
					A2:      "BR",
					A3:      "BRA",
					Numeric: "076",
					Name:    "Brazil, Federative Republic of",
				}, "IO": {
					A2:      "IO",
					A3:      "IOT",
					Numeric: "086",
					Name:    "British Indian Ocean Territory (Chagos Archipelago)",
				}, "VG": {
					A2:      "VG",
					A3:      "VGB",
					Numeric: "092",
					Name:    "British Virgin Islands",
				}, "BN": {
					A2:      "BN",
					A3:      "BRN",
					Numeric: "096",
					Name:    "Brunei Darussalam",
				}, "BG": {
					A2:      "BG",
					A3:      "BGR",
					Numeric: "100",
					Name:    "Bulgaria, People's Republic of",
				}, "BF": {
					A2:      "BF",
					A3:      "BFA",
					Numeric: "854",
					Name:    "Burkina Faso",
				}, "BI": {
					A2:      "BI",
					A3:      "BDI",
					Numeric: "108",
					Name:    "Burundi, Republic of",
				}, "KH": {
					A2:      "KH",
					A3:      "KHM",
					Numeric: "116",
					Name:    "Cambodia, Kingdom of",
				}, "CM": {
					A2:      "CM",
					A3:      "CMR",
					Numeric: "120",
					Name:    "Cameroon, United Republic of",
				}, "CA": {
					A2:      "CA",
					A3:      "CAN",
					Numeric: "124",
					Name:    "Canada",
				}, "CV": {
					A2:      "CV",
					A3:      "CPV",
					Numeric: "132",
					Name:    "Cape Verde, Republic of",
				}, "KY": {
					A2:      "KY",
					A3:      "CYM",
					Numeric: "136",
					Name:    "Cayman Islands",
				}, "CF": {
					A2:      "CF",
					A3:      "CAF",
					Numeric: "140",
					Name:    "Central African Republic",
				}, "TD": {
					A2:      "TD",
					A3:      "TCD",
					Numeric: "148",
					Name:    "Chad, Republic of",
				}, "CL": {
					A2:      "CL",
					A3:      "CHL",
					Numeric: "152",
					Name:    "Chile, Republic of",
				}, "CN": {
					A2:      "CN",
					A3:      "CHN",
					Numeric: "156",
					Name:    "China, People's Republic of",
				}, "CX": {
					A2:      "CX",
					A3:      "CXR",
					Numeric: "162",
					Name:    "Christmas Island",
				}, "CC": {
					A2:      "CC",
					A3:      "CCK",
					Numeric: "166",
					Name:    "Cocos (Keeling) Islands",
				}, "CO": {
					A2:      "CO",
					A3:      "COL",
					Numeric: "170",
					Name:    "Colombia, Republic of",
				}, "KM": {
					A2:      "KM",
					A3:      "COM",
					Numeric: "174",
					Name:    "Comoros, Union of the",
				}, "CD": {
					A2:      "CD",
					A3:      "COD",
					Numeric: "180",
					Name:    "Congo, Democratic Republic of",
				}, "CG": {
					A2:      "CG",
					A3:      "COG",
					Numeric: "178",
					Name:    "Congo, People's Republic of",
				}, "CK": {
					A2:      "CK",
					A3:      "COK",
					Numeric: "184",
					Name:    "Cook Islands",
				}, "CR": {
					A2:      "CR",
					A3:      "CRI",
					Numeric: "188",
					Name:    "Costa Rica, Republic of",
				}, "CI": {
					A2:      "CI",
					A3:      "CIV",
					Numeric: "384",
					Name:    "Cote D'Ivoire, Ivory Coast, Republic of the",
				}, "CU": {
					A2:      "CU",
					A3:      "CUB",
					Numeric: "192",
					Name:    "Cuba, Republic of",
				}, "CY": {
					A2:      "CY",
					A3:      "CYP",
					Numeric: "196",
					Name:    "Cyprus, Republic of",
				}, "CZ": {
					A2:      "CZ",
					A3:      "CZE",
					Numeric: "203",
					Name:    "Czech Republic",
				}, "DK": {
					A2:      "DK",
					A3:      "DNK",
					Numeric: "208",
					Name:    "Denmark, Kingdom of",
				}, "DJ": {
					A2:      "DJ",
					A3:      "DJI",
					Numeric: "262",
					Name:    "Djibouti, Republic of",
				}, "DM": {
					A2:      "DM",
					A3:      "DMA",
					Numeric: "212",
					Name:    "Dominica, Commonwealth of",
				}, "DO": {
					A2:      "DO",
					A3:      "DOM",
					Numeric: "214",
					Name:    "Dominican Republic",
				}, "EC": {
					A2:      "EC",
					A3:      "ECU",
					Numeric: "218",
					Name:    "Ecuador, Republic of",
				}, "EG": {
					A2:      "EG",
					A3:      "EGY",
					Numeric: "818",
					Name:    "Egypt, Arab Republic of",
				}, "SV": {
					A2:      "SV",
					A3:      "SLV",
					Numeric: "222",
					Name:    "El Salvador, Republic of",
				}, "GQ": {
					A2:      "GQ",
					A3:      "GNQ",
					Numeric: "226",
					Name:    "Equatorial Guinea, Republic of",
				}, "ER": {
					A2:      "ER",
					A3:      "ERI",
					Numeric: "232",
					Name:    "Eritrea",
				}, "EE": {
					A2:      "EE",
					A3:      "EST",
					Numeric: "233",
					Name:    "Estonia",
				}, "ET": {
					A2:      "ET",
					A3:      "ETH",
					Numeric: "231",
					Name:    "Ethiopia",
				}, "FO": {
					A2:      "FO",
					A3:      "FRO",
					Numeric: "234",
					Name:    "Faeroe Islands",
				}, "FK": {
					A2:      "FK",
					A3:      "FLK",
					Numeric: "238",
					Name:    "Falkland Islands (Malvinas)",
				}, "FJ": {
					A2:      "FJ",
					A3:      "FJI",
					Numeric: "242",
					Name:    "Fiji, Republic of the Fiji Islands",
				}, "FI": {
					A2:      "FI",
					A3:      "FIN",
					Numeric: "246",
					Name:    "Finland, Republic of",
				}, "FR": {
					A2:      "FR",
					A3:      "FRA",
					Numeric: "250",
					Name:    "France, French Republic",
				}, "GF": {
					A2:      "GF",
					A3:      "GUF",
					Numeric: "254",
					Name:    "French Guiana",
				}, "PF": {
					A2:      "PF",
					A3:      "PYF",
					Numeric: "258",
					Name:    "French Polynesia",
				}, "TF": {
					A2:      "TF",
					A3:      "ATF",
					Numeric: "260",
					Name:    "French Southern Territories",
				}, "GA": {
					A2:      "GA",
					A3:      "GAB",
					Numeric: "266",
					Name:    "Gabon, Gabonese Republic",
				}, "GM": {
					A2:      "GM",
					A3:      "GMB",
					Numeric: "270",
					Name:    "Gambia, Republic of the",
				}, "GE": {
					A2:      "GE",
					A3:      "GEO",
					Numeric: "268",
					Name:    "Georgia",
				}, "DE": {
					A2:      "DE",
					A3:      "DEU",
					Numeric: "276",
					Name:    "Germany",
				}, "GH": {
					A2:      "GH",
					A3:      "GHA",
					Numeric: "288",
					Name:    "Ghana, Republic of",
				}, "GI": {
					A2:      "GI",
					A3:      "GIB",
					Numeric: "292",
					Name:    "Gibraltar",
				}, "GR": {
					A2:      "GR",
					A3:      "GRC",
					Numeric: "300",
					Name:    "Greece, Hellenic Republic",
				}, "GL": {
					A2:      "GL",
					A3:      "GRL",
					Numeric: "304",
					Name:    "Greenland",
				}, "GD": {
					A2:      "GD",
					A3:      "GRD",
					Numeric: "308",
					Name:    "Grenada",
				}, "GP": {
					A2:      "GP",
					A3:      "GLP",
					Numeric: "312",
					Name:    "Guadaloupe",
				}, "GU": {
					A2:      "GU",
					A3:      "GUM",
					Numeric: "316",
					Name:    "Guam",
				}, "GT": {
					A2:      "GT",
					A3:      "GTM",
					Numeric: "320",
					Name:    "Guatemala, Republic of",
				}, "GN": {
					A2:      "GN",
					A3:      "GIN",
					Numeric: "324",
					Name:    "Guinea, Revolutionary People's Rep'c of",
				}, "GW": {
					A2:      "GW",
					A3:      "GNB",
					Numeric: "624",
					Name:    "Guinea-Bissau, Republic of",
				}, "GY": {
					A2:      "GY",
					A3:      "GUY",
					Numeric: "328",
					Name:    "Guyana, Republic of",
				}, "HT": {
					A2:      "HT",
					A3:      "HTI",
					Numeric: "332",
					Name:    "Haiti, Republic of",
				}, "HM": {
					A2:      "HM",
					A3:      "HMD",
					Numeric: "334",
					Name:    "Heard and McDonald Islands",
				}, "VA": {
					A2:      "VA",
					A3:      "VAT",
					Numeric: "336",
					Name:    "Holy See (Vatican City State)",
				}, "HN": {
					A2:      "HN",
					A3:      "HND",
					Numeric: "340",
					Name:    "Honduras, Republic of",
				}, "HK": {
					A2:      "HK",
					A3:      "HKG",
					Numeric: "344",
					Name:    "Hong Kong, Special Administrative Region of China",
				}, "HR": {
					A2:      "HR",
					A3:      "HRV",
					Numeric: "191",
					Name:    "Hrvatska (Croatia)",
				}, "HU": {
					A2:      "HU",
					A3:      "HUN",
					Numeric: "348",
					Name:    "Hungary, Hungarian People's Republic",
				}, "IS": {
					A2:      "IS",
					A3:      "ISL",
					Numeric: "352",
					Name:    "Iceland, Republic of",
				}, "IN": {
					A2:      "IN",
					A3:      "IND",
					Numeric: "356",
					Name:    "India, Republic of",
				}, "ID": {
					A2:      "ID",
					A3:      "IDN",
					Numeric: "360",
					Name:    "Indonesia, Republic of",
				}, "IR": {
					A2:      "IR",
					A3:      "IRN",
					Numeric: "364",
					Name:    "Iran, Islamic Republic of",
				}, "IQ": {
					A2:      "IQ",
					A3:      "IRQ",
					Numeric: "368",
					Name:    "Iraq, Republic of",
				}, "IE": {
					A2:      "IE",
					A3:      "IRL",
					Numeric: "372",
					Name:    "Ireland",
				}, "IL": {
					A2:      "IL",
					A3:      "ISR",
					Numeric: "376",
					Name:    "Israel, State of",
				}, "IT": {
					A2:      "IT",
					A3:      "ITA",
					Numeric: "380",
					Name:    "Italy, Italian Republic",
				}, "JM": {
					A2:      "JM",
					A3:      "JAM",
					Numeric: "388",
					Name:    "Jamaica",
				}, "JP": {
					A2:      "JP",
					A3:      "JPN",
					Numeric: "392",
					Name:    "Japan",
				}, "JO": {
					A2:      "JO",
					A3:      "JOR",
					Numeric: "400",
					Name:    "Jordan, Hashemite Kingdom of",
				}, "KZ": {
					A2:      "KZ",
					A3:      "KAZ",
					Numeric: "398",
					Name:    "Kazakhstan, Republic of",
				}, "KE": {
					A2:      "KE",
					A3:      "KEN",
					Numeric: "404",
					Name:    "Kenya, Republic of",
				}, "KI": {
					A2:      "KI",
					A3:      "KIR",
					Numeric: "296",
					Name:    "Kiribati, Republic of",
				}, "KP": {
					A2:      "KP",
					A3:      "PRK",
					Numeric: "408",
					Name:    "Korea, Democratic People's Republic of",
				}, "KR": {
					A2:      "KR",
					A3:      "KOR",
					Numeric: "410",
					Name:    "Korea, Republic of",
				}, "KW": {
					A2:      "KW",
					A3:      "KWT",
					Numeric: "414",
					Name:    "Kuwait, State of",
				}, "KG": {
					A2:      "KG",
					A3:      "KGZ",
					Numeric: "417",
					Name:    "Kyrgyz Republic",
				}, "LA": {
					A2:      "LA",
					A3:      "LAO",
					Numeric: "418",
					Name:    "Lao People's Democratic Republic",
				}, "LV": {
					A2:      "LV",
					A3:      "LVA",
					Numeric: "428",
					Name:    "Latvia",
				}, "LB": {
					A2:      "LB",
					A3:      "LBN",
					Numeric: "422",
					Name:    "Lebanon, Lebanese Republic",
				}, "LS": {
					A2:      "LS",
					A3:      "LSO",
					Numeric: "426",
					Name:    "Lesotho, Kingdom of",
				}, "LR": {
					A2:      "LR",
					A3:      "LBR",
					Numeric: "430",
					Name:    "Liberia, Republic of",
				}, "LY": {
					A2:      "LY",
					A3:      "LBY",
					Numeric: "434",
					Name:    "Libyan Arab Jamahiriya",
				}, "LI": {
					A2:      "LI",
					A3:      "LIE",
					Numeric: "438",
					Name:    "Liechtenstein, Principality of",
				}, "LT": {
					A2:      "LT",
					A3:      "LTU",
					Numeric: "440",
					Name:    "Lithuania",
				}, "LU": {
					A2:      "LU",
					A3:      "LUX",
					Numeric: "442",
					Name:    "Luxembourg, Grand Duchy of",
				}, "MO": {
					A2:      "MO",
					A3:      "MAC",
					Numeric: "446",
					Name:    "Macao, Special Administrative Region of China",
				}, "MK": {
					A2:      "MK",
					A3:      "MKD",
					Numeric: "807",
					Name:    "Macedonia, the former Yugoslav Republic of",
				}, "MG": {
					A2:      "MG",
					A3:      "MDG",
					Numeric: "450",
					Name:    "Madagascar, Republic of",
				}, "MW": {
					A2:      "MW",
					A3:      "MWI",
					Numeric: "454",
					Name:    "Malawi, Republic of",
				}, "MY": {
					A2:      "MY",
					A3:      "MYS",
					Numeric: "458",
					Name:    "Malaysia",
				}, "MV": {
					A2:      "MV",
					A3:      "MDV",
					Numeric: "462",
					Name:    "Maldives, Republic of",
				}, "ML": {
					A2:      "ML",
					A3:      "MLI",
					Numeric: "466",
					Name:    "Mali, Republic of",
				}, "MT": {
					A2:      "MT",
					A3:      "MLT",
					Numeric: "470",
					Name:    "Malta, Republic of",
				}, "MH": {
					A2:      "MH",
					A3:      "MHL",
					Numeric: "584",
					Name:    "Marshall Islands",
				}, "MQ": {
					A2:      "MQ",
					A3:      "MTQ",
					Numeric: "474",
					Name:    "Martinique",
				}, "MR": {
					A2:      "MR",
					A3:      "MRT",
					Numeric: "478",
					Name:    "Mauritania, Islamic Republic of",
				}, "MU": {
					A2:      "MU",
					A3:      "MUS",
					Numeric: "480",
					Name:    "Mauritius",
				}, "YT": {
					A2:      "YT",
					A3:      "MYT",
					Numeric: "175",
					Name:    "Mayotte",
				}, "MX": {
					A2:      "MX",
					A3:      "MEX",
					Numeric: "484",
					Name:    "Mexico, United Mexican States",
				}, "FM": {
					A2:      "FM",
					A3:      "FSM",
					Numeric: "583",
					Name:    "Micronesia, Federated States of",
				}, "MD": {
					A2:      "MD",
					A3:      "MDA",
					Numeric: "498",
					Name:    "Moldova, Republic of",
				}, "MC": {
					A2:      "MC",
					A3:      "MCO",
					Numeric: "492",
					Name:    "Monaco, Principality of",
				}, "MN": {
					A2:      "MN",
					A3:      "MNG",
					Numeric: "496",
					Name:    "Mongolia, Mongolian People's Republic",
				}, "MS": {
					A2:      "MS",
					A3:      "MSR",
					Numeric: "500",
					Name:    "Montserrat",
				}, "MA": {
					A2:      "MA",
					A3:      "MAR",
					Numeric: "504",
					Name:    "Morocco, Kingdom of",
				}, "MZ": {
					A2:      "MZ",
					A3:      "MOZ",
					Numeric: "508",
					Name:    "Mozambique, People's Republic of",
				}, "MM": {
					A2:      "MM",
					A3:      "MMR",
					Numeric: "104",
					Name:    "Myanmar",
				}, "NA": {
					A2:      "NA",
					A3:      "NAM",
					Numeric: "516",
					Name:    "Namibia",
				}, "NR": {
					A2:      "NR",
					A3:      "NRU",
					Numeric: "520",
					Name:    "Nauru, Republic of",
				}, "NP": {
					A2:      "NP",
					A3:      "NPL",
					Numeric: "524",
					Name:    "Nepal, Kingdom of",
				}, "AN": {
					A2:      "AN",
					A3:      "ANT",
					Numeric: "530",
					Name:    "Netherlands Antilles",
				}, "NL": {
					A2:      "NL",
					A3:      "NLD",
					Numeric: "528",
					Name:    "Netherlands, Kingdom of the",
				}, "NC": {
					A2:      "NC",
					A3:      "NCL",
					Numeric: "540",
					Name:    "New Caledonia",
				}, "NZ": {
					A2:      "NZ",
					A3:      "NZL",
					Numeric: "554",
					Name:    "New Zealand",
				}, "NI": {
					A2:      "NI",
					A3:      "NIC",
					Numeric: "558",
					Name:    "Nicaragua, Republic of",
				}, "NE": {
					A2:      "NE",
					A3:      "NER",
					Numeric: "562",
					Name:    "Niger, Republic of the",
				}, "NG": {
					A2:      "NG",
					A3:      "NGA",
					Numeric: "566",
					Name:    "Nigeria, Federal Republic of",
				}, "NU": {
					A2:      "NU",
					A3:      "NIU",
					Numeric: "570",
					Name:    "Niue, Republic of",
				}, "NF": {
					A2:      "NF",
					A3:      "NFK",
					Numeric: "574",
					Name:    "Norfolk Island",
				}, "MP": {
					A2:      "MP",
					A3:      "MNP",
					Numeric: "580",
					Name:    "Northern Mariana Islands",
				}, "NO": {
					A2:      "NO",
					A3:      "NOR",
					Numeric: "578",
					Name:    "Norway, Kingdom of",
				}, "OM": {
					A2:      "OM",
					A3:      "OMN",
					Numeric: "512",
					Name:    "Oman, Sultanate of",
				}, "PK": {
					A2:      "PK",
					A3:      "PAK",
					Numeric: "586",
					Name:    "Pakistan, Islamic Republic of",
				}, "PW": {
					A2:      "PW",
					A3:      "PLW",
					Numeric: "585",
					Name:    "Palau",
				}, "PS": {
					A2:      "PS",
					A3:      "PSE",
					Numeric: "275",
					Name:    "Palestinian Territory, Occupied",
				}, "PA": {
					A2:      "PA",
					A3:      "PAN",
					Numeric: "591",
					Name:    "Panama, Republic of",
				}, "PG": {
					A2:      "PG",
					A3:      "PNG",
					Numeric: "598",
					Name:    "Papua New Guinea",
				}, "PY": {
					A2:      "PY",
					A3:      "PRY",
					Numeric: "600",
					Name:    "Paraguay, Republic of",
				}, "PE": {
					A2:      "PE",
					A3:      "PER",
					Numeric: "604",
					Name:    "Peru, Republic of",
				}, "PH": {
					A2:      "PH",
					A3:      "PHL",
					Numeric: "608",
					Name:    "Philippines, Republic of the",
				}, "PN": {
					A2:      "PN",
					A3:      "PCN",
					Numeric: "612",
					Name:    "Pitcairn Island",
				}, "PL": {
					A2:      "PL",
					A3:      "POL",
					Numeric: "616",
					Name:    "Poland, Polish People's Republic",
				}, "PT": {
					A2:      "PT",
					A3:      "PRT",
					Numeric: "620",
					Name:    "Portugal, Portuguese Republic",
				}, "PR": {
					A2:      "PR",
					A3:      "PRI",
					Numeric: "630",
					Name:    "Puerto Rico",
				}, "QA": {
					A2:      "QA",
					A3:      "QAT",
					Numeric: "634",
					Name:    "Qatar, State of",
				}, "RE": {
					A2:      "RE",
					A3:      "REU",
					Numeric: "638",
					Name:    "Reunion",
				}, "RO": {
					A2:      "RO",
					A3:      "ROU",
					Numeric: "642",
					Name:    "Romania, Socialist Republic of",
				}, "RU": {
					A2:      "RU",
					A3:      "RUS",
					Numeric: "643",
					Name:    "Russian Federation",
				}, "RW": {
					A2:      "RW",
					A3:      "RWA",
					Numeric: "646",
					Name:    "Rwanda, Rwandese Republic",
				}, "SH": {
					A2:      "SH",
					A3:      "SHN",
					Numeric: "654",
					Name:    "St. Helena",
				}, "KN": {
					A2:      "KN",
					A3:      "KNA",
					Numeric: "659",
					Name:    "St. Kitts and Nevis",
				}, "LC": {
					A2:      "LC",
					A3:      "LCA",
					Numeric: "662",
					Name:    "St. Lucia",
				}, "PM": {
					A2:      "PM",
					A3:      "SPM",
					Numeric: "666",
					Name:    "St. Pierre and Miquelon",
				}, "VC": {
					A2:      "VC",
					A3:      "VCT",
					Numeric: "670",
					Name:    "St. Vincent and the Grenadines",
				}, "WS": {
					A2:      "WS",
					A3:      "WSM",
					Numeric: "882",
					Name:    "Samoa, Independent State of",
				}, "SM": {
					A2:      "SM",
					A3:      "SMR",
					Numeric: "674",
					Name:    "San Marino, Republic of",
				}, "ST": {
					A2:      "ST",
					A3:      "STP",
					Numeric: "678",
					Name:    "Sao Tome and Principe, Democratic Republic of",
				}, "SA": {
					A2:      "SA",
					A3:      "SAU",
					Numeric: "682",
					Name:    "Saudi Arabia, Kingdom of",
				}, "SN": {
					A2:      "SN",
					A3:      "SEN",
					Numeric: "686",
					Name:    "Senegal, Republic of",
				}, "CS": {
					A2:      "CS",
					A3:      "SCG",
					Numeric: "891",
					Name:    "Serbia and Montenegro",
				}, "SC": {
					A2:      "SC",
					A3:      "SYC",
					Numeric: "690",
					Name:    "Seychelles, Republic of",
				}, "SL": {
					A2:      "SL",
					A3:      "SLE",
					Numeric: "694",
					Name:    "Sierra Leone, Republic of",
				}, "SG": {
					A2:      "SG",
					A3:      "SGP",
					Numeric: "702",
					Name:    "Singapore, Republic of",
				}, "SK": {
					A2:      "SK",
					A3:      "SVK",
					Numeric: "703",
					Name:    "Slovakia (Slovak Republic)",
				}, "SI": {
					A2:      "SI",
					A3:      "SVN",
					Numeric: "705",
					Name:    "Slovenia",
				}, "SB": {
					A2:      "SB",
					A3:      "SLB",
					Numeric: "090",
					Name:    "Solomon Islands",
				}, "SO": {
					A2:      "SO",
					A3:      "SOM",
					Numeric: "706",
					Name:    "Somalia, Somali Republic",
				}, "ZA": {
					A2:      "ZA",
					A3:      "ZAF",
					Numeric: "710",
					Name:    "South Africa, Republic of",
				}, "GS": {
					A2:      "GS",
					A3:      "SGS",
					Numeric: "239",
					Name:    "South Georgia and the South Sandwich Islands",
				}, "ES": {
					A2:      "ES",
					A3:      "ESP",
					Numeric: "724",
					Name:    "Spain, Spanish State",
				}, "LK": {
					A2:      "LK",
					A3:      "LKA",
					Numeric: "144",
					Name:    "Sri Lanka, Democratic Socialist Republic of",
				}, "SD": {
					A2:      "SD",
					A3:      "SDN",
					Numeric: "736",
					Name:    "Sudan, Democratic Republic of the",
				}, "SR": {
					A2:      "SR",
					A3:      "SUR",
					Numeric: "740",
					Name:    "Suriname, Republic of",
				}, "SJ": {
					A2:      "SJ",
					A3:      "SJM",
					Numeric: "744",
					Name:    "Svalbard & Jan Mayen Islands",
				}, "SZ": {
					A2:      "SZ",
					A3:      "SWZ",
					Numeric: "748",
					Name:    "Swaziland, Kingdom of",
				}, "SE": {
					A2:      "SE",
					A3:      "SWE",
					Numeric: "752",
					Name:    "Sweden, Kingdom of",
				}, "CH": {
					A2:      "CH",
					A3:      "CHE",
					Numeric: "756",
					Name:    "Switzerland, Swiss Confederation",
				}, "SY": {
					A2:      "SY",
					A3:      "SYR",
					Numeric: "760",
					Name:    "Syrian Arab Republic",
				}, "TW": {
					A2:      "TW",
					A3:      "TWN",
					Numeric: "158",
					Name:    "Taiwan, Province of China",
				}, "TJ": {
					A2:      "TJ",
					A3:      "TJK",
					Numeric: "762",
					Name:    "Tajikistan",
				}, "TZ": {
					A2:      "TZ",
					A3:      "TZA",
					Numeric: "834",
					Name:    "Tanzania, United Republic of",
				}, "TH": {
					A2:      "TH",
					A3:      "THA",
					Numeric: "764",
					Name:    "Thailand, Kingdom of",
				}, "TL": {
					A2:      "TL",
					A3:      "TLS",
					Numeric: "626",
					Name:    "Timor-Leste, Democratic Republic of",
				}, "TG": {
					A2:      "TG",
					A3:      "TGO",
					Numeric: "768",
					Name:    "Togo, Togolese Republic",
				}, "TK": {
					A2:      "TK",
					A3:      "TKL",
					Numeric: "772",
					Name:    "Tokelau (Tokelau Islands)",
				}, "TO": {
					A2:      "TO",
					A3:      "TON",
					Numeric: "776",
					Name:    "Tonga, Kingdom of",
				}, "TT": {
					A2:      "TT",
					A3:      "TTO",
					Numeric: "780",
					Name:    "Trinidad and Tobago, Republic of",
				}, "TN": {
					A2:      "TN",
					A3:      "TUN",
					Numeric: "788",
					Name:    "Tunisia, Republic of",
				}, "TR": {
					A2:      "TR",
					A3:      "TUR",
					Numeric: "792",
					Name:    "Turkey, Republic of",
				}, "TM": {
					A2:      "TM",
					A3:      "TKM",
					Numeric: "795",
					Name:    "Turkmenistan",
				}, "TC": {
					A2:      "TC",
					A3:      "TCA",
					Numeric: "796",
					Name:    "Turks and Caicos Islands",
				}, "TV": {
					A2:      "TV",
					A3:      "TUV",
					Numeric: "798",
					Name:    "Tuvalu",
				}, "VI": {
					A2:      "VI",
					A3:      "VIR",
					Numeric: "850",
					Name:    "US Virgin Islands",
				}, "UG": {
					A2:      "UG",
					A3:      "UGA",
					Numeric: "800",
					Name:    "Uganda, Republic of",
				}, "UA": {
					A2:      "UA",
					A3:      "UKR",
					Numeric: "804",
					Name:    "Ukraine",
				}, "AE": {
					A2:      "AE",
					A3:      "ARE",
					Numeric: "784",
					Name:    "United Arab Emirates",
				}, "GB": {
					A2:      "GB",
					A3:      "GBR",
					Numeric: "826",
					Name:    "United Kingdom of Great Britain & N. Ireland",
				}, "UM": {
					A2:      "UM",
					A3:      "UMI",
					Numeric: "581",
					Name:    "United States Minor Outlying Islands",
				}, "US": {
					A2:      "US",
					A3:      "USA",
					Numeric: "840",
					Name:    "United States of America",
				}, "UY": {
					A2:      "UY",
					A3:      "URY",
					Numeric: "858",
					Name:    "Uruguay, Eastern Republic of",
				}, "UZ": {
					A2:      "UZ",
					A3:      "UZB",
					Numeric: "860",
					Name:    "Uzbekistan",
				}, "VU": {
					A2:      "VU",
					A3:      "VUT",
					Numeric: "548",
					Name:    "Vanuatu",
				}, "VE": {
					A2:      "VE",
					A3:      "VEN",
					Numeric: "862",
					Name:    "Venezuela, Bolivarian Republic of",
				}, "VN": {
					A2:      "VN",
					A3:      "VNM",
					Numeric: "704",
					Name:    "Viet Nam, Socialist Republic of",
				}, "WF": {
					A2:      "WF",
					A3:      "WLF",
					Numeric: "876",
					Name:    "Wallis and Futuna Islands",
				}, "EH": {
					A2:      "EH",
					A3:      "ESH",
					Numeric: "732",
					Name:    "Western Sahara",
				}, "YE": {
					A2:      "YE",
					A3:      "YEM",
					Numeric: "887",
					Name:    "Yemen",
				}, "ZM": {
					A2:      "ZM",
					A3:      "ZMB",
					Numeric: "894",
					Name:    "Zambia, Republic of",
				}, "ZW": {
					A2:      "ZW",
					A3:      "ZWE",
					Numeric: "716",
					Name:    "Zimbabwe",
				},
			},
		},
		Alpha3: &alpha3{
			mu: &sync.RWMutex{},
			data: map[string]*data{
				"AFG": {
					A2:      "AF",
					A3:      "AFG",
					Numeric: "004",
					Name:    "Afghanistan",
				}, "ALB": {
					A2:      "AL",
					A3:      "ALB",
					Numeric: "008",
					Name:    "Albania, People's Socialist Republic of",
				}, "DZA": {
					A2:      "DZ",
					A3:      "DZA",
					Numeric: "012",
					Name:    "Algeria, People's Democratic Republic of",
				}, "ASM": {
					A2:      "AS",
					A3:      "ASM",
					Numeric: "016",
					Name:    "American Samoa",
				}, "AND": {
					A2:      "AD",
					A3:      "AND",
					Numeric: "020",
					Name:    "Andorra, Principality of",
				}, "AGO": {
					A2:      "AO",
					A3:      "AGO",
					Numeric: "024",
					Name:    "Angola, Republic of",
				}, "AIA": {
					A2:      "AI",
					A3:      "AIA",
					Numeric: "660",
					Name:    "Anguilla",
				}, "ATA": {
					A2:      "AQ",
					A3:      "ATA",
					Numeric: "010",
					Name:    "Antarctica (the territory South of 60 deg S)",
				}, "ATG": {
					A2:      "AG",
					A3:      "ATG",
					Numeric: "028",
					Name:    "Antigua and Barbuda",
				}, "ARG": {
					A2:      "AR",
					A3:      "ARG",
					Numeric: "032",
					Name:    "Argentina, Argentine Republic",
				}, "ARM": {
					A2:      "AM",
					A3:      "ARM",
					Numeric: "051",
					Name:    "Armenia",
				}, "ABW": {
					A2:      "AW",
					A3:      "ABW",
					Numeric: "533",
					Name:    "Aruba",
				}, "AUS": {
					A2:      "AU",
					A3:      "AUS",
					Numeric: "036",
					Name:    "Australia, Commonwealth of",
				}, "AUT": {
					A2:      "AT",
					A3:      "AUT",
					Numeric: "040",
					Name:    "Austria, Republic of",
				}, "AZE": {
					A2:      "AZ",
					A3:      "AZE",
					Numeric: "031",
					Name:    "Azerbaijan, Republic of",
				}, "BHS": {
					A2:      "BS",
					A3:      "BHS",
					Numeric: "044",
					Name:    "Bahamas, Commonwealth of the",
				}, "BHR": {
					A2:      "BH",
					A3:      "BHR",
					Numeric: "048",
					Name:    "Bahrain, Kingdom of",
				}, "BGD": {
					A2:      "BD",
					A3:      "BGD",
					Numeric: "050",
					Name:    "Bangladesh, People's Republic of",
				}, "BRB": {
					A2:      "BB",
					A3:      "BRB",
					Numeric: "052",
					Name:    "Barbados",
				}, "BLR": {
					A2:      "BY",
					A3:      "BLR",
					Numeric: "112",
					Name:    "Belarus",
				}, "BEL": {
					A2:      "BE",
					A3:      "BEL",
					Numeric: "056",
					Name:    "Belgium, Kingdom of",
				}, "BLZ": {
					A2:      "BZ",
					A3:      "BLZ",
					Numeric: "084",
					Name:    "Belize",
				}, "BEN": {
					A2:      "BJ",
					A3:      "BEN",
					Numeric: "204",
					Name:    "Benin, People's Republic of",
				}, "BMU": {
					A2:      "BM",
					A3:      "BMU",
					Numeric: "060",
					Name:    "Bermuda",
				}, "BTN": {
					A2:      "BT",
					A3:      "BTN",
					Numeric: "064",
					Name:    "Bhutan, Kingdom of",
				}, "BOL": {
					A2:      "BO",
					A3:      "BOL",
					Numeric: "068",
					Name:    "Bolivia, Republic of",
				}, "BIH": {
					A2:      "BA",
					A3:      "BIH",
					Numeric: "070",
					Name:    "Bosnia and Herzegovina",
				}, "BWA": {
					A2:      "BW",
					A3:      "BWA",
					Numeric: "072",
					Name:    "Botswana, Republic of",
				}, "BVT": {
					A2:      "BV",
					A3:      "BVT",
					Numeric: "074",
					Name:    "Bouvet Island (Bouvetoya)",
				}, "BRA": {
					A2:      "BR",
					A3:      "BRA",
					Numeric: "076",
					Name:    "Brazil, Federative Republic of",
				}, "IOT": {
					A2:      "IO",
					A3:      "IOT",
					Numeric: "086",
					Name:    "British Indian Ocean Territory (Chagos Archipelago)",
				}, "VGB": {
					A2:      "VG",
					A3:      "VGB",
					Numeric: "092",
					Name:    "British Virgin Islands",
				}, "BRN": {
					A2:      "BN",
					A3:      "BRN",
					Numeric: "096",
					Name:    "Brunei Darussalam",
				}, "BGR": {
					A2:      "BG",
					A3:      "BGR",
					Numeric: "100",
					Name:    "Bulgaria, People's Republic of",
				}, "BFA": {
					A2:      "BF",
					A3:      "BFA",
					Numeric: "854",
					Name:    "Burkina Faso",
				}, "BDI": {
					A2:      "BI",
					A3:      "BDI",
					Numeric: "108",
					Name:    "Burundi, Republic of",
				}, "KHM": {
					A2:      "KH",
					A3:      "KHM",
					Numeric: "116",
					Name:    "Cambodia, Kingdom of",
				}, "CMR": {
					A2:      "CM",
					A3:      "CMR",
					Numeric: "120",
					Name:    "Cameroon, United Republic of",
				}, "CAN": {
					A2:      "CA",
					A3:      "CAN",
					Numeric: "124",
					Name:    "Canada",
				}, "CPV": {
					A2:      "CV",
					A3:      "CPV",
					Numeric: "132",
					Name:    "Cape Verde, Republic of",
				}, "CYM": {
					A2:      "KY",
					A3:      "CYM",
					Numeric: "136",
					Name:    "Cayman Islands",
				}, "CAF": {
					A2:      "CF",
					A3:      "CAF",
					Numeric: "140",
					Name:    "Central African Republic",
				}, "TCD": {
					A2:      "TD",
					A3:      "TCD",
					Numeric: "148",
					Name:    "Chad, Republic of",
				}, "CHL": {
					A2:      "CL",
					A3:      "CHL",
					Numeric: "152",
					Name:    "Chile, Republic of",
				}, "CHN": {
					A2:      "CN",
					A3:      "CHN",
					Numeric: "156",
					Name:    "China, People's Republic of",
				}, "CXR": {
					A2:      "CX",
					A3:      "CXR",
					Numeric: "162",
					Name:    "Christmas Island",
				}, "CCK": {
					A2:      "CC",
					A3:      "CCK",
					Numeric: "166",
					Name:    "Cocos (Keeling) Islands",
				}, "COL": {
					A2:      "CO",
					A3:      "COL",
					Numeric: "170",
					Name:    "Colombia, Republic of",
				}, "COM": {
					A2:      "KM",
					A3:      "COM",
					Numeric: "174",
					Name:    "Comoros, Union of the",
				}, "COD": {
					A2:      "CD",
					A3:      "COD",
					Numeric: "180",
					Name:    "Congo, Democratic Republic of",
				}, "COG": {
					A2:      "CG",
					A3:      "COG",
					Numeric: "178",
					Name:    "Congo, People's Republic of",
				}, "COK": {
					A2:      "CK",
					A3:      "COK",
					Numeric: "184",
					Name:    "Cook Islands",
				}, "CRI": {
					A2:      "CR",
					A3:      "CRI",
					Numeric: "188",
					Name:    "Costa Rica, Republic of",
				}, "CIV": {
					A2:      "CI",
					A3:      "CIV",
					Numeric: "384",
					Name:    "Cote D'Ivoire, Ivory Coast, Republic of the",
				}, "CUB": {
					A2:      "CU",
					A3:      "CUB",
					Numeric: "192",
					Name:    "Cuba, Republic of",
				}, "CYP": {
					A2:      "CY",
					A3:      "CYP",
					Numeric: "196",
					Name:    "Cyprus, Republic of",
				}, "CZE": {
					A2:      "CZ",
					A3:      "CZE",
					Numeric: "203",
					Name:    "Czech Republic",
				}, "DNK": {
					A2:      "DK",
					A3:      "DNK",
					Numeric: "208",
					Name:    "Denmark, Kingdom of",
				}, "DJI": {
					A2:      "DJ",
					A3:      "DJI",
					Numeric: "262",
					Name:    "Djibouti, Republic of",
				}, "DMA": {
					A2:      "DM",
					A3:      "DMA",
					Numeric: "212",
					Name:    "Dominica, Commonwealth of",
				}, "DOM": {
					A2:      "DO",
					A3:      "DOM",
					Numeric: "214",
					Name:    "Dominican Republic",
				}, "ECU": {
					A2:      "EC",
					A3:      "ECU",
					Numeric: "218",
					Name:    "Ecuador, Republic of",
				}, "EGY": {
					A2:      "EG",
					A3:      "EGY",
					Numeric: "818",
					Name:    "Egypt, Arab Republic of",
				}, "SLV": {
					A2:      "SV",
					A3:      "SLV",
					Numeric: "222",
					Name:    "El Salvador, Republic of",
				}, "GNQ": {
					A2:      "GQ",
					A3:      "GNQ",
					Numeric: "226",
					Name:    "Equatorial Guinea, Republic of",
				}, "ERI": {
					A2:      "ER",
					A3:      "ERI",
					Numeric: "232",
					Name:    "Eritrea",
				}, "EST": {
					A2:      "EE",
					A3:      "EST",
					Numeric: "233",
					Name:    "Estonia",
				}, "ETH": {
					A2:      "ET",
					A3:      "ETH",
					Numeric: "231",
					Name:    "Ethiopia",
				}, "FRO": {
					A2:      "FO",
					A3:      "FRO",
					Numeric: "234",
					Name:    "Faeroe Islands",
				}, "FLK": {
					A2:      "FK",
					A3:      "FLK",
					Numeric: "238",
					Name:    "Falkland Islands (Malvinas)",
				}, "FJI": {
					A2:      "FJ",
					A3:      "FJI",
					Numeric: "242",
					Name:    "Fiji, Republic of the Fiji Islands",
				}, "FIN": {
					A2:      "FI",
					A3:      "FIN",
					Numeric: "246",
					Name:    "Finland, Republic of",
				}, "FRA": {
					A2:      "FR",
					A3:      "FRA",
					Numeric: "250",
					Name:    "France, French Republic",
				}, "GUF": {
					A2:      "GF",
					A3:      "GUF",
					Numeric: "254",
					Name:    "French Guiana",
				}, "PYF": {
					A2:      "PF",
					A3:      "PYF",
					Numeric: "258",
					Name:    "French Polynesia",
				}, "ATF": {
					A2:      "TF",
					A3:      "ATF",
					Numeric: "260",
					Name:    "French Southern Territories",
				}, "GAB": {
					A2:      "GA",
					A3:      "GAB",
					Numeric: "266",
					Name:    "Gabon, Gabonese Republic",
				}, "GMB": {
					A2:      "GM",
					A3:      "GMB",
					Numeric: "270",
					Name:    "Gambia, Republic of the",
				}, "GEO": {
					A2:      "GE",
					A3:      "GEO",
					Numeric: "268",
					Name:    "Georgia",
				}, "DEU": {
					A2:      "DE",
					A3:      "DEU",
					Numeric: "276",
					Name:    "Germany",
				}, "GHA": {
					A2:      "GH",
					A3:      "GHA",
					Numeric: "288",
					Name:    "Ghana, Republic of",
				}, "GIB": {
					A2:      "GI",
					A3:      "GIB",
					Numeric: "292",
					Name:    "Gibraltar",
				}, "GRC": {
					A2:      "GR",
					A3:      "GRC",
					Numeric: "300",
					Name:    "Greece, Hellenic Republic",
				}, "GRL": {
					A2:      "GL",
					A3:      "GRL",
					Numeric: "304",
					Name:    "Greenland",
				}, "GRD": {
					A2:      "GD",
					A3:      "GRD",
					Numeric: "308",
					Name:    "Grenada",
				}, "GLP": {
					A2:      "GP",
					A3:      "GLP",
					Numeric: "312",
					Name:    "Guadaloupe",
				}, "GUM": {
					A2:      "GU",
					A3:      "GUM",
					Numeric: "316",
					Name:    "Guam",
				}, "GTM": {
					A2:      "GT",
					A3:      "GTM",
					Numeric: "320",
					Name:    "Guatemala, Republic of",
				}, "GIN": {
					A2:      "GN",
					A3:      "GIN",
					Numeric: "324",
					Name:    "Guinea, Revolutionary People's Rep'c of",
				}, "GNB": {
					A2:      "GW",
					A3:      "GNB",
					Numeric: "624",
					Name:    "Guinea-Bissau, Republic of",
				}, "GUY": {
					A2:      "GY",
					A3:      "GUY",
					Numeric: "328",
					Name:    "Guyana, Republic of",
				}, "HTI": {
					A2:      "HT",
					A3:      "HTI",
					Numeric: "332",
					Name:    "Haiti, Republic of",
				}, "HMD": {
					A2:      "HM",
					A3:      "HMD",
					Numeric: "334",
					Name:    "Heard and McDonald Islands",
				}, "VAT": {
					A2:      "VA",
					A3:      "VAT",
					Numeric: "336",
					Name:    "Holy See (Vatican City State)",
				}, "HND": {
					A2:      "HN",
					A3:      "HND",
					Numeric: "340",
					Name:    "Honduras, Republic of",
				}, "HKG": {
					A2:      "HK",
					A3:      "HKG",
					Numeric: "344",
					Name:    "Hong Kong, Special Administrative Region of China",
				}, "HRV": {
					A2:      "HR",
					A3:      "HRV",
					Numeric: "191",
					Name:    "Hrvatska (Croatia)",
				}, "HUN": {
					A2:      "HU",
					A3:      "HUN",
					Numeric: "348",
					Name:    "Hungary, Hungarian People's Republic",
				}, "ISL": {
					A2:      "IS",
					A3:      "ISL",
					Numeric: "352",
					Name:    "Iceland, Republic of",
				}, "IND": {
					A2:      "IN",
					A3:      "IND",
					Numeric: "356",
					Name:    "India, Republic of",
				}, "IDN": {
					A2:      "ID",
					A3:      "IDN",
					Numeric: "360",
					Name:    "Indonesia, Republic of",
				}, "IRN": {
					A2:      "IR",
					A3:      "IRN",
					Numeric: "364",
					Name:    "Iran, Islamic Republic of",
				}, "IRQ": {
					A2:      "IQ",
					A3:      "IRQ",
					Numeric: "368",
					Name:    "Iraq, Republic of",
				}, "IRL": {
					A2:      "IE",
					A3:      "IRL",
					Numeric: "372",
					Name:    "Ireland",
				}, "ISR": {
					A2:      "IL",
					A3:      "ISR",
					Numeric: "376",
					Name:    "Israel, State of",
				}, "ITA": {
					A2:      "IT",
					A3:      "ITA",
					Numeric: "380",
					Name:    "Italy, Italian Republic",
				}, "JAM": {
					A2:      "JM",
					A3:      "JAM",
					Numeric: "388",
					Name:    "Jamaica",
				}, "JPN": {
					A2:      "JP",
					A3:      "JPN",
					Numeric: "392",
					Name:    "Japan",
				}, "JOR": {
					A2:      "JO",
					A3:      "JOR",
					Numeric: "400",
					Name:    "Jordan, Hashemite Kingdom of",
				}, "KAZ": {
					A2:      "KZ",
					A3:      "KAZ",
					Numeric: "398",
					Name:    "Kazakhstan, Republic of",
				}, "KEN": {
					A2:      "KE",
					A3:      "KEN",
					Numeric: "404",
					Name:    "Kenya, Republic of",
				}, "KIR": {
					A2:      "KI",
					A3:      "KIR",
					Numeric: "296",
					Name:    "Kiribati, Republic of",
				}, "PRK": {
					A2:      "KP",
					A3:      "PRK",
					Numeric: "408",
					Name:    "Korea, Democratic People's Republic of",
				}, "KOR": {
					A2:      "KR",
					A3:      "KOR",
					Numeric: "410",
					Name:    "Korea, Republic of",
				}, "KWT": {
					A2:      "KW",
					A3:      "KWT",
					Numeric: "414",
					Name:    "Kuwait, State of",
				}, "KGZ": {
					A2:      "KG",
					A3:      "KGZ",
					Numeric: "417",
					Name:    "Kyrgyz Republic",
				}, "LAO": {
					A2:      "LA",
					A3:      "LAO",
					Numeric: "418",
					Name:    "Lao People's Democratic Republic",
				}, "LVA": {
					A2:      "LV",
					A3:      "LVA",
					Numeric: "428",
					Name:    "Latvia",
				}, "LBN": {
					A2:      "LB",
					A3:      "LBN",
					Numeric: "422",
					Name:    "Lebanon, Lebanese Republic",
				}, "LSO": {
					A2:      "LS",
					A3:      "LSO",
					Numeric: "426",
					Name:    "Lesotho, Kingdom of",
				}, "LBR": {
					A2:      "LR",
					A3:      "LBR",
					Numeric: "430",
					Name:    "Liberia, Republic of",
				}, "LBY": {
					A2:      "LY",
					A3:      "LBY",
					Numeric: "434",
					Name:    "Libyan Arab Jamahiriya",
				}, "LIE": {
					A2:      "LI",
					A3:      "LIE",
					Numeric: "438",
					Name:    "Liechtenstein, Principality of",
				}, "LTU": {
					A2:      "LT",
					A3:      "LTU",
					Numeric: "440",
					Name:    "Lithuania",
				}, "LUX": {
					A2:      "LU",
					A3:      "LUX",
					Numeric: "442",
					Name:    "Luxembourg, Grand Duchy of",
				}, "MAC": {
					A2:      "MO",
					A3:      "MAC",
					Numeric: "446",
					Name:    "Macao, Special Administrative Region of China",
				}, "MKD": {
					A2:      "MK",
					A3:      "MKD",
					Numeric: "807",
					Name:    "Macedonia, the former Yugoslav Republic of",
				}, "MDG": {
					A2:      "MG",
					A3:      "MDG",
					Numeric: "450",
					Name:    "Madagascar, Republic of",
				}, "MWI": {
					A2:      "MW",
					A3:      "MWI",
					Numeric: "454",
					Name:    "Malawi, Republic of",
				}, "MYS": {
					A2:      "MY",
					A3:      "MYS",
					Numeric: "458",
					Name:    "Malaysia",
				}, "MDV": {
					A2:      "MV",
					A3:      "MDV",
					Numeric: "462",
					Name:    "Maldives, Republic of",
				}, "MLI": {
					A2:      "ML",
					A3:      "MLI",
					Numeric: "466",
					Name:    "Mali, Republic of",
				}, "MLT": {
					A2:      "MT",
					A3:      "MLT",
					Numeric: "470",
					Name:    "Malta, Republic of",
				}, "MHL": {
					A2:      "MH",
					A3:      "MHL",
					Numeric: "584",
					Name:    "Marshall Islands",
				}, "MTQ": {
					A2:      "MQ",
					A3:      "MTQ",
					Numeric: "474",
					Name:    "Martinique",
				}, "MRT": {
					A2:      "MR",
					A3:      "MRT",
					Numeric: "478",
					Name:    "Mauritania, Islamic Republic of",
				}, "MUS": {
					A2:      "MU",
					A3:      "MUS",
					Numeric: "480",
					Name:    "Mauritius",
				}, "MYT": {
					A2:      "YT",
					A3:      "MYT",
					Numeric: "175",
					Name:    "Mayotte",
				}, "MEX": {
					A2:      "MX",
					A3:      "MEX",
					Numeric: "484",
					Name:    "Mexico, United Mexican States",
				}, "FSM": {
					A2:      "FM",
					A3:      "FSM",
					Numeric: "583",
					Name:    "Micronesia, Federated States of",
				}, "MDA": {
					A2:      "MD",
					A3:      "MDA",
					Numeric: "498",
					Name:    "Moldova, Republic of",
				}, "MCO": {
					A2:      "MC",
					A3:      "MCO",
					Numeric: "492",
					Name:    "Monaco, Principality of",
				}, "MNG": {
					A2:      "MN",
					A3:      "MNG",
					Numeric: "496",
					Name:    "Mongolia, Mongolian People's Republic",
				}, "MSR": {
					A2:      "MS",
					A3:      "MSR",
					Numeric: "500",
					Name:    "Montserrat",
				}, "MAR": {
					A2:      "MA",
					A3:      "MAR",
					Numeric: "504",
					Name:    "Morocco, Kingdom of",
				}, "MOZ": {
					A2:      "MZ",
					A3:      "MOZ",
					Numeric: "508",
					Name:    "Mozambique, People's Republic of",
				}, "MMR": {
					A2:      "MM",
					A3:      "MMR",
					Numeric: "104",
					Name:    "Myanmar",
				}, "NAM": {
					A2:      "NA",
					A3:      "NAM",
					Numeric: "516",
					Name:    "Namibia",
				}, "NRU": {
					A2:      "NR",
					A3:      "NRU",
					Numeric: "520",
					Name:    "Nauru, Republic of",
				}, "NPL": {
					A2:      "NP",
					A3:      "NPL",
					Numeric: "524",
					Name:    "Nepal, Kingdom of",
				}, "ANT": {
					A2:      "AN",
					A3:      "ANT",
					Numeric: "530",
					Name:    "Netherlands Antilles",
				}, "NLD": {
					A2:      "NL",
					A3:      "NLD",
					Numeric: "528",
					Name:    "Netherlands, Kingdom of the",
				}, "NCL": {
					A2:      "NC",
					A3:      "NCL",
					Numeric: "540",
					Name:    "New Caledonia",
				}, "NZL": {
					A2:      "NZ",
					A3:      "NZL",
					Numeric: "554",
					Name:    "New Zealand",
				}, "NIC": {
					A2:      "NI",
					A3:      "NIC",
					Numeric: "558",
					Name:    "Nicaragua, Republic of",
				}, "NER": {
					A2:      "NE",
					A3:      "NER",
					Numeric: "562",
					Name:    "Niger, Republic of the",
				}, "NGA": {
					A2:      "NG",
					A3:      "NGA",
					Numeric: "566",
					Name:    "Nigeria, Federal Republic of",
				}, "NIU": {
					A2:      "NU",
					A3:      "NIU",
					Numeric: "570",
					Name:    "Niue, Republic of",
				}, "NFK": {
					A2:      "NF",
					A3:      "NFK",
					Numeric: "574",
					Name:    "Norfolk Island",
				}, "MNP": {
					A2:      "MP",
					A3:      "MNP",
					Numeric: "580",
					Name:    "Northern Mariana Islands",
				}, "NOR": {
					A2:      "NO",
					A3:      "NOR",
					Numeric: "578",
					Name:    "Norway, Kingdom of",
				}, "OMN": {
					A2:      "OM",
					A3:      "OMN",
					Numeric: "512",
					Name:    "Oman, Sultanate of",
				}, "PAK": {
					A2:      "PK",
					A3:      "PAK",
					Numeric: "586",
					Name:    "Pakistan, Islamic Republic of",
				}, "PLW": {
					A2:      "PW",
					A3:      "PLW",
					Numeric: "585",
					Name:    "Palau",
				}, "PSE": {
					A2:      "PS",
					A3:      "PSE",
					Numeric: "275",
					Name:    "Palestinian Territory, Occupied",
				}, "PAN": {
					A2:      "PA",
					A3:      "PAN",
					Numeric: "591",
					Name:    "Panama, Republic of",
				}, "PNG": {
					A2:      "PG",
					A3:      "PNG",
					Numeric: "598",
					Name:    "Papua New Guinea",
				}, "PRY": {
					A2:      "PY",
					A3:      "PRY",
					Numeric: "600",
					Name:    "Paraguay, Republic of",
				}, "PER": {
					A2:      "PE",
					A3:      "PER",
					Numeric: "604",
					Name:    "Peru, Republic of",
				}, "PHL": {
					A2:      "PH",
					A3:      "PHL",
					Numeric: "608",
					Name:    "Philippines, Republic of the",
				}, "PCN": {
					A2:      "PN",
					A3:      "PCN",
					Numeric: "612",
					Name:    "Pitcairn Island",
				}, "POL": {
					A2:      "PL",
					A3:      "POL",
					Numeric: "616",
					Name:    "Poland, Polish People's Republic",
				}, "PRT": {
					A2:      "PT",
					A3:      "PRT",
					Numeric: "620",
					Name:    "Portugal, Portuguese Republic",
				}, "PRI": {
					A2:      "PR",
					A3:      "PRI",
					Numeric: "630",
					Name:    "Puerto Rico",
				}, "QAT": {
					A2:      "QA",
					A3:      "QAT",
					Numeric: "634",
					Name:    "Qatar, State of",
				}, "REU": {
					A2:      "RE",
					A3:      "REU",
					Numeric: "638",
					Name:    "Reunion",
				}, "ROU": {
					A2:      "RO",
					A3:      "ROU",
					Numeric: "642",
					Name:    "Romania, Socialist Republic of",
				}, "RUS": {
					A2:      "RU",
					A3:      "RUS",
					Numeric: "643",
					Name:    "Russian Federation",
				}, "RWA": {
					A2:      "RW",
					A3:      "RWA",
					Numeric: "646",
					Name:    "Rwanda, Rwandese Republic",
				}, "SHN": {
					A2:      "SH",
					A3:      "SHN",
					Numeric: "654",
					Name:    "St. Helena",
				}, "KNA": {
					A2:      "KN",
					A3:      "KNA",
					Numeric: "659",
					Name:    "St. Kitts and Nevis",
				}, "LCA": {
					A2:      "LC",
					A3:      "LCA",
					Numeric: "662",
					Name:    "St. Lucia",
				}, "SPM": {
					A2:      "PM",
					A3:      "SPM",
					Numeric: "666",
					Name:    "St. Pierre and Miquelon",
				}, "VCT": {
					A2:      "VC",
					A3:      "VCT",
					Numeric: "670",
					Name:    "St. Vincent and the Grenadines",
				}, "WSM": {
					A2:      "WS",
					A3:      "WSM",
					Numeric: "882",
					Name:    "Samoa, Independent State of",
				}, "SMR": {
					A2:      "SM",
					A3:      "SMR",
					Numeric: "674",
					Name:    "San Marino, Republic of",
				}, "STP": {
					A2:      "ST",
					A3:      "STP",
					Numeric: "678",
					Name:    "Sao Tome and Principe, Democratic Republic of",
				}, "SAU": {
					A2:      "SA",
					A3:      "SAU",
					Numeric: "682",
					Name:    "Saudi Arabia, Kingdom of",
				}, "SEN": {
					A2:      "SN",
					A3:      "SEN",
					Numeric: "686",
					Name:    "Senegal, Republic of",
				}, "SCG": {
					A2:      "CS",
					A3:      "SCG",
					Numeric: "891",
					Name:    "Serbia and Montenegro",
				}, "SYC": {
					A2:      "SC",
					A3:      "SYC",
					Numeric: "690",
					Name:    "Seychelles, Republic of",
				}, "SLE": {
					A2:      "SL",
					A3:      "SLE",
					Numeric: "694",
					Name:    "Sierra Leone, Republic of",
				}, "SGP": {
					A2:      "SG",
					A3:      "SGP",
					Numeric: "702",
					Name:    "Singapore, Republic of",
				}, "SVK": {
					A2:      "SK",
					A3:      "SVK",
					Numeric: "703",
					Name:    "Slovakia (Slovak Republic)",
				}, "SVN": {
					A2:      "SI",
					A3:      "SVN",
					Numeric: "705",
					Name:    "Slovenia",
				}, "SLB": {
					A2:      "SB",
					A3:      "SLB",
					Numeric: "090",
					Name:    "Solomon Islands",
				}, "SOM": {
					A2:      "SO",
					A3:      "SOM",
					Numeric: "706",
					Name:    "Somalia, Somali Republic",
				}, "ZAF": {
					A2:      "ZA",
					A3:      "ZAF",
					Numeric: "710",
					Name:    "South Africa, Republic of",
				}, "SGS": {
					A2:      "GS",
					A3:      "SGS",
					Numeric: "239",
					Name:    "South Georgia and the South Sandwich Islands",
				}, "ESP": {
					A2:      "ES",
					A3:      "ESP",
					Numeric: "724",
					Name:    "Spain, Spanish State",
				}, "LKA": {
					A2:      "LK",
					A3:      "LKA",
					Numeric: "144",
					Name:    "Sri Lanka, Democratic Socialist Republic of",
				}, "SDN": {
					A2:      "SD",
					A3:      "SDN",
					Numeric: "736",
					Name:    "Sudan, Democratic Republic of the",
				}, "SUR": {
					A2:      "SR",
					A3:      "SUR",
					Numeric: "740",
					Name:    "Suriname, Republic of",
				}, "SJM": {
					A2:      "SJ",
					A3:      "SJM",
					Numeric: "744",
					Name:    "Svalbard & Jan Mayen Islands",
				}, "SWZ": {
					A2:      "SZ",
					A3:      "SWZ",
					Numeric: "748",
					Name:    "Swaziland, Kingdom of",
				}, "SWE": {
					A2:      "SE",
					A3:      "SWE",
					Numeric: "752",
					Name:    "Sweden, Kingdom of",
				}, "CHE": {
					A2:      "CH",
					A3:      "CHE",
					Numeric: "756",
					Name:    "Switzerland, Swiss Confederation",
				}, "SYR": {
					A2:      "SY",
					A3:      "SYR",
					Numeric: "760",
					Name:    "Syrian Arab Republic",
				}, "TWN": {
					A2:      "TW",
					A3:      "TWN",
					Numeric: "158",
					Name:    "Taiwan, Province of China",
				}, "TJK": {
					A2:      "TJ",
					A3:      "TJK",
					Numeric: "762",
					Name:    "Tajikistan",
				}, "TZA": {
					A2:      "TZ",
					A3:      "TZA",
					Numeric: "834",
					Name:    "Tanzania, United Republic of",
				}, "THA": {
					A2:      "TH",
					A3:      "THA",
					Numeric: "764",
					Name:    "Thailand, Kingdom of",
				}, "TLS": {
					A2:      "TL",
					A3:      "TLS",
					Numeric: "626",
					Name:    "Timor-Leste, Democratic Republic of",
				}, "TGO": {
					A2:      "TG",
					A3:      "TGO",
					Numeric: "768",
					Name:    "Togo, Togolese Republic",
				}, "TKL": {
					A2:      "TK",
					A3:      "TKL",
					Numeric: "772",
					Name:    "Tokelau (Tokelau Islands)",
				}, "TON": {
					A2:      "TO",
					A3:      "TON",
					Numeric: "776",
					Name:    "Tonga, Kingdom of",
				}, "TTO": {
					A2:      "TT",
					A3:      "TTO",
					Numeric: "780",
					Name:    "Trinidad and Tobago, Republic of",
				}, "TUN": {
					A2:      "TN",
					A3:      "TUN",
					Numeric: "788",
					Name:    "Tunisia, Republic of",
				}, "TUR": {
					A2:      "TR",
					A3:      "TUR",
					Numeric: "792",
					Name:    "Turkey, Republic of",
				}, "TKM": {
					A2:      "TM",
					A3:      "TKM",
					Numeric: "795",
					Name:    "Turkmenistan",
				}, "TCA": {
					A2:      "TC",
					A3:      "TCA",
					Numeric: "796",
					Name:    "Turks and Caicos Islands",
				}, "TUV": {
					A2:      "TV",
					A3:      "TUV",
					Numeric: "798",
					Name:    "Tuvalu",
				}, "VIR": {
					A2:      "VI",
					A3:      "VIR",
					Numeric: "850",
					Name:    "US Virgin Islands",
				}, "UGA": {
					A2:      "UG",
					A3:      "UGA",
					Numeric: "800",
					Name:    "Uganda, Republic of",
				}, "UKR": {
					A2:      "UA",
					A3:      "UKR",
					Numeric: "804",
					Name:    "Ukraine",
				}, "ARE": {
					A2:      "AE",
					A3:      "ARE",
					Numeric: "784",
					Name:    "United Arab Emirates",
				}, "GBR": {
					A2:      "GB",
					A3:      "GBR",
					Numeric: "826",
					Name:    "United Kingdom of Great Britain & N. Ireland",
				}, "UMI": {
					A2:      "UM",
					A3:      "UMI",
					Numeric: "581",
					Name:    "United States Minor Outlying Islands",
				}, "USA": {
					A2:      "US",
					A3:      "USA",
					Numeric: "840",
					Name:    "United States of America",
				}, "URY": {
					A2:      "UY",
					A3:      "URY",
					Numeric: "858",
					Name:    "Uruguay, Eastern Republic of",
				}, "UZB": {
					A2:      "UZ",
					A3:      "UZB",
					Numeric: "860",
					Name:    "Uzbekistan",
				}, "VUT": {
					A2:      "VU",
					A3:      "VUT",
					Numeric: "548",
					Name:    "Vanuatu",
				}, "VEN": {
					A2:      "VE",
					A3:      "VEN",
					Numeric: "862",
					Name:    "Venezuela, Bolivarian Republic of",
				}, "VNM": {
					A2:      "VN",
					A3:      "VNM",
					Numeric: "704",
					Name:    "Viet Nam, Socialist Republic of",
				}, "WLF": {
					A2:      "WF",
					A3:      "WLF",
					Numeric: "876",
					Name:    "Wallis and Futuna Islands",
				}, "ESH": {
					A2:      "EH",
					A3:      "ESH",
					Numeric: "732",
					Name:    "Western Sahara",
				}, "YEM": {
					A2:      "YE",
					A3:      "YEM",
					Numeric: "887",
					Name:    "Yemen",
				}, "ZMB": {
					A2:      "ZM",
					A3:      "ZMB",
					Numeric: "894",
					Name:    "Zambia, Republic of",
				}, "ZWE": {
					A2:      "ZW",
					A3:      "ZWE",
					Numeric: "716",
					Name:    "Zimbabwe",
				},
			},
		},
	}
}
