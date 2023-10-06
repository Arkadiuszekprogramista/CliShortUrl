package models

// Coordinates struct represents the coordinates in the JSON data.
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Timezone struct represents the timezone in the JSON data.
type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

// Location struct represents the location in the JSON data.
type Location struct {
	Street     struct {
		Number int    `json:"number"`
		Name   string `json:"name"`
	} `json:"street"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	Country    string      `json:"country"`
	Postcode   interface{}         `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone   Timezone    `json:"timezone"`
}

// Name struct represents the name in the JSON data.
type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

// Login struct represents the login information in the JSON data.
type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

// Dob struct represents the date of birth in the JSON data.
type Dob struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

// Registered struct represents the registration date in the JSON data.
type Registered struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

// Picture struct represents the picture URLs in the JSON data.
type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

// Result struct represents the individual result in the JSON data.
type Result struct {
	Gender     string    `json:"gender"`
	Name       Name      `json:"name"`
	Location   Location  `json:"location"`
	Email      string    `json:"email"`
	Login      Login     `json:"login"`
	Dob        Dob       `json:"dob"`
	Registered Registered `json:"registered"`
	Phone      string    `json:"phone"`
	Cell       string    `json:"cell"`
	ID         struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	} `json:"id"`
	Picture Picture `json:"picture"`
	Nat     string  `json:"nat"`
}

// Info struct represents the additional information in the JSON data.
type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}

// Response struct represents the entire JSON response.
type Response struct {
	Results []Result `json:"results"`
	Info    Info     `json:"info"`
}