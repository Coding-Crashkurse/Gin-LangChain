package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/generate", generateCompletion)
	}

	router.Run(":8080")
}


func generateCompletion(c *gin.Context) {
	var requestData struct {
		Prompt string `json:"prompt"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	llm, err := openai.New()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	completion, err := llm.Call(context.Background(), requestData.Prompt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"completion": completion})
}

