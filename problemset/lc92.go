package lc92

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLL(last *ListNode, head *ListNode, n int) *ListNode {
	var ret *ListNode = nil
	lastNode := head
	var nextNode *ListNode = nil
	cur := head.Next
	// ? -> [n] -> ?
	// [last -> cur -> next]
	// [last <- cur: next]
	for i := 0; i < n-1; i++ {
		nextNode = cur.Next
		cur.Next = lastNode
		lastNode = cur
		cur = nextNode
	}

	ret = lastNode
	head.Next = cur

	if last != nil {
		last.Next = lastNode
	}
	return ret
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	origin := head
	var last *ListNode = nil
	for i := 0; i < m-1; i++ {
		last = head
		head = head.Next
	}
	ret := reverseLL(last, head, n-m+1)
	if last == nil && n != m { // r-start is arr[0].
		origin = ret
	}
	return origin
}
