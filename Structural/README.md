## Структурні патерни (Structural)

Асоціація - це відношення, коли об'єкти двох класів можуть посилатися один на одного. Наприклад, параметр класу 
містить екземпляр іншого класу.

Агрегація - приватна форма асоціації. Агрегація застосовується,
коли один об'єкт повинен бути контейнером для інших об'єктів і час існування цих 
об'єктів ніяк не залежить від часу існування об'єкта-контейнера. В основному, якщо 
контейнер буде знищений, то що входять в нього об'єкти не постраждають. Наприклад, 
ми створили об'єкт, а потім передали його в об'єкт-контейнер, будь-яким чином, 
можна в метод об'єкта-контейнера передати або привласнити відразу параметр 
контейнера ззовні. Значить, при видаленні контейнера ми ніяк не торкнемося 
нашого створеного об'єкту, який може взаємодіяти і з іншими контейнерами.

Композиція - Теж саме, що і агрегація, але складові об'єкти не можуть існувати окремо
від об'єкта контейнера і якщо контейнер буде знищений, то весь його вміст буде знищено
теж. Наприклад, ми створили об'єкт в методі об'єкту-контейнера і привласнили його
параметри об'єкту-контейнеру. Із зовні, про наш створений об'єкті ніхто не знає,
значить, при видаленні контейнера, створений об'єкт буде видалений, тому що на нього
немає посилання ззовні.

До патернів рівня класу відноситься тільки «Адаптер». Сенс його роботи в тому, що якщо 
у вас є клас і його інтерфейс не сумісний з бібліотеками вашої системи, то що б вирішити 
цей конфлікт, ми не змінюємо код цього класу, а пишемо для нього адаптер.

Всі структурні патерни відповідають за створення правильної структури системи, в якій 
без зусиль зможуть взаємодіяти між собою вже наявні класи і об'єкти.

* [Адаптер (Adapter)](Adapter) 
* [Міст (Bridge)](Bridge) 
* [Композиція (Composite)](Composite) 
* [Декоратор (Decorator)](Decorator) 
* [Фасад (Facade)](Facade) 
* [Легковажний (Flyweight)](Flyweight) 
* [Замісник (Proxy)](Proxy)
 
