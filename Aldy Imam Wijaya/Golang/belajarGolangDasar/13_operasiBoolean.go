package main;

import "fmt";

func main() {
	
	//operasi boolean 
	// " && " --> dan
	// " || " --> atau
	// " ! "  --> kebalikan

	var nilaiQuiz int8 = 80;
	var nilaiUjian uint = 75;

	var lulusNilaiQuiz bool = nilaiQuiz > 60;
	var lulusNilaiUjian bool = nilaiUjian > 60;

	var kelulusan bool = lulusNilaiQuiz && lulusNilaiUjian;
	fmt.Println(kelulusan);


};