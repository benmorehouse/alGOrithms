# B-Tree explanation

Thus far we have learned about how to access data that is already on the computer and at the ready. But often times, especially
in the industry, we have to access data that we have stored from a disk. Disk access takes a lot of expensive operations 
even when you have very simple pieces of data to access. On average, it can take somehwere between 5-10 seconds for even
small bits of data. The idea behind a b-tree is that we can create a very complex program that, though isn't neccessarily effecient
in coding up, will cut down the amount of reads/writes to the disk.

The idea is intuitive: Binary Search Trees are not efficent because best case scenario we get o(log n) runtime. But if we
have less height than a BST, then we lower the amount of access to the disk.

# Implementation

We must make sure to keep this tree like structure from turning into a binary search tree, otherwise we are then risking
a linkedlist access.

# Operations

Addition is quite the interesting technique. Our goal is to have as cheap of insertions onto the disk as possible.
