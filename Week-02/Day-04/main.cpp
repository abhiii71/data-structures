#include <iostream>
using namespace std;

class DoublyLinkedList {
private:
  struct Node{
  int data;
  Node* next;
  Node* prev;
  Node(int val): data(val), next(nullptr), prev(nullptr) {}
  };
  Node* head;

public:
  DoublyLinkedList(): head(nullptr) {}
  ~DoublyLinkedList();

  void insertAtBeginning(int val);
  void insertAtEnd(int val);
  void insertAtPosition(int data, int position);
  void deleteNode(Node* node);
};


void DoublyLinkedList::insertAtBeginning(int val){
  Node* newNode = new Node(val);
  if (!head) {
    head = newNode;
  }
  if(head) {
    head->prev = newNode;
  }
    newNode->next = head;
    head = newNode;
}

void DoublyLinkedList::insertAtEnd(int val){
  Node* newNode = new Node(val);
  if(!head){
    head = newNode;
  }
  Node* temp = head;
  while(temp->next){
    temp = temp->next;
  }
  temp->next = newNode;
  newNode->prev = temp;
}

void DoublyLinkedList::insertAtPosition(int data, int position){
  Node* newNode = new Node(data);
  if (position == 0) {
    insertAtBeginning(data);
    return;
  } 
  Node* temp = head;
  for(int i = 0; temp && i < position-1; ++i){
    temp = temp->next;
  }
  if (!temp) return;
  newNode->next = temp->next;
  if(temp->next){
    temp->next->prev = newNode;
  }
  temp->next = newNode;
  newNode->prev = temp;
}

void DoublyLinkedList::deleteNode(Node* node){
  if(!head || !node) return;
  if(head == node){
    head = node->next;
  }
  if(node->next) {
    node->next->prev = node->next;
  }
  delete node;
}
