package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"encoding/json"
    "fmt"
    "io/ioutil"
)
type apiResourceList struct {
	Kind         string `json:"kind"`
	APIVersion   string `json:"apiVersion"`
	GroupVersion string `json:"groupVersion"`
	Resources    []struct {
		Name               string   `json:"name"`
		SingularName       string   `json:"singularName"`
		Namespaced         bool     `json:"namespaced"`
		Kind               string   `json:"kind"`
		Verbs              []string `json:"verbs"`
		ShortNames         []string `json:"shortNames,omitempty"`
		StorageVersionHash string   `json:"storageVersionHash,omitempty"`
	} `json:"resources"`
}
type apiGroupList struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Groups     []struct {
		Name     string `json:"name"`
		Versions []struct {
			GroupVersion string `json:"groupVersion"`
			Version      string `json:"version"`
		} `json:"versions"`
		PreferredVersion struct {
			GroupVersion string `json:"groupVersion"`
			Version      string `json:"version"`
		} `json:"preferredVersion"`
	} `json:"groups"`
}
// Basic API Handler
func ApiResourceList(c echo.Context) error {
	jsonFile, err := ioutil.ReadFile("/Users/zeerg/starfleet/helix-honeypot/kube_json/api_resourcelist.json")
    if err != nil {
      fmt.Print(err)
    }
	var data apiResourceList
	err = json.Unmarshal(jsonFile, &data)

	return c.JSON(http.StatusOK, data)
}
func ApiGroupList(c echo.Context) error {
	jsonFile, err := ioutil.ReadFile("/Users/zeerg/starfleet/helix-honeypot/kube_json/api_grouplist.json")
    if err != nil {
      fmt.Print(err)
    }
	var data apiGroupList
	err = json.Unmarshal(jsonFile, &data)

	return c.JSON(http.StatusOK, data)
}