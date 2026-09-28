# Smells to pattern refactorings

The refactorings Refactoring to Patterns offers for each smell. Pick the one that removes the cost you named, and prefer the smallest that does.

| Smell | Refactorings |
|---|---|
| Duplicated Code | Form Template Method, Introduce Polymorphic Creation with Factory Method, Chain Constructors, Replace One/Many Distinctions with Composite, Extract Composite, Unify Interfaces with Adapter, Introduce Null Object |
| Long Method | Compose Method, Move Accumulation to Collecting Parameter, Replace Conditional Dispatcher with Command, Move Accumulation to Visitor, Replace Conditional Logic with Strategy |
| Conditional Complexity | Replace Conditional Logic with Strategy, Move Embellishment to Decorator, Replace State-Altering Conditionals with State, Introduce Null Object |
| Primitive Obsession | Replace Type Code with Class, Replace State-Altering Conditionals with State, Replace Conditional Logic with Strategy, Replace Implicit Tree with Composite, Replace Implicit Language with Interpreter, Move Embellishment to Decorator, Encapsulate Composite with Builder |
| Indecent Exposure | Encapsulate Classes with Factory |
| Solution Sprawl | Move Creation Knowledge to Factory |
| Alternative Classes with Different Interfaces | Unify Interfaces with Adapter |
| Lazy Class | Inline Singleton |
| Large Class | Replace Conditional Dispatcher with Command, Replace State-Altering Conditionals with State, Replace Implicit Language with Interpreter |
| Switch Statements | Replace Conditional Dispatcher with Command, Move Accumulation to Visitor |
| Combinatorial Explosion | Replace Implicit Language with Interpreter |
| Oddball Solution | Unify Interfaces with Adapter |

A smell outside this table, such as Feature Envy, Message Chains, or Shotgun Surgery, takes its cure from Fowler's Refactoring: Move Function, Hide Delegate, Combine Functions into Class.

Smells in test code often take their cure from Meszaros's xUnit Test Patterns: duplicated setup becomes a shared fixture (Creation Method, Transaction Rollback Teardown).
