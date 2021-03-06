package readchangefile

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

func ChangeFile(file string){

	fmt.Printf("Reading file %d\n", file)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil{
		log.Fatal(err)
	}

	fileStr := string(data)

	fileSpl := strings.Split(fileStr, "\n")

	fmt.Printf("Changing file")

	for i := range fileSpl {

		lineSpl := strings.Split(fileSpl[i], ",")
		if len(lineSpl) < 4 {
			continue
		}
		lineSpl[4] = "0"
		lineSpl[5] = "0"
		lineSpl[8] = "0"
		fileSpl[i] = strings.Join(lineSpl, ",")
	}

	fileStr = strings.Join(fileSpl, "\n")

	ioutil.WriteFile(file, []byte(fileStr), 0644)

	fmt.Printf("File changed and saved succesfully!")
}
