/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var newHead *ListNode

    cur := head
    var prev *ListNode
    for cur != nil {
        tempNext := cur.Next

        if cur.Next == nil { // set Tail as new Head
            newHead = cur
            cur.Next = prev
            break
        }

        if prev == nil { // set Head as new Tail
            cur.Next = nil
        } else {
            cur.Next = prev
        }

        prev = cur
        cur = tempNext

    }

    return newHead
}
