package urlshort

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"net/http"
)

type urlPath struct {
	Path string
	Url  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the Path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path

		if val, ok := pathsToUrls[p]; ok {
			http.Redirect(w, r, val, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the Path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - Path: /some-Path
//       Url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the Path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//    {
//      "path": "/urlshort-final",
//      "url": "https://github.com/gophercises/urlshort/tree/solution"
//    }
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(data)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJSON)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]urlPath, error) {
	var paths []urlPath
	err := yaml.Unmarshal(yml, &paths)
	return paths, err
}

func parseJSON(data []byte) ([]urlPath, error) {
	var paths []urlPath
	err := json.Unmarshal(data, &paths)
	return paths, err
}

func buildMap(paths []urlPath) map[string]string {
	pathMap := make(map[string]string, len(paths))

	for _, path := range paths {
		pathMap[path.Path] = path.Url
	}

	return pathMap
}
