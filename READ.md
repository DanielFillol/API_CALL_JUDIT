# API_Call_JUDIT
Simple example script for calling [JUDIT API](https://docs.judit.io/introduction).

## Features
- Log.txt
- Batching processing
- Multithreading Asynchronous request
- Deals with pagination
- Deals with the three requests necessary to complete each call
- JSON returns as .csv

## Setup
- You must have a .csv file with a single column that contains all documents for requesting the API
- You must save it as ```requests.csv``` on ```data``` folder
- The returned files will be saved on ```data/resonses``` for each batch and merged on the root of the project with the name: ```merged_result.csv```
- Create a .env file with your authorization token naming it "AUTH"
```bash
AUTH = myApiKey
```

# Run
```GO
 go run main.go
```