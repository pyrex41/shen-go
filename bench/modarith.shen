\\ Integer arithmetic written against the stdlib Maths package: an LCG through
\\ mod, rotations through floor/mod/power/div, and gcd, lcd, isqrt and power.
\\ Before the natives in cmd/shen/mathsnative.go each of these was a loop in
\\ Shen (floor a ~100 us digit loop, gcd/lcd linear in the smaller argument,
\\ isqrt linear in the root). SHEN_NO_MATHS_NATIVE=1 runs the stdlib
\\ definitions.
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

(define gcd-run
  0 Acc -> Acc
  N Acc -> (gcd-run (- N 1) (+ Acc (gcd (+ 200000 N) (* 3 (+ 70000 N))))))

(define lcd-run
  0 Acc -> Acc
  N Acc -> (lcd-run (- N 1) (+ Acc (lcd (+ 100001 (* 2 N)) (+ 300003 (* 6 N))))))

(define isqrt-run
  0 Acc -> Acc
  N Acc -> (isqrt-run (- N 1) (+ Acc (isqrt (* N 1000003)))))

(define power-run
  0 Acc -> Acc
  N Acc -> (power-run (- N 1) (+ Acc (power 1.0001 (+ 100 N)))))

(run-bench "lcg x20000" (freeze (lcg-run 20000 12345)))
(run-bench "rotr x5000" (freeze (rot-run 5000 305419896)))
(run-bench "gcd x200" (freeze (gcd-run 200 0)))
(run-bench "lcd x200" (freeze (lcd-run 200 0)))
(run-bench "isqrt x200" (freeze (isqrt-run 200 0)))
(run-bench "power x2000" (freeze (power-run 2000 0)))
