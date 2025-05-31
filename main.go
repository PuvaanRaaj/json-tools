package main

import (
    "bytes"
    "encoding/json"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Serve static assets like CSS/JS/images (if any)
    router.Static("/static", "./static")

    // Serve HTML views
    router.GET("/beautify", func(c *gin.Context) {
        c.File("./static/beautify.html")
    })
    router.GET("/minify", func(c *gin.Context) {
        c.File("./static/minify.html")
    })


    // Beautify endpoint
    router.POST("/beautify", func(c *gin.Context) {
        var input struct {
            Raw string `json:"raw"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        var buf bytes.Buffer
        err := json.Indent(&buf, []byte(input.Raw), "", "  ")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"beautified": buf.String()})
    })

    // Minify endpoint
    router.POST("/minify", func(c *gin.Context) {
        var input struct {
            Raw string `json:"raw"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        var out bytes.Buffer
        err := json.Compact(&out, []byte(input.Raw))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"minified": out.String()})
    })

    // Optional: redirect "/" to /beautify
    router.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/beautify")
    })

    router.Run(":8080")
}
