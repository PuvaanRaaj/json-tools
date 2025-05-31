package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // fallback for local testing
    }

    router := gin.Default()

    router.Static("/static", "./static")
    router.GET("/beautify", func(c *gin.Context) {
        c.File("./static/beautify.html")
    })
    router.GET("/minify", func(c *gin.Context) {
        c.File("./static/minify.html")
    })

    router.POST("/beautify", func(c *gin.Context) {
        var input struct {
            Raw string `json:"raw"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        var buf bytes.Buffer
        if err := json.Indent(&buf, []byte(input.Raw), "", "  "); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"beautified": buf.String()})
    })

    router.POST("/minify", func(c *gin.Context) {
        var input struct {
            Raw string `json:"raw"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        var out bytes.Buffer
        if err := json.Compact(&out, []byte(input.Raw)); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"minified": out.String()})
    })

    log.Println("Listening on port", port)
    router.Run(":" + port)
}
