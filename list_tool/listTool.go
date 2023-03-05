package list_tool

import "math/rand"

func GetRandomStringFromList(list []string) string {
	return list[rand.Intn(len(list))]
}
