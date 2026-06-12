4. also implement the file parsers (in another package)
5. implement the buffered reader used in benchmark tests in the API.
6. refactor the tests.

---
per files piccoli leggi il testo e controlla se tutti i caratteri sono validi
per files grandi usa l'euristica per text recognition
 -> più thread in parallelo
 -> non leggere tutto il file
tipi di euristiche:
 -> %ascii vs %utf8
 -> %caratteri speciali
 -> se html riconosci l'header (solo all'inizio del file, altrimenti potrebbe essere uno snippet di codice)
se un file contiene sia HTML che TESTO
se un file contiene all'interno del codice (vedi blackbrid)