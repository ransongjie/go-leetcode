package pass1
/**
 * 剑指 Offer II 021. 删除链表的倒数第 n 个结点
 */

 /**
     * 双指针同向移动
     * - 构建伪头结点
     * - p指向head，pd指向dummy->head
     * - p先走n步
     * - p和pd同时往前走，直到p走到链表尾结点，pd指向倒数第n+1个节点
     * - 删除倒数第n个结点
     * 
     * @param head
     * @param n
     * @return
     */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// xcrj 创建结构体 取地址符号
	var dummy *ListNode=&ListNode{0,head}
	var pd *ListNode=dummy
	var p *ListNode=head
	// 
	for i:=0;i<n;i++{
		p=p.Next
	}
	// xcrj nil null
	for p!=nil {
		p=p.Next
		pd=pd.Next
	}
	//
	pd.Next=pd.Next.Next
	return dummy.Next
}

    /**
     * 栈
     * - 构建虚伪头结点
     * - 全部结点放入栈中
     * - pop n次
     * - 栈顶是倒数第n+1个结点
     * @param head
     * @param n
     * @return
     */
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//
	var dummy *ListNode=&ListNode{0,head}
	var p *ListNode=dummy
	// xcrj slice len必传参数
	var stack []*ListNode=make([]*ListNode,0)
	for p!=nil{
		// xcrj push
		stack=append(stack,p)
		p=p.Next
	}
	// pop
	// for i:=0;i<n;i++{
	// 	stack=stack[:len(stack)-1]
	// }
	// pop n次
	stack=stack[:len(stack)-n]
	// peak
	var pre *ListNode=stack[len(stack)-1]
	//
	pre.Next=pre.Next.Next
	return dummy.Next
}
