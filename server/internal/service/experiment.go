package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	rand2 "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/andreburgaud/crypt2go/ecb"
	streebog512 "github.com/martinlindhe/gogost/gost34112012512"
	eg "github.com/mirzazhar/elgamal"
	magma "github.com/pedroalbanese/gogost/gost341264"
	"gitlab.com/Valghall/diwor/server/internal/plotconfig"
	"gitlab.com/Valghall/diwor/server/internal/results"
	"gitlab.com/Valghall/diwor/server/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/openpgp/elgamal"
	"golang.org/x/crypto/sha3"
	"golang.org/x/tools/benchmark/parse"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bi-zone/ruwireguard-go/crypto/gosthopper"
	"github.com/maoxs2/go-ripemd"
	streebog256 "github.com/martinlindhe/gogost/gost34112012256"
	"github.com/sirupsen/logrus"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type ExperimentService struct {
	storage storage.Experiment
}

const (
	HashAlgorithm   = "Алгоритм хэширования"
	CipherAlgorithm = "Алгоритм шифрования"
	mul2MBperS      = 953.674
)

func NewExperimentService(storage storage.Experiment) *ExperimentService {
	return &ExperimentService{storage: storage}
}

func (es *ExperimentService) GetRecentExperiments(id int) ([]results.ExperimentDigest, error) {
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
	var bench func(b *testing.B)
	var br testing.BenchmarkResult
	var b *parse.Benchmark
	logrus.Println(conf.From)

	x, y := make([]int, 0), make([]int, 0)

	for l := conf.From; l <= conf.To; l += conf.Step {
		textForHashing, _ := generateBytes(l)

		switch alg {
		case BCRYPT:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					bcrypt.GenerateFromPassword(textForHashing, 0)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sample, _ := bcrypt.GenerateFromPassword(textForHashing, 0)
				res = results.HashExpResult{
					Algorithm: BCRYPT,
					Size:      0,
					BlockSize: 0,
					Sample:    fmt.Sprintf("%x", sample),
				}
			}
		case Streebog256:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(streebog256.New(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.HashExpResult{
					Algorithm: Streebog256,
					Size:      streebog256.Size,
					BlockSize: streebog256.BlockSize,
					Sample:    fmt.Sprintf("%x", hash(streebog256.New(), textForHashing)),
				}
			}
		case Streebog512:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(streebog512.New(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(streebog512.New(), textForHashing))

				res = results.HashExpResult{
					Algorithm: Streebog512,
					Size:      streebog512.Size,
					BlockSize: streebog512.BlockSize,
					Sample:    sum,
				}
			}
		case SHA224:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha256.New224(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(sha256.New224(), textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA224,
					Size:      sha256.Size224,
					BlockSize: sha256.BlockSize,
					Sample:    sum,
				}
			}
		case SHA256:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha256.New(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(sha256.New(), textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA256,
					Size:      sha256.Size,
					BlockSize: sha256.BlockSize,
					Sample:    sum,
				}
			}
		case SHA384:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha512.New384(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(sha512.New384(), textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA384,
					Size:      sha512.Size384,
					BlockSize: sha512.BlockSize,
					Sample:    sum,
				}
			}
		case SHA512:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha512.New(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(sha512.New(), textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA512,
					Size:      sha512.Size,
					BlockSize: sha512.BlockSize,
					Sample:    sum,
				}
			}
		case SHA3_224:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha3.New224(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				h := sha3.New224()
				sum := fmt.Sprintf("%x", hash(h, textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA3_224,
					Size:      h.Size(),
					BlockSize: h.BlockSize(),
					Sample:    sum,
				}
			}
		case SHA3_256:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha3.New256(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				h := sha3.New256()
				sum := fmt.Sprintf("%x", hash(h, textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA3_256,
					Size:      h.Size(),
					BlockSize: h.BlockSize(),
					Sample:    sum,
				}
			}
		case SHA3_384:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha3.New384(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)
			h := sha3.New384()

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(h, textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA3_384,
					Size:      h.Size(),
					BlockSize: h.BlockSize(),
					Sample:    sum,
				}
			}
		case SHA3_512:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(sha3.New512(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				h := sha3.New512()
				sum := fmt.Sprintf("%x", hash(h, textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA3_512,
					Size:      h.Size(),
					BlockSize: h.BlockSize(),
					Sample:    sum,
				}
			}
		case SHA3_SHAKE128:
			bench = func(b *testing.B) {
				h := make([]byte, 16)
				for i := 0; i < b.N; i++ {
					sha3.ShakeSum128(h, textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				h := sha3.New512()
				sum := fmt.Sprintf("%x", hash(h, textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA3_SHAKE128,
					Size:      h.Size(),
					BlockSize: h.BlockSize(),
					Sample:    sum,
				}
			}
		case SHA3_SHAKE256:
			bench = func(b *testing.B) {
				h := make([]byte, 32)
				for i := 0; i < b.N; i++ {
					sha3.ShakeSum256(h, textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				h := sha3.New512()
				sum := fmt.Sprintf("%x", hash(h, textForHashing))

				res = results.HashExpResult{
					Algorithm: SHA3_SHAKE256,
					Size:      h.Size(),
					BlockSize: h.BlockSize(),
					Sample:    sum,
				}
			}
		case RIPEMD128:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(ripemd.New128(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(ripemd.New128(), textForHashing))

				res = results.HashExpResult{
					Algorithm: RIPEMD128,
					Size:      ripemd.Size128,
					BlockSize: ripemd.BlockSize128,
					Sample:    sum,
				}
			}
		case RIPEMD160:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(ripemd.New160(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(ripemd.New160(), textForHashing))

				res = results.HashExpResult{
					Algorithm: RIPEMD160,
					Size:      ripemd.Size160,
					BlockSize: ripemd.BlockSize160,
					Sample:    sum,
				}
			}
		case RIPEMD256:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(ripemd.New256(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(ripemd.New256(), textForHashing))

				res = results.HashExpResult{
					Algorithm: RIPEMD256,
					Size:      ripemd.Size256,
					BlockSize: ripemd.BlockSize256,
					Sample:    sum,
				}
			}
		case RIPEMD320:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(ripemd.New320(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(ripemd.New320(), textForHashing))

				res = results.HashExpResult{
					Algorithm: RIPEMD320,
					Size:      ripemd.Size320,
					BlockSize: ripemd.BlockSize320,
					Sample:    sum,
				}
			}
		case MD5:
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(md5.New(), textForHashing)
				}
			}
			br = testing.Benchmark(bench)

			if l >= conf.To {
				sum := fmt.Sprintf("%x", hash(md5.New(), textForHashing))

				res = results.HashExpResult{
					Algorithm: MD5,
					Size:      md5.Size,
					BlockSize: md5.BlockSize,
					Sample:    sum,
				}
			}
		}

		b, _ = parse.ParseLine("Benchmark" + br.String() + br.MemString())

		if l >= conf.To {
			d := (float64(l) * mul2MBperS) / (b.NsPerOp)
			res.Duration = math.Round(d*100) / 100
		}

		x = append(x, l)
		y = append(y, int(b.NsPerOp))
	}

	res.Plot.X = x
	res.Plot.Y = y
	res.Hyst.AlocX = b.AllocsPerOp
	res.Hyst.OpX = b.N
	res.Hyst.Alg = alg

	if res.Duration == math.Inf(1) {
		res.Duration = 0
	}

	return res
}

func (es *ExperimentService) ResearchCipheringAlgorithm(alg string, conf plotconfig.Config) results.CipherExpResult {
	var res results.CipherExpResult
	var bench func(b *testing.B)
	var br struct {
		ciphering   testing.BenchmarkResult
		deciphering testing.BenchmarkResult
	}

	var b []string

	x, y := make([]int, 0), make([]int, 0)

	for l := conf.From; l <= conf.To; l += conf.Step {
		textForCiphering, _ := generateBytes(l)

		switch alg {
		case Grasshopper_GCM:
			key, _ := generateBytes(gosthopper.BlockSize * 2)

			kCipher, _ := gosthopper.NewCipher(key)
			kGCM, _ := cipher.NewGCM(kCipher)
			nonce, _ := generateBytes(kGCM.NonceSize())
			ciphered := GCMSeal(kGCM, nonce, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					GCMSeal(kGCM, nonce, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					GCMOpen(kGCM, nonce, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: "Кузнечик",
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: 32,
				}
			}
		case Grasshopper_ECB:
			key, _ := generateBytes(gosthopper.BlockSize * 2)

			block, _ := gosthopper.NewCipher(key)
			mode := ecb.NewECBEncrypter(block)
			decMode := ecb.NewECBDecrypter(block)
			ciphered := ECBSeal(mode, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBSeal(mode, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBOpen(decMode, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: gosthopper.BlockSize,
				}
			}
		case Grasshopper_OFB:
			key, _ := generateBytes(gosthopper.BlockSize * 2)

			block, _ := gosthopper.NewCipher(key)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, gosthopper.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, gosthopper.BlockSize, textForCiphering)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: gosthopper.BlockSize,
				}
			}
		case Grasshopper_CFB:
			key, _ := generateBytes(gosthopper.BlockSize * 2)

			ghCipher, _ := gosthopper.NewCipher(key)
			ciphered := CFBSeal(ghCipher, gosthopper.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBSeal(ghCipher, gosthopper.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBOpen(ghCipher, gosthopper.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: gosthopper.BlockSize,
				}
			}
		case Grasshopper_CTR:
			key, _ := generateBytes(gosthopper.BlockSize * 2)

			block, _ := gosthopper.NewCipher(key)
			ciphered := CTRSeal(block, gosthopper.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTRSeal(block, gosthopper.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTROpen(block, gosthopper.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: gosthopper.BlockSize,
				}
			}

		case AES128_GCM:
			key, _ := generateBytes(aes.BlockSize)

			aesCipher, _ := aes.NewCipher(key)
			aesGCM, _ := cipher.NewGCM(aesCipher)
			nonce, _ := generateBytes(aesGCM.NonceSize())
			ciphered := GCMSeal(aesGCM, nonce, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					GCMSeal(aesGCM, nonce, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					GCMOpen(aesGCM, nonce, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: aes.BlockSize,
				}
			}
		case AES128_CFB:
			key, _ := generateBytes(aes.BlockSize)

			aesCipher, _ := aes.NewCipher(key)
			ciphered := CFBSeal(aesCipher, aes.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBSeal(aesCipher, aes.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBOpen(aesCipher, aes.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: aes.BlockSize,
				}
			}
		case AES128_ECB:
			key, _ := generateBytes(aes.BlockSize)

			block, _ := aes.NewCipher(key)
			mode := ecb.NewECBEncrypter(block)
			decMode := ecb.NewECBDecrypter(block)
			ciphered := ECBSeal(mode, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBSeal(mode, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBOpen(decMode, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: aes.BlockSize,
				}
			}
		case AES128_OFB:
			key, _ := generateBytes(aes.BlockSize)

			block, _ := aes.NewCipher(key)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, aes.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, aes.BlockSize, textForCiphering)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: aes.BlockSize,
				}
			}
		case AES128_CTR:
			key, _ := generateBytes(aes.BlockSize)

			block, _ := aes.NewCipher(key)
			ciphered := CTRSeal(block, aes.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTRSeal(block, aes.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTROpen(block, aes.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: aes.BlockSize,
				}
			}

		case MGM_CFB:
			key, _ := generateBytes(magma.BlockSize)

			block := magma.NewCipher(key)
			ciphered := CFBSeal(block, magma.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBSeal(block, magma.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBOpen(block, magma.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: magma.BlockSize,
				}
			}
		case MGM_ECB:
			key, _ := generateBytes(magma.BlockSize)

			block := magma.NewCipher(key)
			mode := ecb.NewECBEncrypter(block)
			decMode := ecb.NewECBDecrypter(block)
			ciphered := ECBSeal(mode, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBSeal(mode, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBOpen(decMode, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: magma.BlockSize,
				}
			}
		case MGM_OFB:
			key, _ := generateBytes(magma.BlockSize)

			block := magma.NewCipher(key)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, magma.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, magma.BlockSize, textForCiphering)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: magma.BlockSize,
				}
			}
		case MGM_CTR:
			key, _ := generateBytes(magma.BlockSize)

			block := magma.NewCipher(key)
			ciphered := CTRSeal(block, magma.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTRSeal(block, magma.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTROpen(block, magma.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: magma.BlockSize,
				}
			}

		case DES_CFB:
			key, _ := generateBytes(des.BlockSize)

			desCipher, _ := des.NewCipher(key)
			ciphered := CFBSeal(desCipher, des.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBSeal(desCipher, des.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBOpen(desCipher, des.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: des.BlockSize,
				}
			}
		case DES_ECB:
			key, _ := generateBytes(des.BlockSize)

			block, _ := des.NewCipher(key)
			mode := ecb.NewECBEncrypter(block)
			decMode := ecb.NewECBDecrypter(block)
			ciphered := ECBSeal(mode, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBSeal(mode, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBOpen(decMode, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: des.BlockSize,
				}
			}
		case DES_OFB:
			key, _ := generateBytes(des.BlockSize)

			block, _ := des.NewCipher(key)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, des.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, des.BlockSize, textForCiphering)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: des.BlockSize,
				}
			}
		case DES_CTR:
			key, _ := generateBytes(des.BlockSize)

			block, _ := des.NewCipher(key)
			ciphered := CTRSeal(block, des.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTRSeal(block, des.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTROpen(block, des.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: des.BlockSize,
				}
			}

		case RSA:
			keyPair, _ := rsa.GenerateKey(rand2.Reader, 2048)
			label := []byte("OAEP Encrypted")
			dst, _ := rsa.EncryptOAEP(sha256.New(), rand2.Reader, &keyPair.PublicKey, textForCiphering, label)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					rsa.EncryptOAEP(sha256.New(), rand2.Reader, &keyPair.PublicKey, textForCiphering, label)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					rsa.DecryptOAEP(sha256.New(), rand2.Reader, keyPair, dst, label)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования асимметричный",
					KeyLength: keyPair.Size(),
				}
			}
		case EG:
			key, _ := eg.GenerateKey(64, 1)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					elgamal.Encrypt(rand2.Reader, &elgamal.PublicKey{
						G: key.G,
						P: key.P,
						Y: key.Y,
					}, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)
			c1, c2, _ := elgamal.Encrypt(rand2.Reader, &elgamal.PublicKey{
				G: key.G,
				P: key.P,
				Y: key.Y,
			}, textForCiphering)
			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					elgamal.Decrypt(&elgamal.PrivateKey{}, c1, c2)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования асимметричный",
					KeyLength: -1,
				}
			}

		case BF_CFB:
			key, _ := generateBytes(blowfish.BlockSize)

			bfCipher, _ := blowfish.NewCipher(key)
			ciphered := CFBSeal(bfCipher, blowfish.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBSeal(bfCipher, blowfish.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CFBOpen(bfCipher, blowfish.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: blowfish.BlockSize,
				}
			}
		case BF_ECB:
			key, _ := generateBytes(blowfish.BlockSize)

			block, _ := blowfish.NewCipher(key)
			mode := ecb.NewECBEncrypter(block)
			decMode := ecb.NewECBDecrypter(block)
			ciphered := ECBSeal(mode, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBSeal(mode, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ECBOpen(decMode, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: blowfish.BlockSize,
				}
			}
		case BF_OFB:
			key, _ := generateBytes(blowfish.BlockSize)

			block, _ := blowfish.NewCipher(key)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, blowfish.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					OFBSealOpen(block, blowfish.BlockSize, textForCiphering)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: blowfish.BlockSize,
				}
			}
		case BF_CTR:
			key, _ := generateBytes(blowfish.BlockSize)

			block, _ := blowfish.NewCipher(key)
			ciphered := CTRSeal(block, blowfish.BlockSize, textForCiphering)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTRSeal(block, blowfish.BlockSize, textForCiphering)
				}
			}

			br.ciphering = testing.Benchmark(bench)

			bench = func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CTROpen(block, blowfish.BlockSize, ciphered)
				}
			}

			br.deciphering = testing.Benchmark(bench)

			if l >= conf.To {
				res = results.CipherExpResult{
					Algorithm: alg,
					Type:      "Алгоритм шифрования симметричный",
					KeyLength: blowfish.BlockSize,
				}
			}
		}

		b1, _ := strconv.ParseFloat(strings.Fields(br.ciphering.String())[1], 64)

		b2, _ := strconv.ParseFloat(strings.Fields(br.deciphering.String())[1], 64)

		b = strings.Fields(br.ciphering.String() + br.ciphering.MemString())

		if l == conf.To {
			cd := (float64(l) * mul2MBperS) / (b1)
			dd := (float64(l) * mul2MBperS) / (b2)

			res.CipheringDuration = math.Round(cd*100) / 100
			res.DecipheringDuration = math.Round(dd*100) / 100
		}

		x = append(x, l)
		y = append(y, int(b1))
	}

	res.Plot.X = x
	res.Plot.Y = y

	n, _ := strconv.ParseInt(b[0], 10, 32)
	a, _ := strconv.ParseInt(b[5], 10, 32)

	res.Hyst.AlocX = uint64(a)
	res.Hyst.OpX = int(n)
	res.Hyst.Alg = alg

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

func generateBytes(n int) ([]byte, time.Duration) {
	start := time.Now()

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return b, time.Since(start)
}
