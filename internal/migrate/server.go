package migrate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func ReadJsonFileTitle(c echo.Context) error {
	jsonFile, err := os.Open("data/titles.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var titles Titles

	json.Unmarshal(byteValue, &titles)

	// var ageCertification []string
	// for _, title := range titles {
	// 	if !contains(ageCertification, title.AgeCertification) {
	// 		ageCertification = append(ageCertification, title.AgeCertification)
	// 	}

	// }

	// var s string

	// for _, i := range ageCertification {
	// if i != ""{
	// 	s += fmt.Sprintf("INSERT INTO certification (name) VALUES ('%s');\n", i)
	// }
	// }

	// var genres []string

	// for _, i := range titles {
	// 	f := strings.Split(i.ProductionCountries, "'")
	// 	for _, j := range f {
	// 		if !contains(genres, j) {
	// 			if j != "[]" && j != "]" && j != "[" && j != ", " && j != "" {
	// 				genres = append(genres, j)
	// 			}
	// 		}

	// 	}
	// }

	// for _, i := range genres {
	// 	s += fmt.Sprintf("INSERT INTO genre (name) VALUES ('%s');\n", i)
	// }

	return c.JSON(http.StatusOK, titles)

}

func ReadJsonFileCredit(c echo.Context) error {
	jsonFile, err := os.Open("data/credits.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var credits Credits

	json.Unmarshal(byteValue, &credits)

	return c.JSON(http.StatusOK, credits)

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
