package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"nama":   "Sahril Mahendra",
		"nim":    "16050974009",
		"alamat": "Sidoarjo",
	}
	mahasiswa["gender"] = "Pria"

	fmt.Println(len(mahasiswa))
	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["nim"])
	fmt.Println(mahasiswa["alamat"])

	delete(mahasiswa, "alamat")
	fmt.Println(mahasiswa)
}
