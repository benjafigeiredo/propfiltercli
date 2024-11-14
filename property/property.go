package property

type Property struct {
	squareFootage int
	lighting      string
	price         float64
	rooms         int
	bathrooms     int
	location      []int
	description   string
	ammenities    map[string]bool
}

func (p *Property) SetSquareFootage(squareFootage int) {
	p.squareFootage = squareFootage
}

func (p *Property) SetLighting(lighting string) {
	p.lighting = lighting
}

func (p *Property) SetPrice(price float64) {
	p.price = price
}

func (p *Property) SetRooms(rooms int) {
	p.rooms = rooms
}

func (p *Property) SetBathrooms(bathrooms int) {
	p.bathrooms = bathrooms
}

func (p *Property) SetLocation(location []int) {
	p.location = location
}

func (p *Property) SetDescription(description string) {
	p.description = description
}

func (p *Property) SetAmmenities(ammenities map[string]bool) {
	p.ammenities = ammenities
}

func (p *Property) GetSquareFootage() int {
	return p.squareFootage
}

func (p *Property) GetLighting() string {
	return p.lighting
}

func (p *Property) GetPrice() float64 {
	return p.price
}

func (p *Property) GetRooms() int {
	return p.rooms
}

func (p *Property) GetBathrooms() int {
	return p.bathrooms
}

func (p *Property) GetLocation() []int {
	return p.location
}

func (p *Property) GetDescription() string {
	return p.description
}

func (p *Property) GetAmmenities() map[string]bool {
	return p.ammenities
}

func (p *Property) GetAmmenity(ammenity string) bool {
	return p.ammenities[ammenity]
}
