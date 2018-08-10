package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	t := viper.Get("hobbies") // this would be "steve"
	fmt.Println(t)
	t = viper.Get("clothing.jacket")
	fmt.Println(t)
}
