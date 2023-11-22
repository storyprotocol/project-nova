package v0alpha

/*
{
	pagination?: {
		offset?: number  // starting from 0
		limit?: number
	}
}
*/

/*
	{
		waitForTransaction?: boolean
	}
*/
type QueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
}
