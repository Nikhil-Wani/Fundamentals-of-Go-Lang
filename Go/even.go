package main
import "fmt"

func main(){
	fmt.Print("enter a no");
	var in int;
	fmt.Scanln(&in);
	if (in%2==0){
		fmt.Print("even:");
	} else {
		fmt.Print("odd");
	}
}
