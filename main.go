package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

var FILES = []string{"dist.male.first", "dist.female.first", "dist.all.last"}

// Get names based on a census distribution [1990]
func get_distribution_name(fiLines []string) string {
	var selected = rand.Float64() * 100

	for line := range fiLines {
		var lineRA []string
		lineRA = strings.Fields(fiLines[line])
		var cummulative, _ = strconv.ParseFloat(lineRA[2], 64)
		if cummulative > selected {
			return lineRA[0]
		}
		selected = rand.Float64() * 100
	}
	return ""
}
func get_name(fiLines []string) string {
	var selected = rand.Intn(len(fiLines))
	lineRA := strings.Fields(fiLines[selected])
	return lineRA[0]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maleFirsts, _ := readLines(FILES[0])
	lastNames, _ := readLines(FILES[2])
	// male names only
	for i := 0; i <= 5000; i++ {
		//fmt.Println(get_name(maleFirsts) + " " + get_name(lastNames))
		fmt.Println(get_distribution_name(maleFirsts) + " " + get_distribution_name(lastNames))
	}
}

/*
func get_first_name(gender=None){
    if gender is None:
        gender = random.choice(('male', 'female'))
    if gender not in ('male', 'female'):
        raise ValueError("Only 'male' and 'female' are supported as gender")
    return get_name(FILES['first:%s' % gender]).capitalize()
}

func get_last_name(){
    return get_name(FILES['last']).capitalize()
}


func get_full_name(gender=None){
    return "{0} {1}".format(get_first_name(gender), get_last_name())
}
*/
