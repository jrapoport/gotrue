package config

import (
	"errors"
	"fmt"
	"github.com/jrapoport/gothic/utils"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"strings"
	"time"

	"github.com/imdario/mergo"
)

// JWT holds all the JWT related configuration.
type JWT struct {
	Secret     string `json:"secret"`
	PEM        `yaml:",inline" mapstructure:",squash"`
	Algorithm  string        `json:"algorithm"`
	Expiration time.Duration `json:"expiration"`
	// Issuer is the the entity that issued the token (default: Config.Service)
	Issuer string `json:"issuer"`
	// Audience is an optional comma separated list of resource
	// servers that should accept the token (default: n/a)
	Audience string `json:"audience"`
	// Scope is an optional comma separated list of permission
	// scopes for the token (default: n/a)
	Scope string `json:"scope"`
	sk    jwk.Key
	pk    jwk.Key
}

type PEM struct {
	PrivateKey string `json:"privatekey"`
	PublicKey  string `json:"publickey"`
}

func (j *JWT) normalize(srv Service, def JWT) {
	if def.Issuer == "" {
		def.Issuer = strings.ToLower(srv.Name)
	}
	// no error is possible here since we
	// control the struct entirely
	_ = mergo.Merge(j, def)
}

func (j *JWT) CheckRequired() error {
	if j.Secret == "" && j.PEM.PrivateKey == "" {
		return errors.New("jwt secret or private key is required")
	} else if j.Secret != "" {
		return nil
	}
	if j.PEM.PrivateKey != "" && !utils.PathExists(j.PEM.PrivateKey) {
		err := fmt.Errorf("jwt private key not found: %s", j.PEM.PrivateKey)
		return err
	}
	if j.PEM.PublicKey != "" && !utils.PathExists(j.PEM.PublicKey) {
		err := fmt.Errorf("jwt public key not found: %s", j.PEM.PublicKey)
		return err
	}
	return nil
}

func (j *JWT) PrivateKey() jwk.Key {
	if j.sk != nil {
		return j.sk
	}
	if j.Secret != "" {
		sec := []byte(j.Secret)
		j.sk, _ = jwk.New(sec)
	} else {
		j.sk, _ = pemKey(j.PEM.PrivateKey)
	}
	return j.sk
}

func (j *JWT) PublicKey() jwk.Key {
	if j.pk != nil {
		return j.pk
	}
	if j.Secret != "" {
		sec := []byte(j.Secret)
		j.pk, _ = jwk.New(sec)
	} else {
		pem := j.PEM.PublicKey
		if pem == "" {
			const publicKeyExt = ".pub"
			pem = j.PEM.PrivateKey + publicKeyExt
		}
		j.pk, _ = pemKey(pem)
	}
	return j.pk
}

func pemKey(pem string) (jwk.Key, error) {
	raw, err := ioutil.ReadFile(pem)
	if err != nil {
		return nil, err
	}
	key, err := jwk.ParseKey(raw, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}
