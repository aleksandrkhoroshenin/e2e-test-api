package api

type CustomerWithPrefixResp struct {
	Count int `json:"count"`
	IDs []int64 `json:"ids"`
}

type CustomerPrefixesReq struct {
	Prefixes []string `json:"prefixes"`
}

type DeleteCustomersResp struct {
	Count int `json:"count"`
	IDs []int `json:"ids"`
}

type CreateCustomerReq struct {
	FirstName      *string `json:"first_name"`
	LastName       *string `json:"last_name,omitempty"`
	PatronymicName *string `json:"patronymic_name,omitempty"`
	Phone          *string `json:"phone"`
	Email          *string `json:"email"`
	ID int `json:"id"`
}

type CreateCustomerResp struct {
	CountInserted int `json:"count_inserted"`
}

type CustomerResp struct {
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName,omitempty"`
	PatronymicName *string `json:"patronymicName,omitempty"`
	Phone          *string `json:"phone"`
	Email          *string `json:"email"`
	ID             int     `json:"id"`
}

