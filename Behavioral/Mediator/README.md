### Патерн Mediator відноситься до паттернів рівня об'єкта.

Патерн Mediator надає об'єкт-посередник, що приховує спосіб взаємодії безлічі інших об'єктів-колег.
Mediator робить систему слабопов'язаною, позбавляючи об'єкти від необхідності посилатися один на одного, що дозволяє змінювати взаємодію між ними незалежно.

Наприклад, у нас є посередник між заводом виробництва хлібобулочних виробів, фермером і магазином збуту.
Посередник позбавляє фермера від взаємодії з заводом, який використовує його сировину, а завод від взаємодії з магазином,
в який надходить продукція для збуту.

Для реалізації необхідно:
 1. Інтерфейс Mediator -- посередник описує організацію процесу з обміну інформацією між об'єктами типу Colleague;
 2. Структура ConcreteMediator, яка реалізує інтерфейс Mediator;
 3. Набір структур Colleague -- зокрема колега описує організацію процесу взаємодії об'єктів-колег з об'єктом типу Mediator;
 4. Структура ConcreteColleague -- реалізує методи Colleague. 
 Кожен об'єкт-колега знає тільки про об'єкт-медіатора. Всі об'єкти-колеги обмінюються інформацією тільки через посередника.