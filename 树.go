//                							简单题
// # 剑指 Offer 27. 二叉树的镜像
// # 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
思路1：递归
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
    return root
}

思路2：栈
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {return nil}
    stack := []*TreeNode{root}
    for len(stack) != 0 {
        node := stack[0]
        stack = stack[1:]
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        node.Left, node.Right = node.Right, node.Left
    }
    return root
}

// # 剑指 Offer 28. 对称的二叉树
// # 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
思路：从上往下递归，判断个每个节点对称，节点的左子节点的左子节点应该与节点右子节点的右子节点对称，
节点的左子节点的右子节点应该与节点的右子节点的左子节点对称。
func isSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    return rec(root.Left, root.Right)
}

func rec(L, R *TreeNode) bool {
    if L == nil && R == nil {return true}
    if L == nil || R == nil || L.Val != R.Val {return false}
    return rec(L.Left, R.Right) && rec(L.Right, R.Left)
}

// # 剑指 Offer 54. 二叉搜索树的第k大节点
// # 给定一棵二叉搜索树，请找出其中第k大的节点。
思路：二叉搜索树的中序遍历是递增的，中序遍历的倒序就是递减的，递减的第k个值即为第k大的节点。
var count int
var res int 
func kthLargest(root *TreeNode, k int) int {
    count = k
    dfs(root)
    return res
}

func dfs(root *TreeNode) {
    if root == nil {return}
    dfs(root.Right)
    count --
    if count == 0 {
        res = root.Val
        return
    }
    dfs(root.Left)
}

// 剑指 Offer 55 - I. 二叉树的深度
// 输入一棵二叉树的根节点，求该树的深度。
// 从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
思路1：dfs，二叉树深度为 左子树深度与右子树深度中的最大值 +1
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int{
    if x > y {
        return x
    } else {
        return y
    }
}

思路2：bfs，遍历每层的节点同时res+1
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    queue, res := []*TreeNode{root}, 0
    for len(queue) != 0 {
        tmp := []*TreeNode{}
        for _,node := range queue {
            if node.Left != nil {
                tmp = append(tmp, node.Left)
            }
            if node.Right != nil {
                tmp = append(tmp, node.Right)
            }
        }
        res += 1
        queue = tmp
    }
    return res
}

// # 剑指 Offer 55 - II. 平衡二叉树
// # 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，
// # 那么它就是一棵平衡二叉树。
思路1：后序遍历+剪枝
func isBalanced(root *TreeNode) bool {
    return dfs(root) != -1
}

func dfs(root *TreeNode) int {
    if root == nil {return 0}
    left := dfs(root.Left)
    if left == -1{return -1}
    right := dfs(root.Right)
    if right == -1 {return -1}
    if abs(left-right) > 1 {
        return -1
    }
    return max(left, right) + 1
}

func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}

思路2：前序遍历，自顶向下。
func isBalanced(root *TreeNode) bool {
    if root == nil {return true}
    return isBalanced(root.Left) && isBalanced(root.Right) && abs(dfs(root.Left)-dfs(root.Right)) <= 1
}

func dfs(root *TreeNode) int {
    if root == nil {return 0}
    return max(dfs(root.Left), dfs(root.Right))+1
}

func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}

// # 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
// # 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
思路1：迭代
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for {
        if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        } else if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else {
            return root
        } 
    }
}

思路2：递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } 
    return root
}

// # 剑指 Offer 68 - II. 二叉树的最近公共祖先
// # 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || p == root || q == root {return root}
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left == nil {return right}
    if right == nil {return left}
    return root
  }

// #                                               中等题
// # 剑指 Offer 07. 重建二叉树
// # 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
// # 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
思路：通过根节点创建一个新二叉树，找到根节点在中序遍历中的索引，递归创建根节点左子树和右子树。
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {return nil}
    root := &TreeNode{Val:preorder[0]}
    //通过根节点新建一个二叉树节点
    index := 0
    for i := range inorder {
        if inorder[i] == preorder[0] {
            index = i
        }
    }
    root.Left = buildTree(preorder[1:index+1], inorder[:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])
    return root
}


