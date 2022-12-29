package main;

import "fmt";

func main() {

	// Operasi Matematika 
	// " + " untuk penjumlahan
	// " - " untuk pengurangan 
	// " * " untuk perkalian 
	// " / " untuk pembagian
	// " % " untuk modulus

	var a = 10;
	var b = 6;
	var c = a + b;
	
	fmt.Println(c);

	// atau

	var hasil = 10 - 8;
	fmt.Println(hasil);

	// Augmented Assignments
	// a = a + 10 bisa ditulis menjadi --> a += 10

	var i = 10;
	i += 10;

	fmt.Println(i);

	// atau 

	var k = 20;
	k = k + 20;

	fmt.Println(k);

	// unary operation 
	// " ++ " 	--> a = a + 1
	// " -- "  	--> a = a - 1
	// " + "  	--> positive 
	// " - "  	--> negative
	// " ! "  	--> boolean kebalikan

	var r = 9;
	r++;

	fmt.Println(r);

	var e = 10; 
	e--;
	e--;

	fmt.Println(e);

	var negative = -150;
	var positive = +150;
	
	fmt.Println(negative);
	fmt.Println(positive);

	// var kebalikan = !5
	// !kebalikan;

	// fmt.Println(kebalikan)

};