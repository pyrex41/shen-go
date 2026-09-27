\\ Breadth-first state-space search in the style of a model checker
\\ (issue #57): states are lists, the visited set is keyed by them. bfs
\\ uses the shen.x.map extension, bfs-portable put/get on a bucket vector.
\\ Run: ./shen script bench/bfs.shen

(define next-states
  [P L C] -> (filter (/. S (valid? S)) [[(+ P 1) L C] [P (+ L 1) C] [P L (+ C 1)] [(- P 1) (+ L 1) C]]))

(define valid? [P L C] -> (and (>= P 0) (and (<= P 14) (and (<= L 14) (<= C 14)))))

(define bfs
  [] _ N -> N
  [S | Q] Seen N -> (let Ns (next-states S)
                         New (filter (/. X (not (shen.x.map-has? Seen X))) Ns)
                         _ (map (/. X (shen.x.map-put Seen X true)) New)
                         (bfs (append New Q) Seen (+ N 1))))

(define seen? X V -> (trap-error (get X seen V) (/. E false)))

(define bfs-portable
  [] _ N -> N
  [S | Q] Seen N -> (let Ns (next-states S)
                         New (filter (/. X (not (seen? X Seen))) Ns)
                         _ (map (/. X (put X seen true Seen)) New)
                         (bfs-portable (append New Q) Seen (+ N 1))))

(define repeat 0 _ -> done N F -> (do (thaw F) (repeat (- N 1) F)))

(output "states: ~A ~A~%" (bfs [[0 0 0]] (shen.x.map-new) 0) (bfs-portable [[0 0 0]] (vector 4096) 0))
(output "shen.x.map, 20 searches:~%")
(time (repeat 20 (freeze (bfs [[0 0 0]] (shen.x.map-new) 0))))
(output "put/get on a vector, 20 searches:~%")
(time (repeat 20 (freeze (bfs-portable [[0 0 0]] (vector 4096) 0))))
