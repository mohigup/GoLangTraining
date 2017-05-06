package marshal_unmarshal_encode_decode

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
)

type person struct {
	First string
	Last string `json:"-"`
	Age int `json:"wisdom-score"`
}

func Marshalling (){
	fmt.Println("-------MARSHALLING-STARTED---------")
	p1 := person{"mohit", "gupta", 22}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Println("-------MARSHALLING-ENDED-----------")
}

func UnMarshalling (){
	fmt.Println("\n-------UNMARSHALLING-STARTED---------")
	var p1 person
	fmt.Println(p1)

	bs := []byte(`{"First": "mohit", "Last": "gupta", "wisdom-score":20}`)
	json.Unmarshal(bs,&p1)
	fmt.Println(p1)
	fmt.Println("-------UNMARSHALLING-ENDED-----------")

}

func Encode(){
	p1 := person{"mohit", "gupta", 22}
	json.NewEncoder(os.Stdout).Encode(p1)
	fmt.Println("-------------------------------")
}

func Decode(){
	var p1 person
	input := `{"First": "mohit", "Last": "gupta", "wisdom-score":20}`
	json.NewDecoder(strings.NewReader(input)).Decode(&p1)
	fmt.Println("-------------------------------")
}