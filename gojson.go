package main

import (

	//"strings"

	"strings"

	gojson "github.com/ChimeraCoder/gojson"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

//convenience:
var jQuery = jquery.NewJQuery

type Pet struct {
	name string
}

func NewPet(name string) *js.Object {
	return js.MakeWrapper(&Pet{name})
}

func (p *Pet) Name() string {
	return p.name
}

func (p *Pet) SetName(name string) {
	p.name = name
}

type Gj struct {
	input string
}

func NewGj(input string) *js.Object {
	return js.MakeWrapper(&Gj{input})
}

func (g *Gj) ReadString() string {
	input := strings.NewReader(g.input)
	output, err := gojson.Generate(input, "Test", "FooPkg")
	if err != nil {
		print("err is " + err.Error())
		panic(err)
	} else {
		print("no error")
		print(output)
	}
	print("returning")
	return string(output)
}

func main() {
	js.Global.Set("gj", map[string]interface{}{
		"New": NewGj,
	})

	const btnid = "button#jsonsubmit"
	const inputid = "textarea#jsoninput"
	const outputid = "textarea#gooutput"
	print("registering")

	print("Your current jQuery version is: " + jQuery().Jquery)
	jquery.NewJQuery().Ready(func() {
		print("finding")
		print(jQuery(btnid))
		jQuery(btnid).On(jquery.CLICK, func(e jquery.Event) {
			print("captured click event")

			text := jQuery(inputid).Val()
            print("text is " + text)
			input := strings.NewReader(text)
			output, err := gojson.Generate(input, "Test", "FooPkg")
			if err != nil {
				print("err is " + err.Error())
				panic(err)
			} else {
				print("no error")
				print(output)
			}
			print("returning")
			jQuery(outputid).SetText(string(output))
		})
	})

}
