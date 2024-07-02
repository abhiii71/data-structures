#include <iostream>
using namespace std;

struct Node {
  int data;
  Node* next;
  Node(int val) : data(val), next(nullptr) {}
};

class SinglyLinkedList{
public:
  Node* head;
  SinglyLinkedList(): head(nullptr) {}

  void insertAtBeginning(int data);
  void insertAtEnd(int data);
  void insertAtPosition(int data, int position);
};

void SinglyLinkedList:: insertAtBeginning(int data){
  Node* newNode = new Node(data);
  newNode -> next = head;
  head = newNode; 
}

void SinglyLinkedList:: insertAtEnd(int data){
  Node* newNode = new Node(data);
  if (!head) {
    head = newNode;
    return;
  }
  Node* temp = head;
  while(temp->next){
    temp = temp->next; 
  }
  temp -> next = newNode;
}

void SinglyLinkedList:: insertAtPosition(int data, int position){
  if (position == 0){
    insertAtBeginning(data);
    return;
  }
  Node* newNode = new Node(data);
  Node* temp = head;
  for(int i = 0; i < position - 1 && temp; ++i){
    temp = temp->next;
  }
  if(!temp) return;
  newNode->next = temp->next;
  temp->next = newNode; 
}

int main(){
  SinglyLinkedList list;
  list.insertAtBeginning(10);


  Node* temp = list.head;
  while(temp) {
    cout << temp-> data << " ";
    temp = temp -> next;
  }
  cout << endl;
  return 0;
}

