package modify

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type DataFile struct {
	IP   string `json:"ip"`
	PATH string `json:"path"`
}

func GetIPConfig(arrCmd []string) []DataFile {
	var dataFile DataFile
	var arrDataFile []DataFile
	for i := 0; i < len(arrCmd); i++ {
		data, err := ioutil.ReadFile(arrCmd[i])
		if err != nil {
			fmt.Println(err)
		}
		sRead := string(data)
		removeDistance := strings.ReplaceAll(sRead, " ", "")
		removeLine := strings.ReplaceAll(removeDistance, "\n", "")
		result2 := strings.Index(removeLine, "',REDIS_PORT")
		result1 := strings.Index(removeLine, `REDIS_HOST`)
		ipRedisOfConfig := removeLine[result1+12 : result2]
		dataFile.IP = ipRedisOfConfig
		dataFile.PATH = arrCmd[i]
		//dataFile.dataRead = string(data)
		fmt.Println("IP Of File Config.js", dataFile)
		arrDataFile = append(arrDataFile, dataFile)
		fmt.Println("IP Of File Config.js", ipRedisOfConfig)
	}
	return arrDataFile
}

func ChangeConfig(dataFind []DataFile, addr string) {
	for i := 0; i < len(dataFind); i++ {
		data, err := ioutil.ReadFile(dataFind[i].PATH)
		if err != nil {
			fmt.Println(err)
		}
		//REDIS_HOST: '10.22.7.107',
		f, err := os.OpenFile(dataFind[i].PATH, os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		re := regexp.MustCompile(`REDIS_HOST: ` + `'` + dataFind[i].IP + `'`)
		if _, err = f.WriteString(
			re.ReplaceAllString(string(data), addr)); err != nil {
			panic(err)
		}
	}
}
