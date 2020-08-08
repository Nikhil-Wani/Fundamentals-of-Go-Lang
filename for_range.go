package main
import "fmt"

func main(){
	nums:=[]int{1,2,3}
	sums:=0
	for _,value:=range nums{
		sums+= value
	}
	fmt.Println("sums:",sums)

	for i,num:=range nums{
		if num==3{
			fmt.Println("index:",i)
		}
	}
	kvs:= map[string]string{"1":"mango","2":"apple","3":"banana"}
	for k,v:=range kvs{
		fmt.Println(k,v)
	}
	for k:=range kvs{
		fmt.Println("key:",k)
	}
	for i,c:=range "Hi"{
		fmt.Println(i,c)
	}
}


