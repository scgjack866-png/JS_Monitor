package vo

type PrometheusReq struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

type PrometheusQueryReq struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

type Result struct {
	Metric Metric `json:"metric"`
	Value  []any  `json:"value"`
}

type Metric struct {
	__Name__      string `json:"__name__"`
	AgentHostname string `json:"agent_hostname"`
	Domain        string `json:"domain"`
	Target        string `json:"target"`
	Method        string `json:"method"`
}
