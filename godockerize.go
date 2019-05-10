package main

import "fmt"

// Alpine doesn't do point releases, but if you are reading this, 3.8 downloads
// 3.8.1 or newer, which contains the security fix for this RCE:

const baseDockerImage = "alpine:3.8"

func main() {

	fmt.Println("godockerize")

}
