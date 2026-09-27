\\ Bit-list arithmetic in the style of a pure-Shen SHA-256 (issue #32):
\\ 32-bit words as lists of bits, xor/and/not/rotate/add by recursion over
\\ the lists, so nearly every step conses. Run: ./shen script bench/bitlist.shen
(define n->bits
  0 0 Acc -> Acc
  K N Acc -> (n->bits (- K 1) (shen.div2 N) [(shen.odd N) | Acc]))
(define shen.odd N -> (if (= N (* 2 (shen.div2 N))) 0 1))
(define shen.div2
  N -> (shen.div2h N 0 1073741824))
(define shen.div2h
  N Q 0 -> Q
  N Q P -> (if (>= N (* 2 P)) (shen.div2h (- N (* 2 P)) (+ Q P) (shen.half P)) (shen.div2h N Q (shen.half P))))
(define shen.half
  1 -> 0
  P -> (/ P 2))
(define word N -> (n->bits 32 N []))
(define bxor
  [] [] -> []
  [A | As] [B | Bs] -> [(if (= A B) 0 1) | (bxor As Bs)])
(define band
  [] [] -> []
  [A | As] [B | Bs] -> [(if (and (= A 1) (= B 1)) 1 0) | (band As Bs)])
(define bnot
  [] -> []
  [A | As] -> [(- 1 A) | (bnot As)])
(define rotr
  0 W -> W
  N W -> (rotr (- N 1) (let R (reverse W) [(hd R) | (reverse (tl R))])))
(define add-bits
  As Bs -> (let R (add-h (reverse As) (reverse Bs) 0 []) R))
(define add-h
  [] [] _ Acc -> Acc
  [A | As] [B | Bs] C Acc -> (let S (+ A (+ B C))
                                  (add-h As Bs (if (> S 1) 1 0) [(if (= 1 (shen.odd S)) 1 0) | Acc])))
(define ch X Y Z -> (bxor (band X Y) (band (bnot X) Z)))
(define sigma X -> (bxor (bxor (rotr 2 X) (rotr 13 X)) (rotr 22 X)))
(define rounds
  0 A B C -> A
  N A B C -> (rounds (- N 1) (add-bits (sigma A) (ch A B C)) A B))
(let W1 (word 1779033703) W2 (word 3144134277) W3 (word 1013904242)
  (time (rounds 3000 W1 W2 W3)))
