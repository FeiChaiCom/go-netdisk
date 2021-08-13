package mysqlclient

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDataValue(t *testing.T) {
	mySql, err := NewMySql()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer mySql.Close()

	err = mySql.Open("127.0.0.1", "root", "", "configcenter_ied", 3306, 3000, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	myRows, err := mySql.Query("select * from cc_ApplicationBase limit 10")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range myRows {
		jsonString, _ := json.Marshal(v)
		fmt.Println(string(jsonString))
	}
	type List []interface{}
	list := make(List, 7)
	list[0] = "1"
	list[1] = "aa"
	list[2] = "pod"
	list[3] = "bb"
	list[4] = "service"
	list[5] = "2016-11-16 16:39:47"
	list[6] = "2016-11-16 16:39:47"
	res, err := mySql.Insert("INSERT INTO `configcenter_ied`.`cc_InstanceInclude` (`ID`, `ApplicationID`, `InstanceName`, `InstanceType`, `ConfigName`, `ConfigType`, `CreateTime`, `LastTime`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?)", list)
	fmt.Println(res)
}
