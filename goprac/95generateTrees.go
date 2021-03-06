/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*

the idea is that you can divide into further subproblems
let F(n) := be the set of all unique bst
let G(i,n) := be the set of all unique bst with root node i out on n nodes
F(n) = sum of G(i,n) from i = 0 to n

so for root i , all the lefttree nodes is with bsts with roots in [0,...,i-1] and righttree nodes with roots in [i+1,...,n]

so let basically G(i,n) := [unique bst with root i with with roots in [0,...,i-1] and righttree nodes with roots in [i+1,...,n]]

lets make a function called generateTreesFromSet (gt in diagram) that takes in a set of unqiue values and return a set of unique bst with those values
so visually, and gt will be recursive to recursive generate subbranches

              i
              *
             * *  
            *   *
   gt[0:i-1]    gt[i+1:n]

*/
func generateTrees(n int) []*TreeNode {
    if n <= 0 {
        return []*TreeNode{}
    }
    return generateTreesFromSet(1,n+1)
}

func generateTreesFromSet(start, end int) []*TreeNode {
    if  end-start == 0 {
        return []*TreeNode{nil}
    }
    if end-start == 1 {
        return []*TreeNode{&TreeNode{start,nil,nil}}
    }
    arrOfTrees := []*TreeNode{}
    for i := start; i < end; i++ {
        leftTrees := generateTreesFromSet(start,i)
        rightTrees := generateTreesFromSet(i+1,end)
        root := i
        for k,_ := range leftTrees {
            for l,_ := range rightTrees{
                arrOfTrees = append(arrOfTrees, &TreeNode{root,leftTrees[k],rightTrees[l]})
            }
        }
    }
    return arrOfTrees
}
