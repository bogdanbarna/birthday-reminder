package controller

import (
	"bogdanbarna/hello-rest-gin/core"
	"bogdanbarna/hello-rest-gin/model"
	"bogdanbarna/hello-rest-gin/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context) {
	log.Println("Getting all persons")
	persons, err := repository.FindPersons()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"persons": persons})
}

func GetPerson(c *gin.Context) {
	username := c.Param("username")
	p, err := repository.FindPerson(username)
	log.Println(err)
	if p != (model.Person{}) && err == nil {
		log.Println("Getting person ", p.Username)
		c.IndentedJSON(http.StatusOK, gin.H{"person": p})
	} else {
		log.Println("User not found", username)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

func GetPersonBirthday(c *gin.Context) {
	username := c.Param("username")
	p, err := repository.FindPerson(username)
	log.Println(err)
	if p != (model.Person{}) && err == nil {
		log.Println("Getting birthday for person ", p.Username)
		message := core.ComposeMessage(p.Username, p.Birthday)
		c.IndentedJSON(http.StatusOK, gin.H{"message": message})
	} else {
		log.Println("User not found", username)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

func PutUser(c *gin.Context) {
	var newPerson model.Person
	// using BindJson method to serialize body with struct
	if err := c.BindJSON(&newPerson); err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	newPerson.Username = c.Param("username")
	p, err := repository.FindPerson(newPerson.Username)
	if p != (model.Person{}) && err == nil {
		log.Println("Updating existing person with new birthday")
		repository.UpdatePerson(p, newPerson.Birthday)
		c.IndentedJSON(http.StatusNoContent, gin.H{})
	} else {
		log.Println("Adding person", newPerson)
		repository.CreatePerson(newPerson)
		c.IndentedJSON(http.StatusCreated, gin.H{})
	}
}

func PostUser(c *gin.Context) {
	var newPerson model.Person
	// using BindJson method to serialize body with struct
	if err := c.BindJSON(&newPerson); err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	p, err := repository.FindPerson(newPerson.Username)
	if p != (model.Person{}) && err == nil {
		log.Println("Person already exists", newPerson.Username)
	} else {
		log.Println("Adding person", newPerson)
		repository.CreatePerson(newPerson)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func PatchUser(c *gin.Context) {
	log.Println("Not implemented")
	c.IndentedJSON(http.StatusNotImplemented, gin.H{})
}

func DeleteUser(c *gin.Context) {
	var newPerson model.Person
	newPerson.Username = c.Param("username")

	p, err := repository.FindPerson(newPerson.Username)
	if p != (model.Person{}) && err == nil {
		log.Println("Deleting person ", p.Username)
		repository.SoftDeletePerson(p)
		c.IndentedJSON(http.StatusOK, gin.H{})
	} else {
		log.Println("User not found", newPerson.Username)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}
