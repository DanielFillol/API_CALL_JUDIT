package request

import (
	"API_CALL_JUDIT/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func APIRequest(url string, auth string, request models.ReadCsv, duration time.Duration) (models.ResponseToCSV, error) {
	// Create a new BodyRequest struct with the document ID and pagination settings for the initial API call.
	search := models.Search{
		SearchType: "cpf",
		SearchKey:  fixDocument(request.Document),
	}
	req := models.BodyRequest{Search: search}

	// Serialize the BodyRequest struct to JSON.
	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return models.ResponseToCSV{}, err
	}

	// Create a new buffer with the JSON-encoded request body.
	reqBody := bytes.NewBuffer(jsonReq)

	// Make the API call and get the response.
	res, err := call(url, "POST", auth, reqBody, req)
	if err != nil {
		log.Println(err)
		return models.ResponseToCSV{}, errors.New(err.Error() + "  " + req.Search.SearchKey)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return models.ResponseToCSV{}, err
	}

	var response models.FirstResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return models.ResponseToCSV{}, err
	}

	start := time.Now()
	var totalTime time.Duration
	if response.Status != "completed" {
		for {
			time.Sleep(100 * time.Millisecond)
			res, err = call(url+"/"+response.RequestId, "GET", auth, reqBody, req)
			if err != nil {
				log.Println(err)
				return models.ResponseToCSV{}, errors.New(err.Error() + "  " + req.Search.SearchKey)
			}

			// Read the response body.
			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
				return models.ResponseToCSV{}, err
			}

			err = json.Unmarshal(body, &response)
			if err != nil {
				log.Println(err)
				return models.ResponseToCSV{}, err
			}

			if response.Status == "completed" {
				totalTime = time.Since(start)
				break
			}
		}
	}

	apiData, err := apiRequest("https://requests.prod.judit.io/responses?page=1&request_id="+response.RequestId, "GET", auth, request, duration, response.RequestId)
	if err != nil {
		return models.ResponseToCSV{}, err
	}

	return models.ResponseToCSV{
		Document:  req.Search.SearchKey,
		TotalTime: totalTime,
		R:         apiData,
	}, nil
}

// apiRequest makes an API request to the specified URL using the specified HTTP method and authentication header.
// It returns a models.ResponseJudit struct containing the API response body and an error (if any).
func apiRequest(url string, method string, auth string, request models.ReadCsv, duration time.Duration, requestId string) (models.ResponseJudit, error) {
	// Set a coll down for every request
	time.Sleep(duration)

	// Create a new BodyRequest struct with the document ID and pagination settings for the initial API call.
	search := models.Search{
		SearchType: "cpf",
		SearchKey:  fixDocument(request.Document),
	}

	req := models.BodyRequest{Search: search}

	// Serialize the BodyRequest struct to JSON.
	jsonReq, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return models.ResponseJudit{}, err
	}

	// Create a new buffer with the JSON-encoded request body.
	reqBody := bytes.NewBuffer(jsonReq)

	// Make the API call and get the response.
	res, err := call(url, method, auth, reqBody, req)
	if err != nil {
		log.Println(err)
		return models.ResponseJudit{}, errors.New(err.Error() + "  " + req.Search.SearchKey)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return models.ResponseJudit{}, err
	}

	// Unmarshal the response body into a ResponseJudit struct.
	var response models.ResponseJudit
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("4")
		log.Println(err)
		return models.ResponseJudit{}, err
	}

	//If the API response has more pages of data, make additional API calls and append the results to the response.
	if response.PageCount > 1 {
		lawsuits, err := callNextPage(method, auth, req.Search.SearchKey, requestId)
		if err != nil {
			log.Println(err)
			//return models.ResponseJudit{}, err
		}

		response.PageData = append(response.PageData, lawsuits...)

		return models.ResponseJudit{
			Page:          response.Page,
			PageData:      response.PageData,
			PageCount:     response.PageCount,
			AllCount:      response.AllCount,
			AllPagesCount: response.AllPagesCount,
		}, nil

	}

	return models.ResponseJudit{
		Page:          response.Page,
		PageData:      response.PageData,
		PageCount:     response.PageCount,
		AllCount:      response.AllCount,
		AllPagesCount: response.AllPagesCount,
	}, nil
}

// callNextPage returns a slice of models.Lawsuit structs containing the data from all pages of the API response.
func callNextPage(method string, auth string, req string, requestId string) ([]models.PageDataStr, error) {
	var lawsuits []models.PageDataStr
	page := 2
	for {
		// The API often can't handle to many next-page requests
		time.Sleep(100 * time.Millisecond)

		// Create a new BodyRequest struct with the document ID and updated pagination settings for the next API call.
		search := models.Search{SearchType: "cpf", SearchKey: req}
		request := models.BodyRequest{Search: search}

		// Serialize the BodyRequest struct to JSON.
		jsonReq, err := json.Marshal(request)
		if err != nil {
			log.Println(err)
			return lawsuits, err
		}

		// Create a new buffer with the JSON-encoded request body.
		reqBody := bytes.NewBuffer(jsonReq)

		// Call the API using the provided url, method, authorization, and request body.
		res, err := call("https://requests.prod.judit.io/responses"+"?page="+strconv.Itoa(page)+"&request_id="+requestId, method, auth, reqBody, request)
		if err != nil {
			log.Println(err)
			return lawsuits, errors.New(err.Error() + "  " + request.Search.SearchKey)
		}

		// Read the response body.
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return lawsuits, err
		}

		// Unmarshal the response body into a models.ResponseBody struct.
		var response models.ResponseJudit
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Println(err)
			return lawsuits, err
		}

		// Append the current response to the lawsuits slice.
		lawsuits = append(lawsuits, response.PageData...)

		// If the API response indicates there are no more pages, break out of the loop.
		var responsePage int
		if response.Page == "" {
			responsePage = 0
		} else {
			responsePage, err = strconv.Atoi(response.Page)
			if err != nil {
				log.Println(err)
				return lawsuits, err
			}
		}

		if responsePage == response.PageCount || len(response.PageData) == 0 {
			break
		}

		// Update the cursor for the next API call.
		page++
	}

	return lawsuits, nil
}

// call sends an HTTP request to the specified URL using the specified method and request body, with the specified authorization header.
// It returns the HTTP response or an error if the request fails.
func call(url, method string, AUTH string, body io.Reader, request models.BodyRequest) (*http.Response, error) {
	// Create an HTTP client with a 10-second timeout.
	client := &http.Client{Timeout: time.Second * 10}

	// Create a new HTTP request with the specified method, URL, and request body.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Set the Content-Type and Authorization headers for the request.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", AUTH)

	// Send the request and get the response.
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// If the response status code is not OK, return an error with the status code.
	if response.StatusCode == http.StatusCreated {
		log.Println("status", "Waiting", "HTTP:", strconv.Itoa(response.StatusCode), "document:", request.Search.SearchKey, "url:", url, "request:", request)
	} else if response.StatusCode == http.StatusOK {
		if strings.Contains(url, "page") {
			log.Println("status", "OK", "HTTP:", strconv.Itoa(response.StatusCode), "document:", request.Search.SearchKey, "url:", url, "request:", request)
		} else {
			log.Println("status:", "Waiting", "HTTP:", strconv.Itoa(response.StatusCode))
		}
	} else {
		log.Println("status", "ERROR", "HTTP:", strconv.Itoa(response.StatusCode), "document:", request.Search.SearchKey, "url:", url, "request:", request)
	}

	return response, nil
}
