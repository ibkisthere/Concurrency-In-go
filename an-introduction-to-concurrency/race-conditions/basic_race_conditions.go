package main

import (
	"fmt"
	"time"
)

//a race condition occurs when two condition must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained

// this shows up in a data rae , where one concurrent operation tries to read a variable that is undeterministically being written to by another concurrent operation

// in go you use the go keyword to run a function concurrently , doing so , we create what is called a goroutine

// func main(){
// 	var data int
// 	go func(){
// 		data++
// 	}()
// 	// a lot of developers fall into the trap of putting sleeps in their coe because they feel it will solve their concurrency problems
// 	// no you have not solved your data race this way , the longer you sleep the more your program asymptotically approaches logical correctness , it will never be logically correct 
// 	time.Sleep(1 * time.Second)
// 	if data == 0 {
// 		fmt.Printf("the value is %v.\n", data)
// 	}
// }

// running this code, line 5 is executed first, before line 3 , most times data races are introduced because the developers are thinking about the problem sequentially 

//Atomicity 
// when smthing has the property of atomicity , or is considered atomic , it menas within the context that it is operating,, it is indivisible , or uninterruptible 

// atomicity depends on the currently defined scope

// Something may be atomic in one context, but not another. Operations that are atomic within the context of your process may not be atomic in the context of the operating system; operations that are atomic within the context of the operating system may not be atomic within the con‐ text of your machine; and operations that are atomic within the context of your machine may not be atomic within the context of your application.

// WHY IS ATOMICITY IMPORTANT?
//it is important because if something is atomic , it is safe within concurrency context
// it allows us to compose logically correct programs 

// MEMORY ACCESS SYNCHRONIZATION
func main(){
	var data int
	go func(){
		data++
	}()
	// a lot of developers fall into the trap of putting sleeps in their coe because they feel it will solve their concurrency problems
	// no you have not solved your data race this way , the longer you sleep the more your program asymptotically approaches logical correctness , it will never be logically correct 
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
