\\ Hash-table benchmark. The S42 kernel has no shen.dict (S41's dict.kl is
\\ gone), so this measures what the kernel does have: put/get on a vector
\\ of buckets, which goes through the native hash; and the shen.x.map host
\\ extension (issue #57), keyed by any Shen value with =.
\\ Run: ./shen script bench/dict-bench.shen

(define fill
  _ 0 -> done
  V N -> (do (put (str N) bench N V) (fill V (- N 1))))

(define lookups
  _ 0 -> done
  V N -> (do (get (str N) bench V) (lookups V (- N 1))))

(define map-fill
  _ 0 -> done
  M N -> (do (shen.x.map-put M [N (* 2 N)] N) (map-fill M (- N 1))))

(define map-lookups
  _ 0 -> done
  M N -> (do (shen.x.map-get M [N (* 2 N)] 0) (map-lookups M (- N 1))))

(define repeat
  _ 0 -> done
  F N -> (do (thaw F) (repeat F (- N 1))))

(define run-bench
  Name Thunk ->
    (let T0 (get-time run)
         _  (thaw Thunk)
         T1 (get-time run)
      (output "~A: ~A s~%" Name (- T1 T0))))

(output "~%--- hash-table benchmark ---~%")

(let V (vector 256)
  (do (run-bench "put 2000 string keys" (freeze (fill V 2000)))
      (run-bench "20000 get (2000 keys x 10)" (freeze (repeat (freeze (lookups V 2000)) 10)))))

(let M (shen.x.map-new)
  (do (run-bench "shen.x.map-put 2000 list keys" (freeze (map-fill M 2000)))
      (run-bench "20000 shen.x.map-get (2000 keys x 10)" (freeze (repeat (freeze (map-lookups M 2000)) 10)))))

(output "--- done ---~%")
