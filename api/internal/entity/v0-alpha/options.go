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
	First int
	Skip  int
}
