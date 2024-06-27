package main

import (
	"fmt"
	"full-stack/server/views"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func MakeNames(names []string) []views.Name {
	nameArr := []views.Name{}
	for _, n := range names {
		nameArr = append(nameArr, views.Name{Name: n})
	}

	return nameArr
}

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {
	count := 0
	r := gin.Default()
	r.Static("css", "./server/css")

	r.GET("/", func(c *gin.Context) {
		names := MakeNames([]string{"dovy", "karen", "angel"})

		render(c, http.StatusOK, views.MakeHomePage(names, strconv.Itoa(count)))
	})

	r.POST("/count-up", func(c *gin.Context) {
		count++
		c.Status(200)
		c.String(http.StatusOK, strconv.Itoa(count))
		// ???
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	/* LOCALHOST */
	fmt.Println("Listening on http://localhost:3000/")
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
