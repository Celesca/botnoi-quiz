package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Pokemon struct {
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Name    string `json:"name"`
	Sprites struct {
		BackDefault      *string `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        *string `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     *string `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       *string `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
	}
}

func fetch(url string, v interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func main() {
	e := echo.New()
	e.POST("/pokemon/:id", func(c echo.Context) error {
		id := c.Param("id")
		var pokemon Pokemon
		err := fetch("https://pokeapi.co/api/v2/pokemon/"+id, &pokemon)
		if err != nil {
			log.Fatal(err)
		}
		err = fetch("https://pokeapi.co/api/v2/pokemon-form/"+id, &pokemon)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, pokemon)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
