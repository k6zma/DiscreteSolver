package mathalgos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type VennDiagramBuilder struct {
	apiKey string
}

func NewVennDiagramBuilder(apiKey string) *VennDiagramBuilder {
	return &VennDiagramBuilder{
		apiKey: apiKey,
	}
}

func (v *VennDiagramBuilder) BuildDiagram(expression string) (string, error) {
	vennQuery := fmt.Sprintf("Venn diagram of %s", expression)
	encodedQuery := url.QueryEscape(vennQuery)
	queryURL := fmt.Sprintf("http://api.wolframalpha.com/v2/query?input=%s&format=image&output=JSON&appid=%s", encodedQuery, v.apiKey)

	resp, err := http.Get(queryURL)
	if err != nil {
		return "", fmt.Errorf("error with request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error with request: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var data struct {
		QueryResult struct {
			Pods []struct {
				Title   string `json:"title"`
				Subpods []struct {
					Img struct {
						Src string `json:"src"`
					} `json:"img"`
				} `json:"subpods"`
			} `json:"pods"`
		} `json:"queryresult"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	for _, pod := range data.QueryResult.Pods {
		if pod.Title == "Venn diagram" {
			return pod.Subpods[0].Img.Src, nil
		}
	}

	return "", fmt.Errorf("error while collecting an image")
}
