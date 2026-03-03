package grafana

import (
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"
)

func CmdWiterFile() {
	//需要执行的命令： free -mh
	var timeT = time.Now().Unix()
	cmd := exec.Command("/bin/bash", "-c", `echo -e "{\n  \"version\": \"`+strconv.FormatInt(timeT, 10)+`\",\n  \"configs\": {\n    \"http_response\": {\n    \"`+strconv.FormatInt(timeT, 10)+`\":{\n    \"config\": \"## collect interval\\\n#interval = 15\\\n\\\n[[instances]]\\\ntargets = [\\\n    `+cmdExec()+`\\\n]\",\n      \"format\": \"toml\"\n    }\n    }\n  }\n}" > `+initialize.Grafana.ConfigPath)

	// 获取管道输入
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return
	}

	// 执行Linux命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return
	}

	// 读取所有输出
	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return
	}

	fmt.Printf("打印内存信息：\n\n%s", bytes)

}

func cmdExec() string {
	var domains []entity.Domain
	var whereOrder []mysql.PageWhereOrder
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	err := utils.Find(&entity.Domain{}, &domains, whereOrder...)
	if err != nil {
		return ``
	}
	//var ceshiStr = [4]string{"http://ceshi.com", "https://ceshi1.com", "https://ceshi2.com", "https://ceshi3.com"}

	var dashboards []byte
	for i, domain := range domains {
		dashboards = append(dashboards, []byte(`\\\"`+domain.Domain+`\\\"`)...)
		if i != len(domains)-1 {
			dashboards = append(dashboards, []byte(",\\\\\\n   ")...)
		}
	}
	return string(dashboards)
}
