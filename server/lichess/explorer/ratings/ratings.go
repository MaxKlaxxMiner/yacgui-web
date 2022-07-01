package ratings

type Ratings byte

const (
	None  Ratings = 0x00
	R1600 Ratings = 0x01
	R1800 Ratings = 0x02
	R2000 Ratings = 0x04
	R2200 Ratings = 0x08
	R2500 Ratings = 0x10
	All           = R1600 | R1800 | R2000 | R2200 | R2500

	Min1600 = R1600 | Min1800
	Min1800 = R1800 | Min2000
	Min2000 = R2000 | Min2200
	Min2200 = R2200 | Min2500
	Min2500 = R2500

	Max1600 = R1600
	Max1800 = R1800 | Max1600
	Max2000 = R2000 | Max1800
	Max2200 = R2200 | Max2000
	Max2500 = R2500 | Max2200
)

func (r Ratings) String() string {
	var result string

	switch {
	case r&R1600 != 0:
		result += ",1600"
		fallthrough
	case r&R1800 != 0:
		result += ",1800"
		fallthrough
	case r&R2000 != 0:
		result += ",2000"
		fallthrough
	case r&R2200 != 0:
		result += ",2200"
		fallthrough
	case r&R2500 != 0:
		result += ",2500"
	}

	if len(result) == 0 {
		return result
	}
	return result[1:] // remove first ','
}
