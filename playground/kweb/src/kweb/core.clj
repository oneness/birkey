(ns kweb.core
  (:require [catacumba.core :as ct]
            [unilog.config  :refer [start-logging!]])
  (:import (org.slf4j MDC)))

(def logger
  (start-logging! {:level "all"
                   :console "%p [%d] %t - %c %m%n%X{debug}"}))

(defn info [msg]
  (.info logger msg))

(defn all-handler
  [context]
  (when (Boolean/valueOf (get-in context [:query-params :debug]))
    (MDC/put "debug" (str context))
    (info "all-handler"))
  "Hello World")

(defn foo-handler
  [context]
  (when (Boolean/valueOf (get-in context [:query-params :debug]))
    (MDC/put "debug" (str context))
    (info "foo-handler"))
  "Hello World")

(def app
  (ct/routes [[:all "" #'all-handler]
              [:get "foobar" #'foo-handler]]))

(defn -main
  [& args]
  (ct/run-server app {:port 3030}))
