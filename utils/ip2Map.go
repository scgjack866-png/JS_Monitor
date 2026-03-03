package utils

import (
	"OperationAndMonitoring/utils/convert"
	"strings"
)

func String2Map(ipStr string) ([]string, []string, bool) {
	var ss_jian []string
	var ss_xiang []string
	// 去除所有空格
	noSpaceStr := strings.Replace(ipStr, " ", "", -1)

	// 1.1.1.1-10/20/30,2.2.2.2-10/30
	noCommaStr := strings.Split(noSpaceStr, ",")
	// 1.1.1.1-10/20/30     2.2.2.2-10/30
	for _, differentIPStr := range noCommaStr {

		if len(strings.Split(differentIPStr, ".")) != 4 {
			return nil, nil, false
		}
		if strings.Contains(differentIPStr, "/") {

			// 1 1 1 1-10/20/30
			eachIpMap := strings.Split(differentIPStr, ".")
			// 1-10 20 30
			dSegmentIpMap := strings.Split(eachIpMap[3], "/")

			for _, eachDSegmentIpStr := range dSegmentIpMap {
				if strings.Contains(eachDSegmentIpStr, "-") {
					// 1 10
					oneDSegmentIpMap := strings.Split(eachDSegmentIpStr, "-")

					for i := convert.ToInt(oneDSegmentIpMap[0]); i <= convert.ToInt(oneDSegmentIpMap[1]); i++ {
						ss_xiang = append(ss_xiang, eachIpMap[0]+"."+eachIpMap[1]+"."+eachIpMap[2]+"."+convert.ToString(i))
					}
				} else {
					ss_xiang = append(ss_xiang, eachIpMap[0]+"."+eachIpMap[1]+"."+eachIpMap[2]+"."+eachDSegmentIpStr)
				}
				ss_jian = append(ss_jian, eachIpMap[0]+"."+eachIpMap[1]+"."+eachIpMap[2]+"."+eachDSegmentIpStr)
			}

		} else {
			eachIpMap := strings.Split(differentIPStr, ".")
			if strings.Contains(eachIpMap[3], "-") {

				oneDSegmentIpMap := strings.Split(eachIpMap[3], "-")

				for i := convert.ToInt(oneDSegmentIpMap[0]); i <= convert.ToInt(oneDSegmentIpMap[1]); i++ {
					ss_xiang = append(ss_xiang, eachIpMap[0]+"."+eachIpMap[1]+"."+eachIpMap[2]+"."+convert.ToString(i))
				}
			} else {
				ss_xiang = append(ss_xiang, eachIpMap[0]+"."+eachIpMap[1]+"."+eachIpMap[2]+"."+eachIpMap[3])
			}

			ss_jian = append(ss_jian, eachIpMap[0]+"."+eachIpMap[1]+"."+eachIpMap[2]+"."+eachIpMap[3])

		}

	}
	return ss_jian, ss_xiang, true
}

func map2String(ss []string) string {
	var s string

	return s
}
