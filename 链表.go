							// 简单题
# 剑指 Offer 06. 从尾到头打印链表
# 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
思路1：遍历链表同时将每个元素加入数组，然后倒序数组元素并返回。
func reversePrint(head *ListNode) []int {
    arr := []int{}
    res := []int{}
    if head == nil {return nil}
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    for i := len(arr)-1; i >= 0; i-- {
        res = append(res, arr[i])
    } 
    return res
}

思路2：翻转链表
func reversePrint(head *ListNode) []int {
    if head == nil {return nil}
    var p *ListNode
    for head != nil {
        tmp := head.Next
        head.Next = p
        p = head
        head = tmp
    }
    res := []int{}
    for p != nil {
        res = append(res, p.Val)
        p = p.Next
    }
    return res
}

思路3：递归
func reversePrint(head *ListNode) []int {
    if head == nil {return nil}
    return append(reversePrint(head.Next), head.Val)
}

// # 剑指 Offer 18. 删除链表的节点
// # 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// # 返回删除后的链表的头节点。
// # Definition for singly-linked list.