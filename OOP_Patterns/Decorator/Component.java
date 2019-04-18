
abstract class Component {
    methodA() {}
    methodB() {}
}

class ConcreteComponentA extends Component {}
class ConcreteComponentB extends Component {}

// component has basic functionality
// concrete component extends functonality

// decorator adds some functionality on top
// good thing is that you can reuse that functionality
// for other concrete components
// the only problem is that we have only basic interface to use

// existing base concrete decorator ?

// is good for processing, filtering

class ConcreteDecoratorA extends Component {
    Component component;
    // concrete implementation of methods
}
class ConcreteDecoratorB extends Component {
    Component component;
    // implementation
}
