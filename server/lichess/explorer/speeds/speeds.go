package speeds

type Speeds byte

const (
	None             Speeds = 0x00
	UltraBullet      Speeds = 0x01
	Bullet           Speeds = 0x02
	Blitz            Speeds = 0x04
	Rapid            Speeds = 0x08
	Classical        Speeds = 0x10
	Correspondence   Speeds = 0x20
	All                     = UltraBullet | Bullet | SlowerThanBullet
	SlowerThanBullet        = Blitz | SlowerThanBlitz
	SlowerThanBlitz         = Rapid | SlowerThanRapid
	SlowerThanRapid         = Classical | Correspondence
)

func (s Speeds) String() string {
	var result string

	switch {
	case s&UltraBullet != 0:
		result += ",ultraBullet"
		fallthrough
	case s&Bullet != 0:
		result += ",bullet"
		fallthrough
	case s&Blitz != 0:
		result += ",blitz"
		fallthrough
	case s&Rapid != 0:
		result += ",rapid"
		fallthrough
	case s&Classical != 0:
		result += ",classical"
		fallthrough
	case s&Correspondence != 0:
		result += ",correspondence"
	}

	if len(result) == 0 {
		return result
	}
	return result[1:] // remove first ','
}
