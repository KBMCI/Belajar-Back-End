package main;

import "fmt";

func main() {
	
	// operator perbandingan 
	// " > "  --> lebih dari 
	// " < "  --> kurang dari 
	// " >= " --> lebih dari sama dengan
	// " <= " --> kurang dari sama dengan
	// " == " --> sama dengan
	// " != " --> tidak sama dengan

	var nama1 string = "Aldy";
	var nama2 string = "Imam";

	var hasil bool = nama1 != nama2;
	fmt.Println(hasil);

	var nilaiA = 100;
	var nilaiB = 200;

	fmt.Println(nilaiA > nilaiB);
	fmt.Println(nilaiA < nilaiB);
	fmt.Println(nilaiA == nilaiB);
	fmt.Println(nilaiA != nilaiB);
	

}
