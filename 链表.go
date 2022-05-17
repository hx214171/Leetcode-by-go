							// 简单题
// # 剑指 Offer 06. 从尾到头打印链表
// # 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 思路1：遍历链表同时将每个元素加入数组，然后倒序数组元素并返回。
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

// 思路2：翻转链表
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

// 思路3：递归
func reversePrint(head *ListNode) []int {
    if head == nil {return nil}
    return append(reversePrint(head.Next), head.Val)
}

// # 剑指 Offer 18. 删除链表的节点
// # 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// # 返回删除后的链表的头节点。
// # Definition for singly-linked list.
// 思路：双指针cur和pre，如果要删除第一个节点则直接返回head.Next，
// 否则通过pre.Next = cur.Next删除指定节点。
func deleteNode(head *ListNode, val int) *ListNode {
    var cur *ListNode
    var pre *ListNode
    if head.Val == val {return head.Next}
    pre, cur = head, head.Next
    for cur != nil {
        if cur.Val == val {
            pre.Next = cur.Next
        }
        pre, cur = cur, cur.Next
    }
    return head
}

// # 剑指 Offer 22. 链表中倒数第k个节点
// # 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，
// # 即链表的尾节点是倒数第1个节点。
// # 例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。
// # 这个链表的倒数第3个节点是值为4的节点。
// # Definition for singly-linked list.
// 思路：快慢指针，快指针cur先走k步，然后cur和pre一起走，当cur走完，pre即为倒数第k个节点。
func getKthFromEnd(head *ListNode, k int) *ListNode {
    var cur *ListNode
    var pre *ListNode
    cur = head
    pre = head
    for i := 0; i < k; i++ {
        cur = cur.Next
    }
    for cur != nil {
        pre, cur = pre.Next, cur.Next
    }
    return pre
}

// # 剑指 Offer 24. 反转链表
// # 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点
// 思路1：通过三个变量cur pre tmp， tmp := cur.Next；cur.Next, pre = pre, cur； cur = tmp 反转每个节点。
// 可以作为剑指 Offer 06. 从尾到头打印链表的解法。
func reverseList(head *ListNode) *ListNode {
    // if head == nil {return nil}
    var pre *ListNode
    cur := head
    for cur != nil {
        tmp := cur.Next
        cur.Next, pre = pre, cur
        cur = tmp 
    }
    return pre
}

思路2：递归
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { 
        //只有一个或者没有节点的时候返回 
        return head
    }
    p := reverseList(head.Next)
    //剩下两个节点head和head.Next,翻转两个节点的指针。
    head.Next.Next = head
    head.Next = nil
    return p
}

// # 剑指 Offer 25. 合并两个排序的链表
// # 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
思路1：双指针，分别比较遍历并比较两个链表，将较小的元素加入新链表，遍历完成一个链表后，
将剩下的链表整体加入。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dum := new(ListNode)
    cur := dum
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
             cur.Next = l1
             l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 != nil {
        cur.Next = l1
    } else {
        cur.Next = l2
    }
    return dum.Next
}

思路2：递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {return l2}
    if l2 == nil {return l1}
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    }
}

// # 剑指 Offer 52.两个链表的第一个公共节点
// # 输入两个链表，找出它们的第一个公共节点。
思路：双指针，如果有交点双指针分别遍历两个链表即可相遇。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {return nil}
    p, q := headA, headB
    //注意是公共节点，意思是节点指针相等 p == q 而不是值相等p.val = q.val
    for p != q {
        if p == nil {
            p = headB
        } else {
            p = p.Next
        }
        if q == nil {
            q = headA
        } else {
            q = q.Next
        }
    }
    return p
} 

//                                               中等题
// # 剑指 Offer 35. 复杂链表的复制
// # 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，
// 每个节点除了有一个 next 指针指向下一个节点，
// # 还有一个 random 指针指向链表中的任意节点或者 null。
// 思路：遍历两次原始链表，第一次遍历同时创建新节点保存到哈希表里，
// 第二次将新节点的random和next指针分别指向节点。
func copyRandomList(head *Node) *Node {
    if head == nil {return nil}
    dic := map[*Node]*Node{}
    for cur := head; cur != nil; cur = cur.Next {
        dic[cur] = &Node{Val:cur.Val}
    }
    for cur := head; cur != nil; cur = cur.Next {
        dic[cur].Next = dic[cur.Next]
        dic[cur].Random = dic[cur.Random] 
    }  
    return dic[head]
}