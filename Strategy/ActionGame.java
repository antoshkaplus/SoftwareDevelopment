
// probably need a way to get to Character properties
// Character should have legs arms... how they are located 
interface WeaponBehavior {
    void useWeapon();
} 

class KnifeBevavior implements WeaponBehavior {}
class AxeBehavior implements WeaponBehavior {}
class BowAndArrowBehavior implements WeaponBehavior {}
class SwordBehavior implements WeaponBehavior {}

class Character {
    WeaponBehavior behavior;
    fight();
}

class King extends Character {}
class Queen extends Character {}
class Troll extends Character {}
class Knight extends Character {}
