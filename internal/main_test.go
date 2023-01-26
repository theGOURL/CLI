package main
// main.go
// main_test.go
// version.go

import (
  "testing"
  "log"
  "time"
)

func TestVersion(test *testing.T){
  result   := PrintVersion("2.0.1");
  expected := "Version: 2.0.1";
  if result != expected {
    log.Println(time.Now(),"\n[ERROR]\nResult: ",result,"\nExpected: ",expected);
  }
}

func Benchmarkmain(benchmark *testing.B){
  for i := 0; i < benchmark.N; i++{
    main();
  }
}

func Examplemain(){
  main();
  //Output:
  //Version:  2.0.1
  //linux
  //[Command successfully executed]
 //gabriel🧠 UNDEFINIED COMMAND 
}
