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

func main() {
	e := echo.New()
	e.POST("/pokemon/:id", func(c echo.Context) error {
		id := c.Param("id")
		response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + id)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var pokemon Pokemon
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			log.Fatal(err)
		}

		second_response, err := http.Get("https://pokeapi.co/api/v2/pokemon-form/" + id)
		if err != nil {
			log.Fatal(err)
		}

		second_body, err := io.ReadAll(second_response.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(second_body, &pokemon)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, pokemon)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
