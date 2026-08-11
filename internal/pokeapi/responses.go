package pokeapi

type BaseResponse struct {
	Count		int	`json:"count"`
        Next            string	`json:"next"`
        Previous        string	`json:"previous"`
}

type ResponseLocations struct {
	BaseResponse
	Results []struct {
		Name	string	`json:"name"`
		URL 	string	`json:"url"`
	} `json:"results"`
}
