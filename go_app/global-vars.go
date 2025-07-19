package main

var (
	createDB    bool
	startServer bool
	server      string
	port        int
	dbPath      string
	dbName      string
)

func CreateTableSchemas() []any {
	tableSchemas := make([]any, 0, 1)
	// tableSchemas = append(tableSchemas, CreateNewCigarTable())
	tableSchemas = append(tableSchemas, &CigarTable{})
	return tableSchemas
}

func CreateNewCigarTable() *CigarTable {
	return &CigarTable{
		Brand:                "Default",
		Name:                 "Default",
		Wrapper:              "Default",
		Profile:              "String",
		TastyTip:             false,
		Binder:               "Default",
		Spicy:                0,
		Rating:               0,
		Length:               54,
		Ring:                 60,
		Review:               "",
		JohnRating:           0,
		JohnReview:           "",
		KyleRating:           0,
		KyleReview:           "",
		AuthenticHumanReview: "",
	}
}
