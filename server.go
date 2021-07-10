package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Characters struct {
	totalAngka int `json:"totalAngka"`
}

type Faktorial struct {
	Nilai      int `json:"nilai" form:"nilai" query:"nilai"`
	TotalAngka int `json:"totalAngka" form:"totalAngka" query:"totalAngka"`
}

func main() {
	e := echo.New()
	e.GET("/faktor", func(c echo.Context) error {
		u := new(Faktorial)
		if err := c.Bind(u); err != nil {
			return err
		}
		fmt.Println(`Reading File ....`)
		result := findFactor(u.Nilai)
		fmt.Println(`Done ....`)
		return c.JSON(http.StatusOK, result)
	})

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}

func findFactor(total int) Faktorial {
	var faktorial Faktorial
	maximum := total
	counter := 0
	for i := 1; i <= maximum; i++ {
		faktor := 1
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				if faktor == 6 && j == i {
					// fmt.Println(`faktor`, faktor, i)
					counter = counter + 1
					faktor = 1
					break
				} else {
					faktor = faktor + 1
				}
			} else {
				continue
			}
		}
	}
	fmt.Println(`------------------------`, counter)
	faktorial.TotalAngka = counter
	return faktorial
}
