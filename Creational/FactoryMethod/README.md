### Патерн Factory Method відноситься до паттернів рівня класу і сфокусований лише на відносинах між класами.

Патерн Factory Method корисний, коли система повинна залишатися легко маштабованою шляхом додавання об'єктів нових 
типів. **Цей патерн є основою для всіх породжуючих патернів і може легко трансформуватися під потреби системи.** 
Тому, якщо перед розробником є не чіткі вимоги для продукту або незрозумілий спосіб організації взаємодії між 
продуктами, то для початку можна скористатися патерном Factory Method, поки повністю не сформуються всі вимоги.

Патерн Factory Method застосовується для створення об'єктів з певним інтерфейсом, реалізації якого надаються нащадками.
Іншими словами, є базовий абстрактний клас фабрики, який говорить, що кожна його дочірня фабрика(нащадок) повинна 
реалізувати такий-то метод для створення своїх продуктів.

Реалізація фабричного методу може бути різною, це може бути поліморфізм або параметризований метод.

Потрібно для реалізації:
 
 1. Базовий абстрактний клас **Creator**, що описує інтерфейс, який має реалізувати конкретна фабрика для виробництва 
 продуктів. Цей базовий клас описує фабричний метод.
 2. Базовий клас **Product**, що описує інтерфейс продукту, який повертає фабрика. Всі продукти, щл повертаються 
 фабрикою повинні дотримуватися єдиного інтерфейсу.
 3. Клас конкретної фабрики з виробництва продуктів **ConcreteCreator**. Цей клас повинен реалізувати фабричний метод;
 4. Клас реального продукту ConcreteProductA;
 5. Клас реального продукту ConcreteProductB;
 6. Клас реального продукту ConcreteProductC.
 
**FactoryMethod відрізняється від AbstractFactory**, тим, що AbstractFactory виробляє сімейство об'єктів, ці об'єкти різні, 
мають різні інтерфейсами, але взаємодіють між собою. У той час як Factory Method виробляє продукти дотримуються одного 
інтерфейсу і ці продукти не пов'язані між собою, не вступають у взаємодію.