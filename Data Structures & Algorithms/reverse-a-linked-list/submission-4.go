/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    // Base cases: empty list or single node
    if head == nil || head.Next == nil {
        return head
    }
    
    // Reverse the rest of the list recursively
    newHead := reverseList(head.Next)
    
    // Adjust pointers
    head.Next.Next = head  // Make next node point back to current
    head.Next = nil        // Break the forward link
    
    return newHead
}