package silence

import (
	"OperationAndMonitoring/initialize"
	"github.com/parnurzeal/gorequest"
	"time"
)

var (
	request = *gorequest.New()
)

func CreateSilences(ip string, endsTime time.Time) (string, bool) {
	var startsAt = time.Now().UTC().Format(time.RFC3339)
	endsAt := endsTime.UTC().Format(time.RFC3339)
	res, body, _ := request.
		Post(initialize.Grafana.ApiUrl+"/api/alertmanager/grafana/api/v2/silences").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(`{
	 "startsAt": "` + startsAt + `",
	 "endsAt": "` + endsAt + `",
	 "comment": "created ` + startsAt + `",
	 "createdBy": "admin",
	 "matchers": [
	   {
	     "name": "IP",
	     "value": "` + ip + `",
	     "isEqual": true,
	     "isRegex": false
	   }
	 ]
	}`).
		End()
	if res.StatusCode == 202 {
		return body, true
	}
	return body, false
}

func DeleteSilences(silenceUID string) {
	request.
		Delete(initialize.Grafana.ApiUrl+"/api/alertmanager/grafana/api/v2/silence/"+silenceUID).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		End()
}
