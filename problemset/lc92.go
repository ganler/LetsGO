/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseLL(last *ListNode, head *ListNode, n int) *ListNode {
	var ret *ListNode = nil
	last_node := head
	var next_node *ListNode = nil
	cur := head.Next
	// ? -> [n] -> ?
	// [last -> cur -> next]
	// [last <- cur: next]
	for i := 0; i < n-1; i++ {
		next_node = cur.Next
		cur.Next = last_node
		last_node = cur
		cur = next_node
	}

	ret = last_node
	head.Next = cur

	if last != nil {
		last.Next = last_node
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