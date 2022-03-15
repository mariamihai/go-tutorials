package main

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var initAlbums = []album{
	{ID: "1", Title: "Title 1", Artist: "Artist 1", Price: 10.12},
	{ID: "2", Title: "Title 2", Artist: "Artist 2", Price: 14.48},
	{ID: "3", Title: "Title 3", Artist: "Artist 1", Price: 9.56},
}
