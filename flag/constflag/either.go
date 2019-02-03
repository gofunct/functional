package constflag

import (
	"flag"
	"fmt"
	"github.com/gofunct/functional/errors"
	"os"
	"strings"
)

// Either tries to parse the argument using `Either`, and if that fails, using `Or`.
// `ChoseEither` is true if the first attempt succeed.
type Either struct {
	Either   flag.Value
	Or       flag.Value
	ChoseEither bool
	Env 	 string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *Either) Help() string {
	if fv.Either != nil && fv.Or != nil {
		if eitherHelp, ok := fv.Either.(interface {
			Help() string
		}); ok {
			if orHelp, ok := fv.Or.(interface {
				Help() string
			}); ok {
				return fmt.Sprintf("either %s, or %s", eitherHelp.Help(), orHelp.Help())
			}
		}
	}
	return ""
}

// Set is flag.Value.Set
func (fv *Either) Set(v string) error {
	if err := os.Setenv(fv.Env, v); err != nil {
		return errors.Wrap(err, "failed to bind env to flag value")
	}
	err := fv.Either.Set(v)
	fv.ChoseEither = err == nil
	if err != nil {
		return fv.Or.Set(v)
	}
	return nil
}

func (fv *Either) String() string {
	if fv.Env != "" {
		if s, ok := os.LookupEnv(strings.ToUpper(fv.Env)); ok == true {
			return s
		}
	}
	if fv.ChoseEither {
		return fv.Either.String()
	}
	if fv.Or != nil {
		return fv.Or.String()
	}
	return ""
}