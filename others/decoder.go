package gotest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//type (
//	gResult struct {
//		GsearchResultClass  string `json:"GsearchResultClass"`
//		UnescapedURL        string `json:"unescapedUrl"`
//		URL                 string `json:"url"`
//		VisibleURL          string `json:"visibleUrl"`
//		CacheURL            string `json:"cacheUrl"`
//		Title               string `json:"title"`
//		TitleNoFormatting   string `json:"titleNoFormatting"`
//		Content             string `json:"content"`
//	}
//
//	gResponse struct {
//		ResponseData struct{
//			Results  []gResult `json:"results"`
//		}  `json:"responseData"`
//	}
//)

type gResponse struct {
	ResponseData    string `json:"responseData"`
	ResponseDetails string `json:"responseDetails"`
	ResponseStatus  int    `json:"responseStatus"`
}

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	defer resp.Body.Close()

	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(gr)

}
