#include <iostream>
using namespace std;

class Node {
public:
  int data;
  Node* next;

  Node(int val) : data(val), next(nullptr) {}
};

// SinglyLinkedList class definition
class SinglyLinkedList {
private:
  Node* head;

public:
  SinglyLinkedList() : head(nullptr) {}

  // insertAtBeginning
  void insertAtBeginning(int data) {
    Node* newNode = new Node(data);
    newNode->next = head;
    head = newNode;
  }

// insertAtEnd
  void insertAtEnd(int val){
    Node* newNode = new Node(val);
    if(!head) {
      head = newNode;
      return;
    }
    Node* temp = head;
    while(temp->next) {
      temp = temp->next;
    }
    temp->next = newNode;
  }

  // detectingLoop 
  bool detectLoop(){
    Node* slow = head;
    Node* fast = head;
  
    while (slow && fast && fast->next) {
      slow = slow->next;
      fast = fast->next->next;
      if(slow == fast) {
        return true;
      }
    }
    return false;
  }

  // remove a loop
  void removeloop(Node* loopNode, Node* head){
    Node* ptr1 = head;
    Node* ptr2;
    while (1) {
      ptr2 = loopNode;
      while (ptr2->next == ptr1) {
        break;
      }
      ptr1 = ptr1->next;
    }
    ptr2->next = nullptr; // break the loop
  }

  // Find the middle
  Node* findMiddle(Node* head){
    if(!head) return nullptr;
    Node* slow = head;
    Node* fast = head;
    while(fast && fast->next) {
      slow = slow->next;
      fast = fast->next->next;
    }
    return slow; // this is middle node 
  }

  // display linked list
  void printList(){
    Node* temp = head;
    while(temp->next){
      cout << temp->data << " ";
      temp = temp->next;
    }
    cout << "nil" << endl;
  }

};


int main(){
  SinglyLinkedList list;

  // inserting element at beginning
  list.insertAtBeginning(21);
  list.insertAtBeginning(71);
  list.insertAtBeginning(7);
  list.insertAtBeginning(33);
  list.insertAtBeginning(21);

  // inserting element at end
  list.insertAtEnd(9);

  // print the list
  list.printList();

  // Loop checking
  if (list.detectLoop()) 
    cout <<  "Loop detected\n";
  else 
    cout << "No loop\n";

}

