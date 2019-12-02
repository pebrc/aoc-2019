(ns day1
  [:require [clojure.java.io :as io]])


(defn calc-fuel [i]
  (let [cost (- (Math/floor (/ i  3)) 2)]
    (if (pos? cost) cost 0)))

(defn calc-fuel-recur [s i]
  (let [f (calc-fuel i)]
    (if (pos? f)
      (recur (+ s f) f)
      s)))

(def problem1
  (with-open [rdr (io/reader "../inputs/day1_1")]
    (->> (line-seq rdr)
         (map #(Integer/parseInt %))
         (map calc-fuel)
         (reduce + 0))))

(def problem2
  (with-open [rdr (io/reader "../inputs/day1_1")]
    (->> (line-seq rdr)
         (map #(Integer/parseInt %))
         (map #(calc-fuel-recur 0 %))
         (reduce + 0))))



