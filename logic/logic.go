package logic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	SchoolList []*SchoolInfo
)

func InitLogic(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("open file %s failed, err is %v", filename, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var id int
	for {
		line, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			break
		}
		if errRet != nil {
			err = errRet
			log.Print(err)
			return
		}

		strSplit := strings.Split(line, "\t")
		if len(strSplit) != 4 {
			fmt.Printf("err read line %s\n", line)
			continue
		}

		var schoolInfo SchoolInfo
		id++
		schoolInfo.SchoolId = id
		schoolInfo.Provice = strings.TrimSpace(strSplit[0])
		schoolInfo.City = strings.TrimSpace(strSplit[1])
		schoolInfo.SchoolName = strings.TrimSpace(strSplit[2])
		st, err := strconv.Atoi(strings.TrimSpace(strSplit[3]))
		if err != nil {
			log.Printf("err nums :%v\n", err)
		}
		schoolInfo.SchoolType = st

		SchoolList = append(SchoolList, &schoolInfo)

		fmt.Printf("school info :%+v \n", schoolInfo)
	}
	return
}

func Search(keyword string, limit int) (schools []*SchoolInfo) {
	for _, s := range SchoolList {
		if strings.HasPrefix(s.SchoolName, keyword) == true {
			schools = append(schools, s)
			if len(schools) > limit {
				return
			}
		}
	}
	return
}
