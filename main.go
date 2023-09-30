package main

import (
	"fmt"
	"limite/models"
	"strings"
	"golang.org/x/text/unicode/norm"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"unicode"

)

type allsentences []models.Sentence

var sentences = allsentences{
	{
		Words:       "Hello",
		Limit:       0.5,
		Ideal_limit: 0.5,
	},
}

func quitarAcentos(cadena string) string {
	// Crea una forma de normalización NFD
	forma := norm.NFD

	// Aplica la normalización para descomponer los caracteres
	cadenaNormalizada := forma.String(cadena)

	// Elimina los caracteres de marca diacrítica
	cadenaSinAcentos := strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			// Mantener letras mayúsculas sin cambios
			return r
		}
		if r >= 'a' && r <= 'z' {
			// Mantener letras minúsculas sin cambios
			return r
		}
		if unicode.Is(unicode.Mn, r) {
			// Eliminar marcas diacríticas
			return -1
		}
		return r
	}, cadenaNormalizada)

	return cadenaSinAcentos
}

func contarLetras(cadena string) map[rune]int {
	resultado := make(map[rune]int)

	

	// Convierte la cadena a minúsculas para que no distinga entre mayúsculas y minúsculas
	cadena = strings.ToLower(cadena)

	for _, char := range cadena {
		if 'a' <= char && char <= 'z' {
			resultado[char]++
		}
	}

	return resultado
}
func dividirConteo(resultados map[rune]int, divisor float64) map[rune]float64 {
	resultadoDividido := make(map[rune]float64)

	// Itera sobre los resultados y divide el conteo por el divisor
	for letra, count := range resultados {
		resultadoDividido[letra] = float64(count) / float64(divisor)
	}

	return resultadoDividido
}
func sumarValores(resultadosDivididos map[rune]float64) float64 {
	suma := 0.0

	// Itera sobre los valores y suma
	for _, valor := range resultadosDivididos {
		suma += valor
	}

	return suma
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./client/dist")
	app.Get("/sentences", func(c *fiber.Ctx) error {

		return c.JSON(sentences)
	})

	app.Post("/sentences", func(c *fiber.Ctx) error {
		newSentence := new(models.Sentence)

		if err := c.BodyParser(newSentence); err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(200)
		}
		c.BodyParser(&newSentence)
		wordsWhtoutSpace := strings.ReplaceAll(newSentence.Words, " ", "")
		fmt.Println(wordsWhtoutSpace)
        cnew:=quitarAcentos(wordsWhtoutSpace)
		fmt.Println(cnew)
		newSentence.Limit = float64(len(cnew))
		fmt.Println(newSentence.Limit)

		resultado := contarLetras(cnew)

		for letra, count := range resultado {
			fmt.Printf("'%c': %d\n", letra, count)
		}

		resultadoDividido := dividirConteo(resultado, newSentence.Limit)
		for letra, count := range resultadoDividido {
			fmt.Printf("'%c': %.16f\n", letra, count)
		}

		suma := sumarValores(resultadoDividido)

		fmt.Printf("Suma de valores divididos: %.16f\n", suma)

		newSentence.Ideal_limit = suma

		return c.JSON(newSentence)
	})
	app.Listen(":3000")

}
