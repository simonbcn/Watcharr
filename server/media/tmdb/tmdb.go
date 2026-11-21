package tmdb

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	gocache "github.com/robfig/go-cache"
	"github.com/sbondCo/Watcharr/database/entity"
)

var ContentStore = gocache.New(time.Hour*24, time.Minute)

type ContentProvider interface {
	CacheContentShow(content ShowDetails, onlyUpdate bool) (entity.Content, error)
	CacheContentMovie(content MovieDetails, onlyUpdate bool) (entity.Content, error)
}

// Language used when none is configured, and as the fallback for
// fields that have no translation in the configured language.
const DefaultLang = "en-US"

type TMDB struct {
	Key string
	// Returns the configured metadata language (may be empty).
	// A func so a config change applies without restarting.
	lang            func() string
	contentProvider ContentProvider
}

func NewTMDB(key string, lang func() string) *TMDB {
	return &TMDB{
		Key:  key,
		lang: lang,
	}
}

func (t *TMDB) AddContentProvider(contentProvider ContentProvider) {
	t.contentProvider = contentProvider
}

func (t *TMDB) GetKey() string {
	if t.Key != "" {
		return t.Key //Config.TMDB_KEY
	}
	return "d047fa61d926371f277e7a83c9c4ff2c"
}

// Language to request metadata in (e.g. "es-ES").
func (t *TMDB) GetLang() string {
	if t.lang != nil {
		if l := t.lang(); l != "" {
			return l
		}
	}
	return DefaultLang
}

func (t *TMDB) apiRequest(ep string, p map[string]string) ([]byte, error) {
	slog.Debug("tmdbAPIRequest", "endpoint", ep, "params", p)
	base, err := url.Parse("https://api.themoviedb.org/3")
	if err != nil {
		return nil, errors.New("failed to parse api uri")
	}

	// Path params
	base.Path += ep

	// Query params
	params := url.Values{}
	params.Add("api_key", t.GetKey())
	// Callers can override the language per request (used for the
	// fallback when a translation is missing).
	if _, ok := p["language"]; !ok {
		params.Add("language", t.GetLang())
	}
	for k, v := range p {
		params.Add(k, v)
	}

	// Add params to url
	base.RawQuery = params.Encode()

	// Run get request
	res, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		slog.Error("TMDB non 200 status code:", "status_code", res.StatusCode)
		return nil, errors.New(string(body))
	}
	return body, nil
}

func (t *TMDB) req(ep string, p map[string]string, resp interface{}) error {
	body, err := t.apiRequest(ep, p)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return err
	}
	return nil
}

// Copy of `p` with the language forced to DefaultLang.
func fallbackParams(p map[string]string) map[string]string {
	m := map[string]string{"language": DefaultLang}
	for k, v := range p {
		if k != "language" {
			m[k] = v
		}
	}
	return m
}
