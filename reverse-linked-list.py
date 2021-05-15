## a -> b -> c -> d
## d -> c -> b -> a

def reverseLinkedList(head):
    next_node = None
    curr_node = head
    previous_node = None
    while curr_node:
        next_node = curr_node.next
        
        curr_node.next = previous_node
        
        previous_node = curr_node
        
        curr_node = next_node

    return previous_node


class LinkedList():
    def __init__(self, value):
        self.value = value
        self.next = None


if __name__ == '__main__':
    first  = LinkedList(1)
    second = LinkedList(2)
    third  = LinkedList(3)
    fourth = LinkedList(4)

    first.next = second
    
    second.next = third

    third.next = fourth

    print(first.value)
    print(first.next.value)
    print(first.next.next.value)
    print(first.next.next.next.value)

    reversed = reverseLinkedList(first)

    print(reversed.value)
    print(reversed.next.value)
    print(reversed.next.next.value)
    print(reversed.next.next.next.value)

