package jpush

func isValidPlatform(platform string) bool {
	switch platform {
	case ALL, PLATFORM_IOS, PLATFORM_ANDROID, PLATFORM_WP:
		return true
	}
	return false
}

type Platform struct {
	value []string
}

func NewPlatform() *Platform {
	return &Platform{}
}

// "all" or ["ios", "android"]
func (p *Platform) Value() interface{} {
	if p.Has(ALL) {
		return ALL
	}

	return p.value
}

func (p *Platform) All() {
	p.value = []string{ALL}
}

func (p *Platform) Add(platform string) error {
	if !isValidPlatform(platform) {
		return ErrInvalidPlatform
	}
	if !p.Has(platform) {
		if p.value == nil {
			p.value = make([]string, 0)
		}
		p.value = append(p.value, platform)
	}
	return nil
}

func (p *Platform) Has(platform string) bool {
	if p.value == nil {
		return false
	}
	for _, v := range p.value {
		if v == ALL || v == platform {
			return true
		}
	}
	return false
}
