package main

import "fmt"

// Media - Generic type
// TODO: embedding Rateable here is idiomatic Go or not?
type Media interface {
	MovieDisk | MusicDisk
	Rateable
}

// Could be used for types other than Media. Like Studios, People (Artists)
type Rateable interface {
	getRating() float32
}

// common fields are repeated.
// TODO: explore type embedding for simulating Java/C++ like inheritance seperately,
// but need to think differently in Go
type MovieDisk struct {
	name             string
	genre            string
	length           float32
	year             int
	metacriticRating float32
	imdbRating       float32
}

type MusicDisk struct {
	name                string
	genre               string
	length              float32
	year                int
	googleRating        float32
	rollingStonesRating float32
	ucrRating           float32
}

func (movieDisk MovieDisk) getRating() float32 {

	return (movieDisk.imdbRating + movieDisk.metacriticRating) / 2

}

func (musicDisk MusicDisk) getRating() float32 {
	return (musicDisk.googleRating + musicDisk.ucrRating + musicDisk.rollingStonesRating) / 3
}

func main() {

	fmt.Println("Generics in Go...")

	inception := MovieDisk{
		name:             "Inception",
		genre:            "Thriller/Sci-Fi",
		length:           148,
		year:             2010,
		metacriticRating: 7.4,
		imdbRating:       8.8,
	}

	djangoUnchained := MovieDisk{
		name:             "Django Unchained",
		genre:            "Western",
		length:           165,
		year:             2012,
		metacriticRating: 8.1,
		imdbRating:       8.4,
	}

	_12AngryMen := MovieDisk{
		name:             "12 Angry Men",
		genre:            "Drama",
		length:           96,
		year:             1957,
		metacriticRating: 9.6,
		imdbRating:       9,
	}

	ledZeppelinIII := MusicDisk{
		name:                "Led Zeppelin III",
		genre:               "Heavy Metal",
		year:                1970,
		length:              43.04,
		googleRating:        9.5,
		ucrRating:           9.6,
		rollingStonesRating: 9.4,
	}

	ledZeppelinII := MusicDisk{
		name:                "Led Zeppelin II",
		genre:               "Heavy Metal",
		year:                1960,
		length:              41.38,
		googleRating:        9.8,
		ucrRating:           9.6,
		rollingStonesRating: 9.0,
	}

	ledZeppelinI := MusicDisk{
		name:                "Led Zeppelin I",
		genre:               "Heavy Metal",
		year:                1960,
		length:              44.45,
		googleRating:        9.0,
		ucrRating:           9.1,
		rollingStonesRating: 8.0,
	}

	movieDisks := []MovieDisk{djangoUnchained, inception, _12AngryMen}
	musicDisks := []MusicDisk{ledZeppelinII, ledZeppelinI, ledZeppelinIII}

	fmt.Printf("Pre-Sorting : %v \n", movieDisks)
	SortBasedOnRating(&movieDisks)
	fmt.Printf("Post-Sorting : %v \n", movieDisks)

	fmt.Println("************")

	fmt.Printf("Pre-Sorting : %v \n", musicDisks)
	SortBasedOnRating(&musicDisks)
	fmt.Printf("Post-Sorting : %v \n", musicDisks)

}

// Generic type Media being used. MusicDisk and MovieDisk implements method
// getSorting differently.
func SortBasedOnRating[T Media](disks *[]T) {

	for i := 0; i < len(*disks); i++ {
		for j := i + 1; j < len(*disks); j++ {
			if ((*disks)[i]).getRating() < (*disks)[j].getRating() {
				temp := (*disks)[i]
				(*disks)[i] = (*disks)[j]
				(*disks)[j] = temp
			}
		}

	}

}
