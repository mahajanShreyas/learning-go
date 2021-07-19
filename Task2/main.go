package main

func main() {

	ll1 := LinkedList{}

	ll1.Insert(2)
	ll1.Insert(3)
	ll1.Insert(4)
	ll1.Insert(5)
	ll1.Insert(6)

	ll1.Traverse()

	ll2 := LinkedList{}

	ll2.Insert(20)
	ll2.Insert(30)
	ll2.Insert(40)
	ll2.Insert(50)
	ll2.Insert(60)

	ll2.Traverse()

	ll3 := Merge(ll1, ll2)
	ll3.Traverse()


}