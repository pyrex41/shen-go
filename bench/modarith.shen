\\ Modular 32-bit word arithmetic written against the stdlib: an LCG through
\\ mod, rotations through floor and mod, and div. Before the native
\\ floor/ceiling/round/mod (cmd/shen/mathsnative.go) each mod or floor was a
\\ ~100 us digit loop. SHEN_NO_MATHS_NATIVE=1 runs the stdlib definitions.
\\ Run: ./shen script bench/modarith.shen

(define lcg X -> (mod (+ (* X 1664525) 1013904223) 4294967296))

(define lcg-run
  0 X -> X
  N X -> (lcg-run (- N 1) (lcg X)))

(define rotr
  X N -> (let P (power 2 N)
           (mod (+ (floor (/ X P)) (* (mod X P) (power 2 (- 32 N)))) 4294967296)))

(define rot-run
  0 X -> X
  N X -> (rot-run (- N 1) (+ (rotr X 7) (div X 3))))

(define run-bench
  Name Thunk -> (let T0 (get-time run)
                     R (thaw Thunk)
                     T1 (get-time run)
                  (do (output "~A: ~A s  result ~A~%" Name (- T1 T0) R) R)))

(run-bench "lcg x20000" (freeze (lcg-run 20000 12345)))
(run-bench "rotr x5000" (freeze (rot-run 5000 305419896)))
