(ns day4)

(def input [158126 624574])

(defn digits [n]
  (if (pos? n)
    (conj (digits (quot n 10)) (mod n 10) )
    []))

(defn pw-rules [d-op n]
  (let [d (digits n)
        sorted (= d (sort d))
        dup (->> (partition-by identity d)
                 (map count)
                 (some #(d-op 2 %)))]
    (and sorted dup)))

(defn solve [rules]
  (->> (range (first input) (inc (last input)))
       (filter rules)
       count))

(comment
  (solve (partial pw-rules <=)) ;part1
  (solve (partial pw-rules =)) ;part2
)
