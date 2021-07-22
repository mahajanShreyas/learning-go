package main

func Merge(ll1 LinkedList, ll2 LinkedList) LinkedList {

	newList := LinkedList{}
	temp1 := ll1.start
	temp2 := ll2.start

	for temp1 != nil && temp2 != nil {
		if temp1.data < temp2.data {
			newList.Insert(temp1.data)
			temp1 = temp1.next
		} else {
			newList.Insert(temp2.data)
			temp2 = temp2.next
		}
	}

	for temp1 != nil {
		newList.Insert(temp1.data)
		temp1 = temp1.next
	}

	for temp2 != nil {
		newList.Insert(temp2.data)
		temp2 = temp2.next
	}

	return newList
}