package main

import (
"fmt"
"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Tudo certo para começar!")
}

func goRoutes(c =gin.Engine) =gin.Engine {
	c .GET("/heart", routeHeart)

	return c 
}