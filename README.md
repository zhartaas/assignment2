# assignment2

Single Responsibility Principle (SRP):
Each class does just one job. For example, the Librarian class only deals with librarian tasks like adding books and magazines. Similarly, the Patron class only handles tasks related to patrons, such as borrowing items.

Open/Closed Principle (OCP):
The system can be extended without changing existing code. If we want to add new types of items like CDs, we can do so without having to edit the existing code. This is possible because we define interfaces that specify what an item should do, and new item types just need to follow those rules.

Liskov Substitution Principle (LSP):
We can swap a subclass with its parent class without causing problems. So, if a function expects any item from the library, it will work fine whether it gets a book or a magazine because they both behave as library items.

Interface Segregation Principle (ISP):
Interfaces are tailored to what the user needs. In this case, the LibraryItem interface only includes what's necessary for users, like displaying an item. It doesn't include anything extra that users don't need.

Dependency Inversion Principle (DIP):
High-level parts of the code don't rely on specific low-level details. For example, the librarian and patron classes don't depend on the specific details of books or magazines. Instead, they rely on a more general idea of what items in the library can do.
In simpler terms, the code is set up so that it's easy to add new features without breaking what's already there. Each part of the code has a clear job, and everything works together smoothly
