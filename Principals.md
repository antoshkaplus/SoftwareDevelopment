
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

Avoid inheriting from concrete base classes. Need behavior - add non-member function. Need state - add member variable.

“Don't inherit publicly to reuse code (that exists in the base class); inherit publicly in order to be reused (by existing code that already uses base objects polymorphically)”

“A square "is-a" rectangle (mathematically) but a Square is not a Rectangle (behaviorally)”

“Nonvirtual Interface (NVI) pattern”

