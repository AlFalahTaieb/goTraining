package main

import "testing"

func TestHello( t *testing.T){
    got := Hello()
    want := "Hello,World :) "

    if got != want {
        t.Errorf("got %q want %q", got , want)
    }
}

// Declaring a Variable X := var ;
// t.Errorf we are calling the Errorf method on our t which will print out a message 
// and fail the test. check : https://pkg.go.dev/fmt#hdr-Printing

