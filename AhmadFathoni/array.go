package main

import "fmt";

func main() {

	// membuat array manual 
	var name [3]string;

	name[0] = "Ahmad";
	name[1] = "Fathoni";
	name[2] = "Gans";

	fmt.Println(name[0]);
	fmt.Println(name[1]);
	fmt.Println(name[2]);

	// membuat array langsung 

	var values = [3] int {
		10,
		20,
		30,
	};

	fmt.Println(values);
	fmt.Println(values[0]);
	fmt.Println(values[1]);
	fmt.Println(values[2]);

	// " len(array) 		   --> menghitung jumlah panjang array"
	// " array[index]          --> mendapatkan data di posisi index"
	// " array[index] = value  --> mengubah data di posisi index"
	fmt.Println(len(name));
	fmt.Println(len(values));
	
	values[2] = 100;
	fmt.Println(values);
};