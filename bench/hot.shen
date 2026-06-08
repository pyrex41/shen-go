(define tak
  X Y Z ->
    (if (not (< Y X))
        Z
        (tak (tak (- X 1) Y Z)
             (tak (- Y 1) Z X)
             (tak (- Z 1) X Y))))

(define fib
  N -> N           where (< N 2)
  N -> (+ (fib (- N 1)) (fib (- N 2))))

(define bench-tak
  X Y Z 0 -> done
  X Y Z N -> (let _ (tak X Y Z) (bench-tak X Y Z (- N 1))))
