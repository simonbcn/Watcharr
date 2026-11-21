package tmdb

import (
	"errors"
	"log/slog"
	"sort"
	"strings"
)

// A selectable metadata language (for the TMDB_LANG setting).
type Language struct {
	// TMDB `language` value, e.g. "es-ES"
	Code string `json:"code"`
	// Human readable label, e.g. "Spanish (ES)"
	Name string `json:"name"`
}

// Languages TMDB has translations for (its primary translations).
func (t *TMDB) Languages() ([]Language, error) {
	var codes []string
	err := t.req(
		"/configuration/primary_translations",
		map[string]string{},
		&codes)
	if err != nil {
		slog.Error("Languages: Translations request failed", "error", err)
		return nil, errors.New("request failed")
	}
	var names []struct {
		Iso6391     string `json:"iso_639_1"`
		EnglishName string `json:"english_name"`
	}
	err = t.req("/configuration/languages", map[string]string{}, &names)
	if err != nil {
		slog.Error("Languages: Languages request failed", "error", err)
		return nil, errors.New("request failed")
	}
	nameByIso := make(map[string]string, len(names))
	for _, n := range names {
		nameByIso[n.Iso6391] = n.EnglishName
	}
	langs := make([]Language, 0, len(codes))
	for _, code := range codes {
		lang, region, _ := strings.Cut(code, "-")
		name := nameByIso[lang]
		if name == "" {
			name = lang
		}
		if region != "" {
			name += " (" + region + ")"
		}
		langs = append(langs, Language{Code: code, Name: name})
	}
	sort.Slice(langs, func(a, b int) bool {
		return langs[a].Name < langs[b].Name
	})
	return langs, nil
}
