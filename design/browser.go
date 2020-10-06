package design

type BrowserHistory struct {
	History []string
	Cur     int
	Size    int
}

func NewBrowser(homepage string) *BrowserHistory {
	return &BrowserHistory{[]string{homepage}, 0, 1}
}

func (h *BrowserHistory) Visit(url string) {
	h.Cur++
	h.History = append(h.History[:h.Cur], url)
	h.Size = len(h.History)
}

func (h *BrowserHistory) Back(steps int) string {
	h.Cur -= steps
	if h.Cur < 0 {
		h.Cur = 0
	}
	return h.History[h.Cur]
}

func (h *BrowserHistory) Forward(steps int) string {
	h.Cur += steps
	if h.Cur >= h.Size {
		h.Cur = h.Size - 1
	}
	return h.History[h.Cur]
}
