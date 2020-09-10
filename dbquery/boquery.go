package dbquery

// Boquery service
type Boquery interface {
	Query(sql string, noRet, csvOut int) ([]byte, error)
	QueryRetCsv(sql string) (string, error)
	QueryRetBinary(sql string) ([][]interface{}, error)
	QueryNoRet(sql string) error
	UploadCsv(wstable string, values [][]string) error
}
