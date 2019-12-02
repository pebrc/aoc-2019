(ns day2
  [:require [clojure.string :as str]])

(def input (->>  (slurp "../inputs/day2_1")
                 (#(str/split % #","))
                 (map #(Integer/parseInt %))))

(defn aop [op p ip]
  (let [dst (aget p (+ 3 ip))
        x (aget p  (aget p (inc ip)))
        y (aget p (aget p (+ 2 ip)))]
    (aset p dst (op x y))))

(defn int-code [p ip]
  (let [op (aget p ip)]
    (case op
      99 (aget p 0)
      1 (do (aop + p ip )
            (recur p (+ 4 ip)))
      2 (do (aop * p ip)
            (recur p (+ 4 ip))))))

(defn with-input [xs n v]
  (let [p  (into-array Integer/TYPE xs)]
    (aset p 1 n)
    (aset p 2 v)
    p))

#_(=  2 (int-code (into-array Integer/TYPE [1 0 0 0 99]) 0))

#_(= 30 (int-code (into-array Integer/TYPE [1 1 1 4 99 5 6 0 99]) 0))

(def part1 (int-code (with-input input 12 2) 0))

(def part2 
  (reduce
   (fn [_ [n v]]
     (if (= 19690720 (int-code (with-input input n v) 0)) (reduced (+ (* 100 n) v))))
   0                                    
   (for [n (range 100)
         v (range 100)]
     [n v])))
