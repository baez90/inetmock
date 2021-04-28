package mock

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"gitlab.com/inetmock/inetmock/internal/rules"
)

var (
	ErrUnknownFilterMethod = errors.New("no filter with the given name is known")
	ErrAmbiguousParamCount = errors.New("the supplied number of arguments does not match the expected one")

	knownRequestFilters = map[string]func(args ...rules.Param) (RequestFilter, error){
		"method":      HTTPMethodMatcher,
		"pathpattern": PathPatternMatcher,
		"header":      HeaderValueMatcher,
	}
)

type RequestFilter interface {
	Matches(req *http.Request) bool
}

func RequestFiltersForRoutingRule(rule *rules.Routing) (filters []RequestFilter, err error) {
	if rule.Filters == nil || len(rule.Filters.Chain) == 0 {
		return nil, nil
	}
	filters = make([]RequestFilter, len(rule.Filters.Chain))
	for idx := range rule.Filters.Chain {
		if constructor, ok := knownRequestFilters[strings.ToLower(rule.Filters.Chain[idx].Name)]; !ok {
			return nil, fmt.Errorf("%w %s", ErrUnknownFilterMethod, rule.Filters.Chain[idx].Name)
		} else {
			var instance RequestFilter
			instance, err = constructor(rule.Filters.Chain[idx].Params...)
			if err != nil {
				return
			}
			filters[idx] = instance
		}
	}
	return
}

type RequestFilterFunc func(req *http.Request) bool

func (r RequestFilterFunc) Matches(req *http.Request) bool {
	return r(req)
}

func HTTPMethodMatcher(args ...rules.Param) (RequestFilter, error) {
	if err := validateParameterCount(args, 1); err != nil {
		return nil, err
	}

	var err error

	var expectedMethod string
	if expectedMethod, err = args[0].AsString(); err != nil {
		return nil, err
	}

	return RequestFilterFunc(func(req *http.Request) bool {
		return strings.EqualFold(req.Method, expectedMethod)
	}), nil
}

func PathPatternMatcher(args ...rules.Param) (RequestFilter, error) {
	if err := validateParameterCount(args, 1); err != nil {
		return nil, err
	}

	var err error
	var rawPattern string
	if rawPattern, err = args[0].AsString(); err != nil {
		return nil, err
	}

	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return nil, err
	}

	return RequestFilterFunc(func(req *http.Request) bool {
		return pattern.MatchString(req.URL.Path)
	}), nil
}

func HeaderValueMatcher(args ...rules.Param) (RequestFilter, error) {
	if err := validateParameterCount(args, 2); err != nil {
		return nil, err
	}

	var err error
	var headerName, expectedValue string
	if headerName, err = args[0].AsString(); err != nil {
		return nil, err
	}

	if expectedValue, err = args[1].AsString(); err != nil {
		return nil, err
	}

	return RequestFilterFunc(func(req *http.Request) bool {
		values := req.Header.Values(headerName)
		for idx := range values {
			if strings.EqualFold(expectedValue, values[idx]) {
				return true
			}
		}
		return false
	}), nil
}

func validateParameterCount(params []rules.Param, expected int) error {
	if len(params) != expected {
		return fmt.Errorf("%w: expected %d2 got %d", ErrAmbiguousParamCount, expected, len(params))
	}
	return nil
}
