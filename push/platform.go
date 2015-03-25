package push

import "github.com/DeanThompson/jpush-api-go-client/common"

type Platform struct {
	value []string
}

func NewPlatform() *Platform {
	return &Platform{}
}

// 如果有 "all"，只会返回字符串 "all"
// 其他情况都是 []string{}，包含具体的平台参数
func (p *Platform) Value() interface{} {
	if p.has(ALL) {
		return ALL
	}

	return p.value
}

func (p *Platform) All() {
	p.value = []string{ALL}
}

// 添加 platform，可选传参： "all", "ios", "android", "winphone"
func (p *Platform) Add(platforms ...string) error {
	if len(platforms) == 0 {
		return nil
	}

	if p.value == nil {
		p.value = make([]string, 0)
	}

	// 去重
	platforms = common.UniqString(platforms)

	for _, platform := range platforms {
		if !isValidPlatform(platform) {
			return common.ErrInvalidPlatform
		}

		// 不要重复添加，如果有 set 就方便了
		if !p.has(platform) {
			p.value = append(p.value, platform)
		}
	}
	return nil
}

func (p *Platform) has(platform string) bool {
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

func isValidPlatform(platform string) bool {
	switch platform {
	case ALL, PLATFORM_IOS, PLATFORM_ANDROID, PLATFORM_WP:
		return true
	}
	return false
}
