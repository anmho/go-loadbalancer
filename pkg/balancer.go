package balancer

import "fmt"

type Rule struct {
	ttlMS int
	rate int
}

func NewRule(ttlMS int, rate int) Rule {
	return Rule{ttlMS: ttlMS, rate: rate}
}

func (r*Rule) String() string {
	return fmt.Sprintf("Rule: %d %d", r.ttlMS, r.rate)
}

type Balancer struct {
	Rules map[string]string
	Backends map[string]string
}



func (b *Balancer) AddRule (Rule) {

}