#include <iostream>
using namespace std;

class Node {
public:
  int data;
  Node* next;
  Node* random;
  Node(int val): data(val), next(nullptr), random(nullptr) {}
};

// Copying a Linked List with Random Pointers:
Node* copyRandomList(Node* head){
  if(!head) return nullptr;

  // Step 1: Create a copy of each node and insert it right after the original node. 
  Node* current = head;
  while(current){
    Node* copy = new Node(current->data);
    copy->next = current->next;
    current->next = copy;
    current = copy->next;
  }

  // Step 2: Set the random pointers of the newly added nodes.
  current = head;
  while (current) {
    if(current->random){
      current->next->random = current->random->next;
    }
    current = current->next->next;  
  }

  // Step 3: Separate the newly added nodes to form the new copied list. 
  Node* newHead = head->next;
  current = head;
  while(current){
    Node* copy = current->next;
    current->next = copy->next;
    if(copy->next){
      copy->next = copy->next->next;
    }
    current = current->next;
  }
  return newHead;
}

void printList(Node* head){
  while(head){
    cout << "Data: " << head->data;
    if(head->random){
      cout << "Random: " << head->random->data;
    } else {
      cout << "Random: null";
    }
    cout << endl;
  }

}

int main (){
  // creating a linked list with random pointers 
  Node* head = new Node(1);
  head->next = new Node(2);
  head->next->next = new Node(3);
  head->next->next->next = new Node(4);
  head->next->next->next->next = new Node(5);

  // setting up random pointers 
  head->random = head->next->next; // 1's random points to 3 
  head->next->random = head;

}
