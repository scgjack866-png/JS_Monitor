package vo

type FlowVO struct {
	Ip      string  `json:"ip"`
	Out     float64 `json:"out"`
	In      float64 `json:"in"`
	OutUnit string  `json:"out_unit"`
	InUnit  string  `json:"in_unit"`
}

type SnapshotVO struct {
	Url string `json:"url"`
}

type DsQueryFlowBody struct {
	Results struct {
		In struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					Name  string `json:"name"`
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							AgentHostname string `json:"agent_hostname"`
							Interface     string `json:"interface"`
						} `json:"labels,omitempty"`
						Config0 struct {
							DisplayNameFromDS string `json:"displayNameFromDS"`
						} `json:"config,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"IN"`
		Out struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					Name  string `json:"name"`
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							AgentHostname string `json:"agent_hostname"`
							Interface     string `json:"interface"`
						} `json:"labels,omitempty"`
						Config0 struct {
							DisplayNameFromDS string `json:"displayNameFromDS"`
						} `json:"config,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"OUT"`
	} `json:"results"`
}

type DsQueryCpuBody struct {
	Results struct {
		CPU struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							Name          string `json:"__name__"`
							AgentHostname string `json:"agent_hostname"`
							CPU           string `json:"cpu"`
						} `json:"labels,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"Cpu"`
	} `json:"results"`
}

type DsQueryFreeBody struct {
	Results struct {
		Free struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							Name          string `json:"__name__"`
							AgentHostname string `json:"agent_hostname"`
						} `json:"labels,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"Free"`
	} `json:"results"`
}

type DsQueryLoadBody struct {
	Results struct {
		LoadOne struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							Name          string `json:"__name__"`
							AgentHostname string `json:"agent_hostname"`
						} `json:"labels,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"LoadOne"`
		LoadFive struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							Name          string `json:"__name__"`
							AgentHostname string `json:"agent_hostname"`
						} `json:"labels,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"LoadFive"`
		LoadFifteen struct {
			Status int `json:"status"`
			Frames []struct {
				Schema struct {
					RefID string `json:"refId"`
					Meta  struct {
						Type        string `json:"type"`
						TypeVersion []int  `json:"typeVersion"`
						Custom      struct {
							ResultType string `json:"resultType"`
						} `json:"custom"`
						ExecutedQueryString string `json:"executedQueryString"`
					} `json:"meta"`
					Fields []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						TypeInfo struct {
							Frame string `json:"frame"`
						} `json:"typeInfo"`
						Config struct {
							Interval int `json:"interval"`
						} `json:"config,omitempty"`
						Labels struct {
							Name          string `json:"__name__"`
							AgentHostname string `json:"agent_hostname"`
						} `json:"labels,omitempty"`
					} `json:"fields"`
				} `json:"schema"`
				Data struct {
					Values [][]float64 `json:"values"`
				} `json:"data"`
			} `json:"frames"`
		} `json:"LoadFifteen"`
	} `json:"results"`
}
