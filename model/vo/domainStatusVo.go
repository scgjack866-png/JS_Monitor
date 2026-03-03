package vo

type DomainStatusVo struct {
	Target string  `json:"target"`
	Value  []Value `json:"value"`
}
type Describe struct {
	AgentHostname string `json:"agent_hostname"`
}
type Value struct {
	Describe Describe      `json:"describe"`
	Values   []interface{} `json:"values"`
}
