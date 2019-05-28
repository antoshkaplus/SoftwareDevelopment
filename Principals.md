
What he means by premature pessimisation, I think, is just the opposite of premature optimisation: a fundamental disregard of which data structures and algorithms to use.
Premature optimisation is often concerned with minute details of algorithms that can well be tweaked later and don’t need to be paid attention to at the beginning.
Premature pessimisation, by contrast, concerns the high-level design of code architecture: a fundamentally inefficient interface for your library for instance cannot be fixed later by optimising, since the public interface is pretty much cast in stone.

“Focus on one thing at a time: Prefer to give each entity (variable, class, function, namespace, module, library) one well-defined responsibility. As an entity grows, its scope of responsibility naturally increases, but its responsibility should not diverge.”

“KISS (Keep It Simple Software): Correct is better than fast. Simple is better than complex. Clear is better than cute. Safe is better than insecure”

“It is far, far easier to make a correct program fast than it is to make a fast program correct.”

“Strive for "shared-nothing;" prefer communication (e.g., message queues) over data sharing.”

“Where combinations of primitives can be arbitrary and you cannot capture the reasonable set of usage scenarios in one named operation, there are two alternatives: a) use a callback-based model (i.e., have the caller call a single member function, but pass in the task they want performed as a command or function object); or b) expose locking in the interface in some way.”

“Each level of nesting adds intellectual overhead when reading code because you need to maintain a mental stack”

“cyclic dependencies work against modularity and are a bane of large projects”

“Avoid membership fees: Where possible, prefer making functions nonmember nonfriends.”


# C++

“Destructors, deallocation, and swap never fail.” 
Destructors, memory deallocation function and swap shouldn’t raise exceptions! (terminate is going to be called)

“Virtual functions only "virtually" always behave virtually: Inside constructors and destructors, they don't. Worse, any direct or indirect call to an unimplemented pure virtual function from a constructor or destructor results in undefined behavior. If your design wants virtual dispatch into a derived class from a base class constructor or destructor, you need other techniques such as post-constructors.”

“Make base class destructors public and virtual, or protected and nonvirtual”

To solve warning: variable defined but never used - evaluate variable (void)x

namespace-level objects zero initialized

“In base classes, consider disabling the copy constructor and copy assignment operator, and instead supplying a virtual Clone member function if clients need to make polymorphic (complete, deep) copies.”

“Avoid making any assignment operator virtual”

“Avoid writing a copy assignment operator that relies on a check for self-assignment in order to work properly; often, that reveals a lack of error safety. Check only if self-assignment is frequent.”



# OOP

Avoid inheriting from concrete base classes. Need behavior - add non-member function. Need state - add member variable.

“Don't inherit publicly to reuse code (that exists in the base class); inherit publicly in order to be reused (by existing code that already uses base objects polymorphically)”

“A square "is-a" rectangle (mathematically) but a Square is not a Rectangle (behaviorally)”

“Nonvirtual Interface (NVI) pattern”

