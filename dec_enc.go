package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var option string
	flag.StringVar(&option, "s", "dec", "decode(dec)/encode(enc)?")
	flag.Parse()
	if option == "dec" {
		// decode gfwlist.txt
		gfwlist, err := ioutil.ReadFile("gfwlist.txt")
		if err != nil {
			log.Fatal(err)
		}
		value, err := base64.StdEncoding.DecodeString(string(gfwlist))
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile("dec_gfwlist.txt", value, 0644); err != nil {
			log.Fatal(err)
		}
	} else if option == "enc" {
		// encode dec_gfwlist.txt
		decgfwlist, err := ioutil.ReadFile("dec_gfwlist.txt")
		if err != nil {
			log.Fatal(err)
		}

		decString := base64.StdEncoding.EncodeToString(decgfwlist)
		if err := ioutil.WriteFile("gfwlist.txt", []byte(decString), 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("参数异常")
	}

}
