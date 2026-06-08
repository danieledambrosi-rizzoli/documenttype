1. Implement all the matching logic in the matchers folder (also change the "load.go" naming)
2. Finish implementing the API
3. polishing the API for matchers and types
4. also implement the file parsers (in another package)

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