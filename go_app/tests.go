package main

func TestCigarDBInsertion(dbName string) {
	db := OpenDB(dbName)
	cigars := CreateRandomCigars(10)
	for _, cigar := range cigars {
		db.Create(cigar)
	}
}
