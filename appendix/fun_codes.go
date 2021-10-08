package appendix

//自己实现三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

//组织main函数退出的方法
//func main() {
//	defer func() { for {} }()
//}
//func main() {
//	defer func() { select {} }()
//}
//func main() {
//	defer func() { <-make(chan bool) }()
//}

//基于管道的随机数生成器
func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select { case c <- 0: case c <- 1: }
		}
	}()
	return c
}
