package jpush

// 应用内消息。或者称作：自定义消息，透传消息。
// 此部分内容不会展示到通知栏上，JPush SDK 收到消息内容后透传给 App。
// 收到消息后 App 需要自行处理。
type Message struct {
	Content     string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func NewMessage(content string) *Message {
	return &Message{Content: content}
}

func (m *Message) Validate() error {
	if m.Content == "" {
		return ErrMessageContentMissing
	}

	return nil
}

func (m *Message) AddExtra(key string, value interface{}) {
	if m.Extras == nil {
		m.Extras = make(map[string]interface{})
	}
	m.Extras[key] = value
}
