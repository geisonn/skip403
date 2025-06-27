package functions

import "fmt"

func Help() {
	fmt.Println("Usage:")
	fmt.Println("  -u,  --url         Specify a single target URL")
	fmt.Println("  -l,  --list        Specify a file with a list of target URLs")
	fmt.Println("  -hf, --hide-fails  Hide failed bypass attempts")
	fmt.Println("  -h,  --help        Show this help message")
}
