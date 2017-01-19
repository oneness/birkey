(defproject kweb "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.8.0"]
                 [funcool/catacumba "1.2.0"]
                 [spootnik/unilog "0.7.17"]]
  :main ^:skip-aot kweb.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})
