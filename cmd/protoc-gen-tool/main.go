/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 11:16:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 17:06:39
 * @Description: main
 */
package main

import (
	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
)

const genGoDocURL = "https://developers.google.com/protocol-buffers/docs/reference/go-generated"
const grpcDocURL = "https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code"

func main() {

	// // Define proto option
	// protogenOptions := protogen.Options{}

	// // Add flag.command.Set to paramsFunc
	// // protogenOptions.ParamFunc = flag.CommandLine.Set

	// // To run by protogenOptions
	// protogenOptions.Run(func(plugin *protogen.Plugin) error {

	// 	// Plugin Generate
	// 	// return protocgentoolservice.CreateProtoGenToolService(&protocgentoolservice.Plugin{Plugin: plugin}).Generate()
	// 	return nil
	// })

	// if len(os.Args) == 2 && os.Args[1] == "--version" {
	// 	fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
	// 	os.Exit(0)
	// }
	// if len(os.Args) == 2 && os.Args[1] == "--help" {
	// 	fmt.Fprintf(os.Stdout, "See "+genGoDocURL+" for usage information.\n")
	// 	os.Exit(0)
	// }

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {

		// for _, f := range gen.Files {
		// 	if f.Generate {
		// 		gengo.GenerateFile(gen, f)
		// 	}
		// }

		// return protocgentoolservice.CreateProtoGenToolService(&protocgentoolservice.Plugin{Plugin: gen}).Generate()
		// gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})

}
