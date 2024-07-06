#include <iostream>
using namespace std;

class Node {
public:
  int data;
  Node* next;

  Node(int val) : data(val), next(nullptr) {}
};

class SinglyLinkedList {
private:
  Node* head;

public:
  SinglyLinkedList() : head(nullptr) {}

  void insertAtEnd(int val){ 
    Node* newNode = new Node(val);
    if(!head) {
      head = newNode;
      return;
    }
    Node* temp = head;
    while(temp->next){
      temp = temp->next;
    }
    temp->next = newNode;
  }

  // DeletefromBeginning
  void deleteFromBeginning(){
    if(!head) return;
    Node* temp = head;
    head = head->next;
    delete temp;
  }
 
// deletefromEnd:
  void deleteFromEnd(){
    if(!head) return;
    if(!head->next){
      delete head;
      head= nullptr;
      return;
    }
    Node* temp = head;
    while(temp->next && temp->next->next){
      temp = temp->next;
    }
    delete temp->next;
    temp->next = nullptr;
  }

  // deleteFromPosition
  void deleteFromPosition(int position){
    if (position == 0){
      deleteFromBeginning();
      return;
    }
    Node* temp = head;
      for(int i = 0; i < position-1 && temp; ++i){
      temp = temp->next;
    }
    if (!temp || !temp->next) return;
    Node* nodeToDelete = temp->next;
    temp->next = temp->next->next;
    delete nodeToDelete;
  }

  // PrintList:
  void printList(){
    Node *temp = head;
    while(temp) {
      cout << temp->data << " ";
      temp = temp->next;
    }
    cout << endl;
  }
};

int main (){
  SinglyLinkedList list;
  
  list.insertAtEnd(1);
  list.insertAtEnd(2);
  list.insertAtEnd(3);

  cout << " Initial List: ";
  list.printList();

  list.deleteFromBeginning();
  cout << "After deleting from beginning: ";
  list.printList();

  list.deleteFromEnd();
  cout << "After deleting from end: ";
  list.printList();

  list.deleteFromPosition(1);
  cout << "After deleting from position 1: ";
  list.printList();
  
  return 0;
}
