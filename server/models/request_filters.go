package models

type DefaultFilter struct {
	Limit  int    `schema:"limit,omitempty"`
	Skip   int    `schema:"offset,omitempty"`
	SortBy string `schema:"sortby,omitempty"`
	Asc    bool   `schema:"asc,omitempty"`
}

type ContractsFilter struct {
	Limit        int    `schema:"limit,omitempty"`
	Skip         int    `schema:"skip,omitempty"`
	SortBy       string `schema:"sortby,omitempty"`
	Asc          bool   `schema:"asc,omitempty"`
	ContractName string `schema:"contract_name,omitempty"`
	TokenName    string `schema:"token_name,omitempty"`
	TokenSymbol  string `schema:"token_symbol,omitempty"`
	ErcType      string `schema:"erc_type,omitempty"`
}
