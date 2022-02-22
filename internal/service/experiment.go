package service

import (
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/bi-zone/ruwireguard-go/crypto/gosthopper"
	"github.com/maoxs2/go-ripemd"
	"github.com/sirupsen/logrus"
	"gitlab.com/Valghall/diwor/internal/results"
	"time"

	streebog256 "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012256"
	streebog512 "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012512"
	"gitlab.com/Valghall/diwor/internal/storage"
)

type ExperimentService struct {
	storage storage.Experiment
}

func NewExperimentService(storage storage.Experiment) *ExperimentService {
	return &ExperimentService{storage: storage}
}

func (es *ExperimentService) GetRecentExperiments(id int) []results.ExperimentDigest {
	return es.storage.GetRecentExperiments(id, 5)
}

func (es *ExperimentService) SaveResults(userId int, algType string, reses Result) {
	switch reses.(type) {
	case results.HashAlgorithmsResults:
		es.storage.SaveHashAlgorithmResults(userId, algType, reses.(results.HashAlgorithmsResults))
	case results.CipherAlgorithmsResults:
		es.storage.SaveCipherAlgorithmResults(userId, algType, reses.(results.CipherAlgorithmsResults))
	default:
		logrus.Error("type err while saving results")
	}
}

func (es *ExperimentService) ResearchHashingAlgorithm(alg string, har *results.HashAlgorithmsResults) results.HashExpResult {
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
	har.StartedAt = begin
	har.FinishedAt = time.Now()

	return res
}

func (es *ExperimentService) ResearchCipheringAlgorithm(alg string, car *results.CipherAlgorithmsResults) results.CipherExpResult {
	var res results.CipherExpResult
	var begin time.Time
	cipherDur := time.Duration(0)
	decipherDur := time.Duration(0)
	var dst []byte

	cipherDurChan := make(chan time.Duration)
	_ = make(chan time.Duration)

	switch alg {
	case Grasshopper:
		begin = time.Now()
		key, _ := generateBytes(32)

		for i := 0; i < 200; i++ {
			go func() {
				start := time.Now()

				kCipher, _ := gosthopper.NewCipher(key)
				kGCM, _ := cipher.NewGCM(kCipher)
				nonce, _ := generateBytes(kGCM.NonceSize())
				kGCM.Seal(dst, nonce, textForCiphering, nil)

				end := time.Now()
				cipherDurChan <- end.Sub(start)
			}()
			/*
				go func() {
					start := time.Now()

					kCipher, _ := gosthopper.NewCipher(key)
					kGCM, _ := cipher.NewGCM(kCipher)
					nonce, _ := generateBytes(kGCM.NonceSize())
					kGCM.Seal(dst, nonce, textForCiphering, nil)

					end := time.Now()
					cipherDurChan <- end.Sub(start)
				}()
			*/

		}

		res = results.CipherExpResult{
			Algorithm: "Кузнечик",
			Type:      "Алгоритм шифрования симметричный",
			KeyLength: 32,
		}

	}

	for i := 0; i < 200; i++ {
		cipherDur += <-cipherDurChan
	}

	res.CipheringDuration = cipherDur / 200
	res.DecipheringDuration = decipherDur
	car.StartedAt = begin
	car.FinishedAt = time.Now()

	return res
}

func (es *ExperimentService) GetLastHashExperimentResults(userId int) results.HashAlgorithmsResults {
	return es.storage.GetLastHashExperimentResults(userId)
}

func (es *ExperimentService) GetLastCipherExperimentResults(userId int) results.CipherAlgorithmsResults {
	return es.storage.GetLastCipherExperimentResults(userId)
}

func (es *ExperimentService) GetAllUserExperiments(id int) []results.ExperimentDigest {
	return es.storage.GetAllUserExperiments(id)
}
