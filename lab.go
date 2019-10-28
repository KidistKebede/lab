package main

import "fmt"

	func main(){
		combine([]string{"a","b","c"}, []string{"1","2","3"})
		var n,i,j int
		j=0
		fmt.Print("enter the number")
		fmt.Scanln(&n)
		for i=1;i<=n;i++{
			fmt.Print(fib(j),",")
			j++
			fmt.Println()

		}
     
		
	}

		func combine(a []string, b []string){
			
			if len(a)==len(b){
				var comb []string
				for i:=0; i<len(a); i++{
					
					comb = append(comb, a[i], b[i])
					
				} 
				fmt.Print(comb)
				fmt.Println()
				fmt.Printf("len=%d cap=%d" ,len(comb), cap(comb))
			}else{
				fmt.Print("The given slices can not be combined")
	
			}
		}

	func fib(n int)int{
		if n==0||n==1{
			return n
		}else {
			return (fib(n-1)+fib(n-2))
		}
	}
		