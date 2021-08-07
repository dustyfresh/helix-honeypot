package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"encoding/json"
    "fmt"
    "io/ioutil"
)
type podsStruct struct {
	APIVersion string        `json:"apiVersion"`
	Items      []interface{} `json:"items"`
	Kind       string        `json:"kind"`
	Metadata   struct {
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}
//Pods Handler
func PodsHandler(c echo.Context) error {
	jsonFile, err := ioutil.ReadFile("/Users/zeerg/starfleet/helix-honeypot/kube_json/default_pods.json")
    if err != nil {
      fmt.Print(err)
    }
	var data podsStruct
	err = json.Unmarshal(jsonFile, &data)

	return c.JSON(http.StatusOK, data)
}
