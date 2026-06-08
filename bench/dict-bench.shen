(define fill
  D 0 -> D
  D N -> (let _ (shen.dict-> D (str N) N) (fill D (- N 1))))

(define lookups
  D 0 -> done
  D N -> (let _ (shen.<-dict D (str N)) (lookups D (- N 1))))

(define run-lookups
  D K 0 -> done
  D K Reps -> (let _ (lookups D K) (run-lookups D K (- Reps 1))))

(define run-bench
  Name Thunk ->
    (let T0 (get-time run)
         _  (Thunk)
         T1 (get-time run)
      (output "~A: ~A s~%" Name (- T1 T0))))

(output "~%--- dict benchmark ---~%")

(let D (shen.dict 256)
  (do (run-bench "fill 2000 keys" (freeze (fill D 2000)))
      (run-bench "20000 lookups (2000 keys x 10)" (freeze (run-lookups D 2000 10)))))

(output "--- done ---~%")
