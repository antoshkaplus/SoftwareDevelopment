1. Hello World program. 

2. Inheritance:
```python
class Animal:
    def speak():
        ...

class Dog(Animal):
    def speak():
        print('Guf')

class Cat(Animal):
    def speak():
        print('Myau')

def main():
    Animal dog = Dog()
    dog.speak()

    Animal cat = Cat()
    cat.speak()
```
Should be able to debug by placing breakpoints and also 
traversing line by line.

3. Sqlite inmemory db. Store `Fruit` there.
```python
import datetime 

class Fruit:
    name: str
    created: datetime.datetime
```
Check that add, remove, list work.

4. Expanding (3) we want a command line client to 
    perform actions on `Fruit`.

5. Expanding (3) and (4) we want to split the app on `Client` and `Server`.
    `Client` should still use cli for input. Use any available TCP method for 
    message transfer between them. Can also try some unique methods for 
    message transfer separately.

Adding unit tests and functional tests for above solutions would be nice.  