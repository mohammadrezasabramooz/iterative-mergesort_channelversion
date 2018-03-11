package main



func main() {

}

func getItem(m chan int,item int,size int)int  {
	replace:=make(chan int,size)
	ret:=0
	for i:=0;i<size;i++ {
		replace<-<-m
	}
	for i:=0;i<size;i++ {

		if i==item {
			ret=<-replace
			m<-ret
			continue
		}
		m<-<-replace
	}
	return ret
}
func replaceItem(m chan int,item int,size int,input int,fill int)  {
	replace:=make(chan int,size)
	for i:=0;i<item;i++ {
		replace<-<-m
	}

	replace<-input
<-m
	for i:=item;i<fill-1;i++ {
		replace<-<-m
	}
	for i:=0;i<fill;i++ {
		m<-<-replace
	}


}
func mergesort_iterative_cpu(arr chan int,size int){
	temparr:= make(chan int,size)
	for i:=0;i<size;i++{
		temparr<-0
	}
	var right int
	var rend int
	var i int
	var j int
	var m int

	for k:= 1; k < size; k *= 2 {
		//at each partition size, sort and merge
		for  left := 0; left + k < size; left += k*2 {
			//store the start of the right partition and its end
			right = left + k
			rend = right + k

			//if the partitions are uneven, readjust the end
			if rend > size{
				rend = size
			}
			m = left
			i = left
			j = right

			//merge
			for i < right && j < rend {


				if getItem(arr,i,size) <=getItem(arr,j,size) {
					replaceItem(temparr,m,size,getItem(arr,i,size),size)
					//temparr[m] = arr[i]
					i++
				} else {
					replaceItem(temparr,m,size,getItem(arr,j,size),size)
					//temparr[m] = arr[j]
					j++
				}
				m++
			}
			for i < right {
			//	temparr[m] = arr[i]
				replaceItem(temparr,m,size,getItem(arr,i,size),size)
				i++
				m++
			}
			for j < rend {
				//temparr[m] = arr[j]
				replaceItem(temparr,m,size,getItem(arr,j,size),size)
				j++
				m++
			}
			//copy from temp array into initial array
			for m = left; m < rend; m++ {
				replaceItem(arr,m,size,getItem(temparr,m,size),size)
			//	arr[m] = temparr[m]
			}
		}
	}
}