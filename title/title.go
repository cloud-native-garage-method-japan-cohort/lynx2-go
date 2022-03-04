package title

type title struct {
}

type Title interface {
	Get() string
}

// コンストラクタ
func NewTitle() Title {
	var t title
	return &t
}

func (t *title) Get() string {
	return "Lynx Search 2"
}
