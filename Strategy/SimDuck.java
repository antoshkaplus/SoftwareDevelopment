
// probably interfaces have duck instance
// and they are able to ask particular 
// properties to change

interface FlyBehavior {
    void fly();
}

class FlyWithWings implements FlyBehavior {}
class FlyNoWay implements FlyBehavior {}


interface QuackBehavior {
    void quack();
}

class Quack implements QuackBehavior {}
class Squeak implements QuackBehavior {}
class MuteQuack implements QuackBehavior {}


// every duck should have beak, wings, and clutches
// or we can't really manipulate them with strategy
// behavior manipulates moves of common duck
abstract class Duck {
    FlyBehavior flyBehavior;
    QuackBehavior quackBehavior;
    
    // will change parameters of duck 
    void performQuack();
    void performFly();
    
    // drawing based on duck parameters
    void display()

    // can dynamically change behaviors
    void setFlyBehavior(FlyBehavior flyBehavior) {}
    void setQuackBehavior(QuackBehavior quackBehavior) {}
}

// new type of duck with more parameters
// if need consider more parameters in behavior - 
// create more behaviors with new reference of duck inside



