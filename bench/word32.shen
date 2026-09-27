\\ 32-bit word arithmetic of the kind a pure-Shen SHA-256 / PRNG does: values
\\ between 2^25 and 2^33 that the fixnum range decides whether to box. Only
\\ + - and comparisons, so the stdlib's (pure-Shen, digit-loop) floor and mod
\\ do not dominate. Run: ./shen script bench/word32.shen

(define add32
  X Y -> (let S (+ X Y) (if (> S 4294967295) (- S 4294967296) S)))

\\ Additive lagged Fibonacci: X(n) = X(n-1) + X(n-2) mod 2^32.
(define fib32
  0 A B -> B
  N A B -> (fib32 (- N 1) B (add32 A B)))

(define fill
  V I X -> V where (> I (limit V))
  V I X -> (do (vector-> V I X) (fill V (+ I 1) (add32 X 2654435769))))

(define mix-round
  V I Acc -> Acc where (> I (limit V))
  V I Acc -> (let W (<-vector V I)
                  S (add32 (add32 W Acc) 2246822519)
                  _ (vector-> V I S)
               (mix-round V (+ I 1) (add32 Acc S))))

(define mix
  V 0 Acc -> Acc
  V N Acc -> (mix V (- N 1) (mix-round V 1 Acc)))

(define word-list
  0 X Acc -> Acc
  N X Acc -> (word-list (- N 1) (add32 X 3266489917) [X | Acc]))

(define sum-words
  [] Acc -> Acc
  [W | Ws] Acc -> (sum-words Ws (add32 Acc W)))

(define run-bench
  Name Thunk -> (let T0 (get-time run)
                     R (thaw Thunk)
                     T1 (get-time run)
                  (do (output "~A: ~A s  result ~A~%" Name (- T1 T0) R) R)))

(run-bench "fib32 x3000000" (freeze (fib32 3000000 1 1)))
(run-bench "mix 64 words x40000" (freeze (mix (fill (vector 64) 1 99) 40000 7)))
(run-bench "word list 1000000" (freeze (sum-words (word-list 1000000 5 []) 0)))
