package char_one

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	c.XML(http.StatusOK, Person{FirstName: "Mohamed", LastName: "Labouardy"})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	log.Fatal(router.Run())
}
