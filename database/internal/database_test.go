package mytests

import (
	"fmt"
	"testing"

	"github.com/yingshaoxo/gopython/database"
)

var host = "127.0.0.1"
var port = "27017"
var user = "root"
var password = "yingshaoxo666"

func Test_database(t *testing.T) {
	mongoDB, _ := database.MongoDB(host, port, user, password)

	list, _ := mongoDB.List_database()
	fmt.Println(list)

	mongoDB.Detele_database("config")

	mongoDB.Stop()
}
