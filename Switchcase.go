package main
import "fmt"

func main(){
	fmt.Print("enter a no.");
	var in int;
	fmt.Scanln(&in);

	switch(in){
	case 10:
		fmt.Print(" no 10 \n");fallthrough;
	case 20:
		fmt.Print(" no 20 \n");fallthrough;
	case 30:
		fmt.Print(" no 30\n");fallthrough;
	default:
		fmt.Print("def");

	}
}
