#include <stdio.h>
#include <stdlib.h>
// some help for syntax, etc from here:
// https://www.geeksforgeeks.org/nth-node-from-the-end-of-a-linked-list/

// make a struct to build the linked list
struct Node {
  int data;
  struct Node* next;  // has a pointer to the next node
};

void printList(struct Node* a) {
  while (a != NULL)
  {
    printf("%d ", a->data);
    a = a->next;
  }
  printf("\n-------------\n");
}

void printMergedLists(struct Node* listA, struct Node* listB)
{
  printf("\nPrinting merged lists:\n-------------\n");
  struct Node *currA = listA;
  struct Node *currB = listB;

  while ((currA != NULL) && (currB != NULL))
  {
    if (currA->data < currB->data) {
      printf ("%d ", currA->data);
      currA = currA->next;
    }
    else
    {
      printf ("%d ", currB->data);
      currB = currB->next;
    }
  }

  if (currA == NULL)
  {
    printList(currB);
  }
  else
  {
    printList(currA);
  }

}

void addToList(struct Node** head_ref, int new_data)
{
  struct Node* new_node =
    (struct Node*) malloc(sizeof(struct Node));

  new_node->data = new_data;
  new_node->next = (*head_ref);
  // assigning to a pointer means the value changes in the caller as well
  (*head_ref) = new_node;
}
int main()
{
  struct Node* headA = NULL;
  struct Node* headB = NULL;
  // make list A
  addToList(&headA, 20);
  addToList(&headA, 15);
  addToList(&headA, 10);
  addToList(&headA, 5);
  // make list B
  addToList(&headB, 17);
  addToList(&headB, 13);
  addToList(&headB, 8);
  addToList(&headB, 2);

  printMergedLists(headA, headB);

  // base case: empty List
  struct Node* headEmpty = NULL;
  printMergedLists(headA, headEmpty);

  return 0;
}
