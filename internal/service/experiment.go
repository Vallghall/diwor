package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"gitlab.com/Valghall/diwor/internal/plotconfig"
	"golang.org/x/crypto/blowfish"
	"time"

	"github.com/bi-zone/ruwireguard-go/crypto/gosthopper"
	"github.com/maoxs2/go-ripemd"
	"github.com/sirupsen/logrus"
	"gitlab.com/Valghall/diwor/internal/results"

	streebog256 "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012256"
	streebog512 "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012512"
	"gitlab.com/Valghall/diwor/internal/storage"
)

type ExperimentService struct {
	storage storage.Experiment
}

const (
	HashAlgorithm   = "Алгоритм хэширования"
	CipherAlgorithm = "Алгоритм шифрования"
)

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

func (es *ExperimentService) ResearchHashingAlgorithm(alg string, conf plotconfig.Config) results.HashExpResult {
	var res results.HashExpResult
	x, y := make([]int, 0), make([]int, 0)

	for l := conf.From; l <= conf.To; l += conf.Step {
		dur := time.Duration(0)
		textForHashing, _ := generateBytes(l)

		switch alg {
		case Streebog256:
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := streebog256.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := streebog512.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := sha256.New224()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := sha256.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := sha512.New384()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := sha512.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := ripemd.New128()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := ripemd.New160()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := ripemd.New256()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				hash := ripemd.New320()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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
			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()
				hash := md5.New()
				hash.Write(textForHashing)
				hash.Sum(nil)

				dur += time.Since(start)
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

		res.Duration = dur / time.Duration(conf.NumMeasurements)
		x = append(x, l)
		y = append(y, int(res.Duration.Microseconds()))
	}

	res.Plot.X = x
	res.Plot.Y = y

	return res
}

func (es *ExperimentService) ResearchCipheringAlgorithm(alg string, conf plotconfig.Config) results.CipherExpResult {
	var res results.CipherExpResult
	x, y := make([]int, 0), make([]int, 0)

	for l := conf.From; l <= conf.To; l += conf.Step {
		cipherDur := time.Duration(0)
		decipherDur := time.Duration(0)

		switch alg {
		case Grasshopper:
			key, _ := generateBytes(32)
			var dst, nonce []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				kCipher, _ := gosthopper.NewCipher(key)
				kGCM, _ := cipher.NewGCM(kCipher)
				nonce, _ = generateBytes(kGCM.NonceSize())
				kGCM.Seal(dst, nonce, textForCiphering, nil)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				kCipher, _ := gosthopper.NewCipher(key)
				kGCM, _ := cipher.NewGCM(kCipher)
				kGCM.Open(dst[:0], nonce, dst, nil)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: "Кузнечик",
				Type:      "Алгоритм шифрования симметричный",
				KeyLength: 32,
			}
		case AES128_GCM:
			key, _ := generateBytes(aes.BlockSize)
			var dst, nonce []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				aesCipher, _ := aes.NewCipher(key)
				aesGCM, _ := cipher.NewGCM(aesCipher)
				nonce, _ = generateBytes(aesGCM.NonceSize())
				aesGCM.Seal(dst, nonce, textForCiphering, nil)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				aesCipher, _ := aes.NewCipher(key)
				aesGCM, _ := cipher.NewGCM(aesCipher)
				aesGCM.Open(dst[:0], nonce, dst, nil)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: alg,
				Type:      "Алгоритм шифрования симметричный",
				KeyLength: aes.BlockSize,
			}
		case DES_GCM:
			key, _ := generateBytes(des.BlockSize)
			var dst, nonce []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				desCipher, _ := des.NewCipher(key)
				desGCM, _ := cipher.NewGCM(desCipher)
				nonce, _ = generateBytes(desGCM.NonceSize())
				desGCM.Seal(dst, nonce, textForCiphering, nil)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				desCipher, _ := des.NewCipher(key)
				desGCM, _ := cipher.NewGCM(desCipher)
				desGCM.Open(dst[:0], nonce, dst, nil)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: alg,
				Type:      "Алгоритм шифрования симметричный",
				KeyLength: des.BlockSize,
			}
		case DES_CFB:
			key, _ := generateBytes(des.BlockSize)
			var dst, iv []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				desCipher, _ := des.NewCipher(key)
				iv, _ = generateBytes(des.BlockSize)

				desEncrypter := cipher.NewCFBEncrypter(desCipher, iv)
				desEncrypter.XORKeyStream(dst, textForCiphering)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				desCipher, _ := des.NewCipher(key)
				desDecrypter := cipher.NewCFBDecrypter(desCipher, iv)
				desDecrypter.XORKeyStream(dst[:0], dst)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: alg,
				Type:      "Алгоритм шифрования симметричный",
				KeyLength: des.BlockSize,
			}
		case AES128_CFB:
			key, _ := generateBytes(aes.BlockSize)
			var dst, iv []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				aesCipher, _ := aes.NewCipher(key)
				iv, _ = generateBytes(aes.BlockSize)

				aesEncrypter := cipher.NewCFBEncrypter(aesCipher, iv)
				aesEncrypter.XORKeyStream(dst, textForCiphering)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				aesCipher, _ := aes.NewCipher(key)
				aesDecrypter := cipher.NewCFBDecrypter(aesCipher, iv)
				aesDecrypter.XORKeyStream(dst[:0], dst)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: alg,
				Type:      "Алгоритм шифрования симметричный",
				KeyLength: aes.BlockSize,
			}
		case RSA:
			keyPair, _ := rsa.GenerateKey(rand.Reader, 2048)

			var dst, label []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				label = []byte("OAEP Encrypted")
				rng := rand.Reader
				dst, _ = rsa.EncryptOAEP(sha256.New(), rng, &keyPair.PublicKey, textForCiphering, label)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				rng := rand.Reader
				dst, _ = rsa.DecryptOAEP(sha256.New(), rng, keyPair, dst, label)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: alg,
				Type:      "Алгоритм шифрования асимметричный",
				KeyLength: 256,
			}
		case BF_CFB:
			key, _ := generateBytes(blowfish.BlockSize)
			var dst, iv []byte

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				bfCipher, _ := blowfish.NewCipher(key)
				iv, _ = generateBytes(blowfish.BlockSize)

				bfEncrypter := cipher.NewCFBEncrypter(bfCipher, iv)
				bfEncrypter.XORKeyStream(dst, textForCiphering)

				cipherDur += time.Since(start)
			}

			for i := 0; i < conf.NumMeasurements; i++ {
				start := time.Now()

				bfCipher, _ := blowfish.NewCipher(key)
				bfDecrypter := cipher.NewCFBDecrypter(bfCipher, iv)
				bfDecrypter.XORKeyStream(dst[:0], dst)

				decipherDur += time.Since(start)
			}

			res = results.CipherExpResult{
				Algorithm: alg,
				Type:      "Алгоритм шифрования симметричный",
				KeyLength: blowfish.BlockSize,
			}
		}

		res.CipheringDuration = cipherDur / time.Duration(conf.NumMeasurements)
		res.DecipheringDuration = decipherDur / time.Duration(conf.NumMeasurements)
		x = append(x, l)
		y = append(y, int(res.CipheringDuration.Microseconds()))
	}

	res.Plot.X = x
	res.Plot.Y = y

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

func (es *ExperimentService) GetUserExperimentResultBySortedId(alg string, userId, sortedId int) (Result, error) {
	switch alg {
	case HashAlgorithm:
		return es.storage.GetUserHashExperimentResultBySortedId(userId, sortedId)
	case CipherAlgorithm:
		return es.storage.GetUserCipherExperimentResultBySortedId(userId, sortedId)
	default:
		return nil, errors.New("error wile fetching user user's results by sortedId")
	}
}
