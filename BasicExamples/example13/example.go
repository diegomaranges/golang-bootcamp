/*### Composition over inheritance
Like we stated previously in the introduction. Go is all about composition, in fact, there is no inheritance. If you want to share behaviour and data of something you compose something with that type, no need of subclassing and specifying a complex tree of class hierarchy and inheritance.
Coming from an OOP world this might be weird at first. But composition is actually a well understood concept of OOP and in Go you will use it quite a lot. To learn a bit more about how this works you can read [this section](https://golang.org/doc/effective_go.html#embedding) of Effective Go to see some real-live examples of how you can benefit this.
Lets take a look at a simple example of how you would use this([GoPlay](https://goplay.space/#AnaEldEzcBl)):
```go*/
package main

import "fmt"

// User holds information of a given user.
type User struct {
	ID             int
	Name, Location string
}

// Here you can see that player embeds the User type.
// This way we are saying that a Player holds is composed
// of a User and a GameID.
type Player struct {
	User
	GameID int
}

func main() {
	p := Player{}
	p.ID = 42
	p.Name = "Globant"
	p.Location = "La Plata"
	p.GameID = 90404
	fmt.Printf("%+v", p)
	// This will print the following:
	// {User:{ID:42 Name:Globant Location:La Plata} GameID:90404}
	// You can see that the User type is embedded in the Player
	// structure.
}
