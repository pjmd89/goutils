package systemutils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}
type Argon2 struct {
	aType argonType
	p     *params
}
type argonType int

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

const (
	ARGON_I argonType = iota
	ARGON_ID
)

func NewArgon2() (r *Argon2) {
	r = &Argon2{}
	r.aType = ARGON_I
	r.p = &params{
		memory:      64 * 1024,
		iterations:  4,
		parallelism: 1,
		saltLength:  16,
		keyLength:   32,
	}

	return
}
func (o *Argon2) UseArgon2I() {
	o.aType = ARGON_I
}
func (o *Argon2) UseArgon2ID() {
	o.aType = ARGON_ID
}
func (o *Argon2) SetParams(memory uint32, iterations uint32, parallelism uint8, saltLength uint32, keyLength uint32) {
	o.p.memory = memory
	o.p.iterations = iterations
	o.p.parallelism = parallelism
	o.p.saltLength = saltLength
	o.p.keyLength = keyLength
}
func (o *Argon2) Generate(password string) (r string, err error) {

	return o.generateFromPassword(password)
}
func (o *Argon2) Compare(password string, hash string) (r bool, err error) {

	return o.comparePasswordAndHash(password, hash)
}

func (o *Argon2) generateFromPassword(password string) (encodedHash string, err error) {
	salt, err := o.generateRandomBytes(o.p.saltLength)
	if err != nil {
		return "", err
	}
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	switch o.aType {
	case ARGON_I:
		hash := argon2.IDKey([]byte(password), salt, o.p.iterations, o.p.memory, o.p.parallelism, o.p.keyLength)
		b64Hash := base64.RawStdEncoding.EncodeToString(hash)
		encodedHash = fmt.Sprintf("$argon2i$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, o.p.memory, o.p.iterations, o.p.parallelism, b64Salt, b64Hash)
		break
	case ARGON_ID:
		hash := argon2.IDKey([]byte(password), salt, o.p.iterations, o.p.memory, o.p.parallelism, o.p.keyLength)
		b64Hash := base64.RawStdEncoding.EncodeToString(hash)
		encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, o.p.memory, o.p.iterations, o.p.parallelism, b64Salt, b64Hash)
		break
	}

	return
}
func (o *Argon2) comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := o.decodeHash(encodedHash)
	if err != nil {
		return false, err
	}
	otherHash := argon2.Key([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func (o *Argon2) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (o *Argon2) decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
