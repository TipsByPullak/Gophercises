package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//need to return a httpHandler func
	return func(wr http.ResponseWriter, rd *http.Request) {
		path := rd.URL.Path
		if destPage, ok := pathsToUrls[path]; ok {
			http.Redirect(wr, rd, destPage, http.StatusFound)
			return
		}
		fallback.ServeHTTP(wr, rd)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathsToURLs := buildMap(pathURLs)
	return MapHandler(pathsToURLs, fallback), nil
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathsToURLs := make(map[string]string) //stores all the paths in a map
	for _, pu := range pathURLs {
		pathsToURLs[pu.Path] = pu.URL
	}
	return pathsToURLs
}

func parseYAML(data []byte) ([]pathURL, error) {
	//Use YAML v2 pkg to parse the YML
	var pathURLs []pathURL
	err := yaml.Unmarshal(data, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}
