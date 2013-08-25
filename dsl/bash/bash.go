package bash

import (
	"fmt"
	"strconv"
	"morr.cc/nutsh.git/dsl"
)

func QueryReturn(query string) int {
	dsl.Query(query)
	output := dsl.Query("echo $?")
	value, _ := strconv.Atoi(output[0:len(output)-2])
	return value
}

func Test(expression string) bool {
	return QueryReturn("[[ "+expression+" ]]") == 0
}

func Execute(command string) {
	output := dsl.Query(command)
	valuestring := dsl.Query("echo $?")
	value, _ := strconv.Atoi(valuestring[0:len(valuestring)-2])
	if value != 0 {
		panic(fmt.Sprintf("executing `%s` failed: %s", command, output))
	}
}
