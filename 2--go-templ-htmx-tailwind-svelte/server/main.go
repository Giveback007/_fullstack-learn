// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/a-h/templ"
// )

// func main() {
// 	component := hello("John!")

// 	http.Handle("/", templ.Handler(component))

// 	fmt.Println("Listening on http://localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Listening on http://localhost:3000/")
}
