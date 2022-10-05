package main

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

var routes = []route{
	newRoute("GET", "/.well-known/webfinger", finger),
	/*newRoute("GET", "/", home),
	newRoute("GET", "/contact", contact),
	newRoute("GET", "/api/widgets", apiGetWidgets),
	newRoute("POST", "/api/widgets", apiCreateWidget),
	newRoute("POST", "/api/widgets/([^/]+)", apiUpdateWidget),
	newRoute("POST", "/api/widgets/([^/]+)/parts", apiCreateWidgetPart),
	newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", apiUpdateWidgetPart),
	newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/delete", apiDeleteWidgetPart),
	newRoute("GET", "/([^/]+)", widget),
	newRoute("GET", "/([^/]+)/admin", widgetAdmin),
	newRoute("POST", "/([^/]+)/image", widgetImage),*/
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
