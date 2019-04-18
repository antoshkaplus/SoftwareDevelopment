
// maybe interface ?
abstract class Beverage {
    String description;
    
    String getDescription() {}
    BigDecimal cost() {}
}


class HouseBlend extends Beverage {}
class DarkRoast extends Beverage {}
class Epresso extends {}
class Decaf extends {}

// just need interface
class Milk extends Beverage {
    Beverage beverage;
    // implementation
}

class Mocha extends Beverage {}
class Soy extends Beverage {}
class Whip extends Beverage {}