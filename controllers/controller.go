package controllers

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func UrlScrapper(c *gin.Context) {
	var requestUrls []string
	if err := c.ShouldBind(&requestUrls); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, url := range requestUrls {
		wg.Add(1)
		go FetchUrl(url, c)
	}
	wg.Wait()

}

func FetchUrl(url string, c *gin.Context) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't get url", "error": err.Error()})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't read url body", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "fetched content of url", url: body})
}
