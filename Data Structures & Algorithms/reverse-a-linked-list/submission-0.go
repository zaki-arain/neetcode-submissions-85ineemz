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

func setNext(node *ListNode, next *ListNode) *ListNode {
    node.Next = next
    return node
}

// 1, 2, 3, 4, 5

// 1 -> 2 -> 3 -> 4 -> 5

// loop, start with cur
// if prev is nil
    // tempNext = cur.Next
    // set 1 -> nil (cur.Next = nil)
    // set cur = 1 (cur = tempNext)
    // set 1 = prev (prev = cur)
// set prev = 1 (prev = cur)

// 2
// 3
