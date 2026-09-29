package main

func main() {
	var umur int8
	var suhu float32

	suhu = 36.3
	suhu = 10

	println("Umur: ", umur)
	println("Suhu: ", suhu)
	//mengetahui memori
	println("Alamat memori dari var suhu ", &suhu)
	println("Alamat memori dari var umur ", &umur)
}
