# Deck Of Cards

![](https://i.imgur.com/oJybtFo.png) <br/>
Semplice Libreria riguardante la gestione basilare di un mazzo di carte.
Implementato attraverso i tutorial di [Gophercises](https://courses.calhoun.io).

---

Funzioni della libreria: <br/>
 + `New(*option*)`: Crea un nuovo mazzo. **(52 carte)** <br/> 
   + `New(Sort(Less))`: Crea un nuovo mazzo ordinato
   + `New(Joker(n))`: Crea un nuovo mazzo con *n* Joker
   + `New(Shuffle)`: Crea un nuovo mazzo e mischialo
   + `New(Filter)`: Crea un nuovo mazzo filtrato
        + Filter è la regola di filtraggio del mazzo 
          <br/> `es: Filter: card.Rank == 3` nel mazzo creato non ci saranno 3 di alcun seme <br/>
          `es: Filter: card.Suit == Spade` nel mazzo creato non ci saranno carte del seme di picche <br/>
     <p></p> <br/>
 + `Deck(n)`: Aggiunge al mazzo iniziale *n* mazzi uguali
   + **NB:** Con uguali si intende che mantengono eventuali *filtri* e *Joker* precedentemente inseriti

---

### Argomenti trattati con questo progetto:
1. Iota
2. Struct
3. Slices
4. Function
5. Print
   1. PrintF
6. Foreach
7. Test
   1. Test di una funzione non deterministica


---

NB: il fine più importante di questo progetto per me era di prendere più dimestichezza con il linguaggio GO attraverso esempi concreti
    e per questo vi sono *appunti* sotto forma di commenti all'interno del codice.

---



