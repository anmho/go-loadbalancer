package balancer

type Rule struct {
	ttlMS int
	rate int
}

func NewRule(ttlMS int, rate int) Rule {
	return Rule{ttlMS: ttlMS, rate: rate}
}

type Balancer struct {
	Rules map[string]string
	Backends map[string]string
}

func (b *Balancer) AddRule (Rule) {

}