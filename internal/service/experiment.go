package service

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/maoxs2/go-ripemd"
	"gitlab.com/Valghall/diwor/internal/results"
	"os"
	"time"

	streebog "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012256"
	"gitlab.com/Valghall/diwor/internal/storage"
)

const (
	Streebog  = "Streebog"
	SHA224    = "SHA-224"
	SHA256    = "SHA-256"
	SHA384    = "SHA-384"
	SHA512    = "SHA-512"
	RIPEMD128 = "RIPEMD-128"
	RIPEMD160 = "RIPEMD-160"
	RIPEMD256 = "RIPEMD-256"
	RIPEMD320 = "RIPEMD-320"
	MD5       = "MD5"
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

func (es *ExperimentService) ResearchHashingAlgorithm(alg string, has *results.HashAlgorithmsResults) Result {
	var start time.Time
	var end time.Time

	switch alg {
	case Streebog:
		start = time.Now()

		hash := streebog.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: Streebog,
			Duration:  end.Sub(start),
			Size:      streebog.Size,
			BlockSize: streebog.BlockSize,
			Sample:    sum,
		}
	case SHA224:
		start = time.Now()

		hash := sha256.New224()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))
		time.Sleep(time.Second)
		end = time.Now()

		return results.HashExpResult{
			Algorithm: SHA224,
			Duration:  end.Sub(start),
			Size:      sha256.Size224,
			BlockSize: sha256.BlockSize,
			Sample:    sum,
		}
	case SHA256:
		start = time.Now()

		hash := sha256.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: SHA256,
			Duration:  end.Sub(start),
			Size:      sha256.Size,
			BlockSize: sha256.BlockSize,
			Sample:    sum,
		}
	case SHA384:
		start = time.Now()

		hash := sha512.New384()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: SHA384,
			Duration:  end.Sub(start),
			Size:      sha512.Size384,
			BlockSize: sha512.BlockSize,
			Sample:    sum,
		}
	case SHA512:
		start = time.Now()

		hash := sha512.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: SHA512,
			Duration:  end.Sub(start),
			Size:      sha512.Size,
			BlockSize: sha512.BlockSize,
			Sample:    sum,
		}
	case RIPEMD128:
		start = time.Now()

		hash := ripemd.New128()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: RIPEMD128,
			Duration:  end.Sub(start),
			Size:      ripemd.Size128,
			BlockSize: ripemd.BlockSize128,
			Sample:    sum,
		}
	case RIPEMD160:
		start = time.Now()

		hash := ripemd.New160()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: RIPEMD160,
			Duration:  end.Sub(start),
			Size:      ripemd.Size160,
			BlockSize: ripemd.BlockSize160,
			Sample:    sum,
		}
	case RIPEMD256:
		start = time.Now()

		hash := ripemd.New256()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: RIPEMD256,
			Duration:  end.Sub(start),
			Size:      ripemd.Size256,
			BlockSize: ripemd.BlockSize256,
			Sample:    sum,
		}
	case RIPEMD320:
		start = time.Now()

		hash := ripemd.New320()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: RIPEMD320,
			Duration:  end.Sub(start),
			Size:      ripemd.Size320,
			BlockSize: ripemd.BlockSize320,
			Sample:    sum,
		}
	case MD5:
		start = time.Now()

		hash := md5.New()
		hash.Write(textForHashing)
		sum := fmt.Sprintf("%x", hash.Sum(nil))

		end = time.Now()

		return results.HashExpResult{
			Algorithm: MD5,
			Duration:  end.Sub(start),
			Size:      md5.Size,
			BlockSize: md5.BlockSize,
			Sample:    sum,
		}
	}

	has.StartedAt = start
	has.FinishedAt = end

	return nil
}
