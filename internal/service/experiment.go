package service

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/maoxs2/go-ripemd"
	"github.com/sirupsen/logrus"
	"gitlab.com/Valghall/diwor/internal/results"
	"os"
	"time"

	streebog256 "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012256"
	streebog512 "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012512"
	"gitlab.com/Valghall/diwor/internal/storage"
)

const (
	Streebog256 = "Streebog-256"
	Streebog512 = "Streebog-512"
	SHA224      = "SHA-224"
	SHA256      = "SHA-256"
	SHA384      = "SHA-384"
	SHA512      = "SHA-512"
	RIPEMD128   = "RIPEMD-128"
	RIPEMD160   = "RIPEMD-160"
	RIPEMD256   = "RIPEMD-256"
	RIPEMD320   = "RIPEMD-320"
	MD5         = "MD5"
)

var (
	textForHashing []byte
)

func init() {
	textForHashing, _ = os.ReadFile("lavkraft.txt")
}

type ExperimentService struct {
	storage storage.Experiment
}

func NewExperimentService(storage storage.Experiment) *ExperimentService {
	return &ExperimentService{storage: storage}
}

func (es *ExperimentService) SaveResults(userId int, algType string, results results.HashAlgorithmsResults) {
	es.storage.SaveResults(userId, algType, results)
}

func (es *ExperimentService) ResearchHashingAlgorithm(alg string, has *results.HashAlgorithmsResults) results.HashExpResult {
	var res results.HashExpResult
	var begin time.Time
	dur := time.Duration(0)

	durChan := make(chan time.Duration)

	switch alg {
	case Streebog256:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := streebog256.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := streebog256.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: Streebog256,
			Size:      streebog256.Size,
			BlockSize: streebog256.BlockSize,
			Sample:    sum,
		}
	case Streebog512:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := streebog512.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := streebog512.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: Streebog512,
			Size:      streebog512.Size,
			BlockSize: streebog512.BlockSize,
			Sample:    sum,
		}
	case SHA224:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := sha256.New224()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := sha256.New224()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: SHA224,
			Size:      sha256.Size224,
			BlockSize: sha256.BlockSize,
			Sample:    sum,
		}
	case SHA256:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := sha256.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := sha256.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: SHA256,
			Size:      sha256.Size,
			BlockSize: sha256.BlockSize,
			Sample:    sum,
		}
	case SHA384:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := sha512.New384()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := sha512.New384()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: SHA384,
			Size:      sha512.Size384,
			BlockSize: sha512.BlockSize,
			Sample:    sum,
		}
	case SHA512:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := sha512.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := sha512.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: SHA512,
			Size:      sha512.Size,
			BlockSize: sha512.BlockSize,
			Sample:    sum,
		}
	case RIPEMD128:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := ripemd.New128()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := ripemd.New128()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: RIPEMD128,
			Size:      ripemd.Size128,
			BlockSize: ripemd.BlockSize128,
			Sample:    sum,
		}
	case RIPEMD160:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := ripemd.New160()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := ripemd.New160()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: RIPEMD160,
			Size:      ripemd.Size160,
			BlockSize: ripemd.BlockSize160,
			Sample:    sum,
		}
	case RIPEMD256:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := ripemd.New256()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := ripemd.New256()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: RIPEMD256,
			Size:      ripemd.Size256,
			BlockSize: ripemd.BlockSize256,
			Sample:    sum,
		}
	case RIPEMD320:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				hash := ripemd.New320()
				hash.Write(textForHashing)
				hash.Sum(nil)

				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := ripemd.New320()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: RIPEMD320,
			Size:      ripemd.Size320,
			BlockSize: ripemd.BlockSize320,
			Sample:    sum,
		}
	case MD5:
		begin = time.Now()

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()
				hash := md5.New()
				hash.Write(textForHashing)
				hash.Sum(nil)
				end := time.Now()
				durChan <- end.Sub(start)
			}()
		}

		hash := md5.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		res = results.HashExpResult{
			Algorithm: MD5,
			Size:      md5.Size,
			BlockSize: md5.BlockSize,
			Sample:    sum,
		}
	}

	for i := 0; i < 200; i++ {
		dur += <-durChan
	}
	res.Duration = dur / 200
	logrus.Print(res.Duration)
	has.StartedAt = begin
	has.FinishedAt = time.Now()

	return res
}
