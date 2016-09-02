(ns omnia.etc.output
  (:require [clojure.string :as str]
            [riemann.graphite :refer :all]
            [riemann.config :refer :all]))

(defn graphite-path-statsd [event]
  (let [host (:host event)
        app (re-find #"^.*?\." (:service event))
        service (str/replace-first (:service event) #"^.*?\." "")
        split-host (if host (str/split host #"\.") [])
        split-service (if service (str/split service #" ") [])]
    (str app, (str/join "." (concat (reverse split-host) split-service)))))

(defn add-environment-to-graphite [event]
  (condp = (:plugin event)
    "docker"
      (if (:com.docker.compose.service event)
        (str "docker.", (:com.docker.compose.service event), ".", (riemann.graphite/graphite-path-percentiles event))
        (str "docker.", (riemann.graphite/graphite-path-percentiles event)))
    "statsd" (str "", (graphite-path-statsd event))
    (str "hosts.", (riemann.graphite/graphite-path-percentiles event))))

(def graph (async-queue! :graphite {:queue-size 1000}
            (graphite {:host "graphite" :port 2003 :path add-environment-to-graphite})))