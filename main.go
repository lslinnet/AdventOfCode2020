package main

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lslinnet/adventofcode2020/day1"
)

type stubMapping map[string]interface{}

// StubStorage is a list of callable funcs
var StubStorage = stubMapping{}

func main() {
	StubStorage = map[string]interface{}{
		"1a": day1.Day1a,
		"1b": day1.Day1b,
	}

	r := gin.Default()

	r.POST("/day/:day", func(c *gin.Context) {
		day := c.Param("day")
		input := strings.Fields(c.PostForm("input"))
		var i2 = []int{}

		for _, i := range input {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			i2 = append(i2, j)
		}

		resA, _ := Call(day+"a", i2)
		resB, _ := Call(day+"b", i2)

		c.JSON(200, gin.H{
			"day" + day + "a_result": resA,
			"day" + day + "b_result": resB,
		})
	})

	r.Run()
}

// Call is a helper function to make it easier to call the stubbed functions.
func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is out of index")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLinesToIntArray(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var value int
		value, err = strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}
	return lines, scanner.Err()
}
