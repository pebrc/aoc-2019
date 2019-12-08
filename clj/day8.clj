(ns day8
  [:require [clojure.string :refer [join]]])

(def width 25)
(def height 6)

(def layer-size (* width height))

(def input (->> (seq (slurp "../inputs/day8_1"))
     (map #(Integer/parseInt (str %)))
     (partition layer-size)))

(def part1 (->> input
                (map frequencies)
                (apply min-key #(get % 0))
                ((juxt #(get % 1) #(get % 2)))
                (reduce *)))

(defn overlay [input]
  (->> (apply map vector input)
       (map #(first (filter #{0 1} %)))))

(defn render [i]
  (->>
   (map #(if (= 0 %) " " "*") i)
   (partition width )
   (map #(println (join "" %)))))

(def part2
  (render (overlay input)))