// # 剑指 Offer 26. 树的子结构
// # 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
// # B是A的子结构， 即 A中有出现和B相同的结构和节点值。
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil {return false}
    return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
    //遍历每个节点的同时判断每个节点为根节点的子树是否包含树B
}

func recur(A *TreeNode, B *TreeNode) bool {
    if B == nil {return true}
    if A == nil || A.Val != B.Val {return false}
    return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

// # 剑指 Offer 32 - I. 从上到下打印二叉树
// # 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
func levelOrder(root *TreeNode) []int {
    queue, res := []*TreeNode{root}, []int{}
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        if node != nil {
            res = append(res, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return res
}

// # 剑指 Offer 32 - II. 从上到下打印二叉树 II
// # 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
思路：和32-1不同的地方在于每层加入一个节点数组。
func levelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    res := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        l := len(queue)
        tmp := []int{}
        for i := 0; i < l; i++{
            node := queue[0]
            queue = queue[1:]
            if node != nil {
                tmp = append(tmp, node.Val)
                if node.Left != nil {
                    queue = append(queue, node.Left)
                }         
                if node.Right != nil {
                    queue = append(queue, node.Right)
                }  
            }
        }
    res = append(res, tmp)
    }
    return res
}

// # 剑指 Offer 32 - III. 从上到下打印二叉树 III
// # 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，
// # 第三行再按照从左到右的顺序打印，其他行以此类推。
func levelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    queue := []*TreeNode{root}
    res := [][]int{} //此时len(res)为0
    for len(queue) > 0 {
        l := len(queue)
        tmp := make([]int, l)
        for i := 0; i < l; i++ {
            if len(res) % 2 == 0{
                tmp[i] = queue[i].Val
            } else {
                tmp[i] = queue[l-1-i].Val
            }
            if queue[i].Left != nil {queue = append(queue, queue[i].Left)}
            if queue[i].Right != nil {queue = append(queue, queue[i].Right)}
        }
        queue = queue[l:]
        res = append(res, tmp)
    }
    return res
}

// # 剑指 Offer 33. 二叉搜索树的后序遍历序列
// # 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。
// # 假设输入的数组的任意两个数字都互不相同。
思路1：递归，后续遍历是左子树右子树根节点，所以后续遍历postorder列表中最后一个值为根节点，又因为是二叉搜索树，左子树的节点值都应该小于根节点的值，
所以从postorder列表的首元素往后遍历找到第一个大于根节点值的元素，这个值应该是根节点的右子树的开始节点，从这个值开始到根节点的元素值都应该大于根节点，
否则返回false。

func verifyPostorder(postorder []int) bool {
    return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i int, j int) bool {
    if i >= j {return true}
    mid := i
    root := postorder[j]
    for postorder[mid] < root {
        mid++
    }
    tmp := mid
    for tmp < j {
        tmp++
        if postorder[tmp] < root{
            return false
        }
    }
    return recur(postorder, i, mid-1) && recur(postorder, mid, j-1)
}

思路2：辅助栈


// # 剑指 Offer 34. 二叉树中和为某一值的路径
// # 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
// # 从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
func pathSum(root *TreeNode, target int) [][]int {
    res := [][]int{}
    path := []int{}
    dfs(root, target, path, &res)
    return res
}

func dfs(root *TreeNode, target int, path []int, res *[][]int) {
    if root == nil { return}
    if root.Val == target && root.Left == nil && root.Right == nil {
        tmp := make([]int, len(path)+1)
        copy(tmp, append(path, root.Val))
        *res = append(*res, tmp)
    }
    dfs(root.Left, target-root.Val, append(path, root.Val), res)
    dfs(root.Right, target-root.Val, append(path, root.Val), res)
}


// # 剑指 Offer 36. 二叉搜索树与双向链表
// # 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，
// # 只能调整树中节点指针的指向。
func Convert(root *TreeNode ) *TreeNode {
    var dfs func(cur *TreeNode)
    var pre, head *TreeNode
    dfs = func(cur *TreeNode){
        if cur == nil {return}
        dfs(cur.Left)
        if pre != nil{
            pre.Right = cur
        } else {
            head = cur
        }
        cur.Left = pre
        pre = cur
        dfs(cur.Right)
    }
    dfs(root)
    return head
}
