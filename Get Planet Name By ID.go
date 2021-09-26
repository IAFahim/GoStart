package main

func GetPlanetName(ID int) string {
	var arr []string = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}
	return arr[ID-1]
}

func main() {

}
