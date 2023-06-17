package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Document struct {
	DocumentID   int    `json:"documentID"`
	DocumentName string `json:"documentName"`
	DocumentPath string `json:"documentPath"`
}

var router = gin.Default()

func main() {
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Welcome!": "This is the Giggling Giraffe API",
		})
	})

	router.GET("/api/documents", getDocuments)
	router.POST("/api/documents", writeDocuments)

	router.Run()

}

func getDocuments(c *gin.Context) {

	data, err := os.ReadFile("data/documents.json")
	if err != nil {
		log.Fatal(err)
	}

	var documents []Document

	err = json.Unmarshal(data, &documents)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, documents)

}

func writeDocuments(c *gin.Context) {
	var documents []Document

	err := c.BindJSON(&documents)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(documents, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./documents.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Documents written successfully",
	})
}
