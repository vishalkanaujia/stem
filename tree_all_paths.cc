#include <iostream>

using namespace std;

class node {
  public:
      int data;
      node *left;
      node *right;
};


int getHeight(node *root)
{
    if (root == NULL) {
        return -1;
    }

    int lheight = getHeight(root->left);
    int rheight = getHeight(root->right);

    return(max(lheight, rheight) + 1);
}


