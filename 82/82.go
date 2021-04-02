package _2

func main() {
	//用两个指针来移动
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	pro := &ListNode{0, head}
	pre := pro
	cur := pro.Next
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			x := cur.Val
			for cur.Next != nil && x == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return pro.Next
}
