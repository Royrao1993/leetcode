package design

type ProductOfNumbers struct {
	Prefix []int
}

func NewProductNums() *ProductOfNumbers {
	return &ProductOfNumbers{[]int{}}
}

func (p *ProductOfNumbers) Add(num int) {
	l := len(p.Prefix)
	if num == 0 {
		p.Prefix = p.Prefix[:0]
	} else {
		if l == 0 {
			p.Prefix = append(p.Prefix, num)
		} else {
			p.Prefix = append(p.Prefix, num*p.Prefix[l-1])
		}
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	l := len(p.Prefix)
	if k > l {
		return 0
	} else if k == l {
		return p.Prefix[l-1]
	} else {
		return p.Prefix[l-1] / p.Prefix[l-k-1]
	}
}
