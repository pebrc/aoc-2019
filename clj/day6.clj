(ns day6
  [:require
   [clojure.java.io :as io]
   [clojure.string :as str]])


(defn orbit-map [xs]
  (reduce (fn [acc p]
            (-> (assoc acc :obj (first p) )
                (update :next conj {:obj (second p) :next []})))
          {:next []}
          xs))


(defn uom [cur d]
  (let [next (map #(uom (get orbits (:obj %)) (inc d)) (:next cur))]
    (assoc cur :next next :depth d)))


(def pairs (->> (line-seq (io/reader "../inputs/day6_1"))
                (sort)
                (map #(str/split % #"\)"))))

(def orbits (->> pairs
                 (partition-by first)
                 (map orbit-map)
                 (reduce (fn [acc o] (assoc acc (:obj o) o)) {})))

(def uom-graph (uom (get orbits "COM") 0))


(defn count-orbits [uom]
  (+ (:depth uom) (reduce +  (map count-orbits (:next uom)))))

#_(count-orbits uom-graph)

