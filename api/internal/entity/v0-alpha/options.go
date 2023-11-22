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
		Offset int `json:"offset,omitempty"`
		Limit  int `json:"limit,omitempty"`
	} `json:"pagination,omitempty"`
}
