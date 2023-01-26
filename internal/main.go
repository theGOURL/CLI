package main

import (
  "github.com/theGOURL/cli"
  "fmt"
)

//This constant is responsible for storing the version of CLI
const VERSION = "2.0.1";

func main(){
	PrintVersion(VERSION);
	cli.Start();
}

//This function s responsible for printing the CLI version on the console
func PrintVersion(version string) string{
  fmt.Println("Version: ", version);
  return "Version: " + version;
}
