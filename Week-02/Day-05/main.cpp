#include <iostream>
#include <iterator>
using namespace std;

class CircularSinglyLinkedList {
private:
  struct Node {
    int data;
    Node* next;
    Node(int data): data(data), next(nullptr) {}
  };
  Node* head;
public:
  CircularSinglyLinkedList(): head(nullptr){}
  ~CircularSinglyLinkedList();

  void insertAtBeginning(int data);
  void insertAtEnd(int data);
  void insertAtPosition(int data, int position);
  void printList();
  void deleteNode(int key);
};

CircularSinglyLinkedList::~CircularSinglyLinkedList(){
if(!head) return;
  Node* current = head;
  Node* nextNode;
  do{
    nextNode = current->next;
    delete current;
    current = nextNode;
  } while (current != head);
}

// Doubly linked list 

class CircularDoublyLinkedList{
private:
  struct Node {
    int data;
    Node* next;
    Node* prev;
    Node(int data): data(data), next(nullptr), prev(nullptr) {}
  };
    Node* head;

public:
  CircularDoublyLinkedList(): head(nullptr) {}
  ~CircularDoublyLinkedList();

  void insertAtBeginning(int data);
  void insertAtEnd(int data);
  void insertAtPosition(int position, int data);
  void deleteNode(int key);
  void printList();
  
};

CircularDoublyLinkedList::~CircularDoublyLinkedList() {
  if(!head) return;
  Node* current = head;
  do{
    Node* nextNode = current->next;
    delete current;
    current = nextNode;
  } while(current != head);
}

void CircularSinglyLinkedList::insertAtBeginning(int data){
  Node* newNode = new Node(data);
  if(!head){
    head = newNode;
    newNode->next = head;
  }else {
    Node* temp = head;
    while (temp->next != head) {
      temp = temp->next;
    }
    newNode->next = head;
    temp->next = newNode;
    head = newNode;
  }
}

void CircularSinglyLinkedList::insertAtEnd(int data){
  Node* newNode = new Node(data);
  if (!head) {
    head = newNode;
    newNode->next = head;
  } else {
    Node* temp = head;
    while (temp->next != head) {
      temp = temp->next;
    }
    temp->next = newNode;
    newNode->next = head;
  }
}

// insertAtPosition
void CircularSinglyLinkedList::insertAtPosition(int data, int position){
  if(position == 0) {
    insertAtBeginning(data);
    return;
  }
  Node* newNode = new Node(data);
  Node* temp = head;
  for (int i = 0; i < position - 1 && temp->next != head; ++i){
    temp = temp->next;
  }
  if(temp->next == head){
    insertAtEnd(data);
  } else {
    newNode->next = temp->next;
    temp->next = newNode;
  }
}

// deleteNode
void CircularSinglyLinkedList::deleteNode(int key){
  if(!head) return;
  Node* current = head;  
  Node* prev = nullptr;
  do{
    if(current->data == key){
      if (prev) {
      prev->next = current->next;
      if(current == head) head = current->next;
    } else {
      Node* temp = head;
      while (temp->next != head) temp = temp->next;
      head = head->next;
      temp->next = head;
      }
    delete current;
    return;
    }
    prev = current;
    current = current->next;
    } while(current != head);

}

// printList
void CircularSinglyLinkedList::printList() {
  if (!head) return;
  Node* temp = head;
  do {
    cout << temp->data << " ";
    temp = temp->next;
  } while (temp != head);
  cout << endl;
}


// InsertAtBeginning
void CircularDoublyLinkedList:: insertAtBeginning(int data){
    Node* newNode = new Node(data);
    if(!head){
      head = newNode;
      newNode->next = head;
      newNode->prev = head;
    } else {
      Node* temp = head;
      while (temp->next != head) {
        temp = temp->next;
      }
      newNode->next = head;
      newNode->prev = temp;
      head->prev = newNode;
      temp->next = newNode;
      head = newNode;
    }
}

// InsertAtEnd
void CircularDoublyLinkedList:: insertAtEnd(int data){
  Node* newNode = new Node(data);
    if(!head) {
      head = newNode;
      newNode->next = head;
      newNode->prev = head;
    } else {
      Node* temp = head;
      while (temp->next != head) {
        temp = temp->next;
      }
      temp->next = newNode;
      newNode->next = head;
      newNode->prev = temp;
      head->prev = newNode;
    }
}

// insertAtPosition
void CircularDoublyLinkedList::insertAtPosition(int data, int position){
  if (position == 0){
    insertAtBeginning(data);
    return;
  } 

  Node* newNode = new Node(data);
  Node* temp = head;
  for(int i = 0; i < position-1 && temp->next != head; ++i){
    temp = temp->next;
  }

  if(temp->next == head){
    insertAtEnd(data);
  } else {
    newNode->next = temp->next;
    temp->next->prev = newNode;
    temp->next = newNode;
    newNode->prev = temp;
  }
}

//deleteNode
void CircularDoublyLinkedList::deleteNode(int key){
  if(!head) return;
  Node* current = head;
  do{
    if(current->data == key) {
      if(current->prev){
        current->prev->next = current->next;
        current->next->prev = current->prev;
      } else {
        Node* temp = head;
        while (temp->next != head) temp = temp->next;
        head = head->next;
        head->prev = temp;
        temp->next = head;
      }
      delete current;
      return;
    }
    current = current->next;
  } while (current != head);
}


void CircularDoublyLinkedList::printList() {
  if (!head) return;
  Node* temp = head;
  do {
    cout << temp->data << " ";
    temp = temp->next;
  } while (temp != head);
  cout << endl;
}

int main () {
  // circularSinglyList
  cout << " Circular Singly Linked List: \n";
  CircularSinglyLinkedList listS;
  listS.insertAtEnd(21);
  listS.insertAtBeginning(71);
  listS.insertAtPosition(2,2);
  listS.printList();
  listS.deleteNode(2);
  listS.printList();

  // Doubly
  cout << " Circular Doubly Linked List: \n";
  CircularDoublyLinkedList listD;
  listD.insertAtBeginning(54);
  listD.insertAtBeginning(18);
  listD.insertAtPosition(2,5);
  listD.insertAtBeginning(23);
  listD.insertAtEnd(7);
  listD.printList();
  listD.deleteNode(2);
  listD.printList();



  return 0;
}











