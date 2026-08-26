(define typed-sum
  Acc N -> (if (= N 0) Acc (typed-sum (+ Acc N) (- N 1))))

(define typed-bool-loop
  N Acc -> (if (= N 0) Acc
              (if (number? N)
                  (typed-bool-loop (- N 1) true)
                  (typed-bool-loop (- N 1) Acc))))

(define typed-string-loop
  N S -> (if (= N 0) S (typed-string-loop (- N 1) (cn S "λ"))))

(define typed-list-loop
  N L -> (if (= N 0) L (typed-list-loop (- N 1) (cons N L))))

(define typed-vector-loop
  N V -> (if (= N 0) V
             (do (address-> V 0 N)
                 (typed-vector-loop (- N 1) V))))

(define typed-dynamic-apply
  F N -> (if (= N 0) 0 (do (F N) (typed-dynamic-apply F (- N 1)))))

(define typed-fallback-heavy
  N X -> (if (= N 0) X
             (typed-fallback-heavy (- N 1)
               (if (number? X) (+ X 1) (cn X "x")))))

(define typed-ir-bench
  -> (do
       (get-time run)
       (typed-sum 0 100000)
       (typed-bool-loop 8000 false)
       (typed-string-loop 200 "λ")
       (typed-list-loop 2000 [])
       (typed-vector-loop 8000 (absvector 1))
       (typed-dynamic-apply (lambda X (+ X 1)) 8000)
       (typed-fallback-heavy 2000 "x")
       (get-time run)))
