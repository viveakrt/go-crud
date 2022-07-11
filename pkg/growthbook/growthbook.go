package growthbook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myapp/config"
	"net/http"

	growthbook "github.com/growthbook/growthbook-golang"
)

type GrowthBookApiResp struct {
	Features json.RawMessage
	Status   int
}

func GetFeatureMap() []byte {
	var apiResp GrowthBookApiResp
	resp, err := http.Get(config.Config.GrowthBookURL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &apiResp)

	if err != nil {
		fmt.Println(err)
	}

	return apiResp.Features
}

func Gbook(id string) *growthbook.GrowthBook {

	featureMap := GetFeatureMap()
	features := growthbook.ParseFeatureMap(featureMap)

	context := growthbook.NewContext().
		WithFeatures(features).
		// TODO: Real user attributes
		WithAttributes(growthbook.Attributes{
			"id":       id,
			"deviceId": "foo",
			"company":  "foo",
			"loggedIn": true,
			"employee": true,
			"country":  "foo",
			"browser":  "foo",
			"url":      "foo",
		}).
		// TODO: Track in your analytics system
		WithTrackingCallback(func(experiment *growthbook.Experiment, result *growthbook.ExperimentResult) {

			log.Println(fmt.Sprintf("Experiment: %s, Variation: %d", experiment.Key, result.VariationID))
		})

	gb := growthbook.New(context)
	return gb
}
