package main

func main() {
	// counter vehicle in parking lot
	var cars int8 = 10      // can store from -128 to 127
	var bikes int16 = 300   // can store from -32,768 to 32,767
	var trucks int32 = 5000 // can store from -2,147,483,648 to 2,147,483,647

	// download list
	var picture byte = 255   // can store from 0 to 255 or equivalently uint8
	var video rune = 40000   // can store from -2,147,483,648 to 2,147,483,647 or equivalently int32
	var movies uint = 100000 // can store from 0 to 4,294,967,295 or equivalently minimal uint32
	var songs int = 50000    // can store from -2,147,483,648 to 2,147,483,647 or equivalently minimal int32

	// GPS coordinates
	var latitude float32 = 37.7749    // can store up to 7 decimal digits of precision
	var longitude float64 = -122.4194 // can store up to 15 decimal digits of precision

	println("we have cars: ", cars)
	println("we have bikes: ", bikes)
	println("we have trucks: ", trucks)

	println("picture size in KB: ", picture)
	println("video size in MB: ", video)
	println("movies size in MB: ", movies)
	println("songs size in MB: ", songs)

	println("latitude: ", latitude)
	println("longitude: ", longitude)
}
