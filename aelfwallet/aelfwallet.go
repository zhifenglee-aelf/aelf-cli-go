package aelfwallet

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/foxnut/go-hdwallet"
)

type Wallet struct {
	PrivateKey string
	PublicKey  string
	Address    string
	Mnemonic   string
}

func CreateWalletFromMnemonic(mnemonic string) (Wallet, error) {
	masterKey, error := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
		hdwallet.Password(""),
		hdwallet.Language(hdwallet.English),
	)
	if error != nil {
		return Wallet{}, error
	}

	path := hdwallet.Path("m/44'/1616'/0'/0/0")
	key, _ := masterKey.GetChildKey(path)

	//ecPrivateKey, ecPublicKey := secp256k1.PrivKeyFromBytes(key.Private.Serialize())

	publicKey := key.PublicHex(false)
	publicKeyUncompressed, _ := hex.DecodeString(publicKey)
	//publicKeyUncompressed := ecPublicKey.SerializeUncompressed()
	//publicKey := hex.EncodeToString(publicKeyUncompressed)

	hash := sha256.Sum256(publicKeyUncompressed)
	hash = sha256.Sum256(hash[:])

	hashCheckSum := hashCheckEncode(hash[:])
	//log.Default().Println("hashCheckSum", (hex.EncodeToString(hashCheckSum)))

	address := base58.Encode(hashCheckSum)

	hashCheckSumLen := len(hashCheckSum)
	for i := 0; i < hashCheckSumLen && hashCheckSum[i] == 0; i++ {
		address = "1" + address
	}

	return Wallet{
		PrivateKey: key.PrivateHex(),
		//PrivateKey: hex.EncodeToString(ecPrivateKey.Serialize()),
		PublicKey: publicKey,
		Address:   address,
		Mnemonic:  mnemonic,
	}, nil
}

func CreateWallet() (Wallet, error) {
	mnemonic, error := hdwallet.NewMnemonic(12, hdwallet.English)
	if error != nil {
		return Wallet{}, error
	}

	return CreateWalletFromMnemonic(mnemonic)
}

func hashCheckEncode(input []byte) []byte {
	b := make([]byte, 0, len(input)+4)
	b = append(b, input[:]...)
	cksum := checksum(b)
	b = append(b, cksum[:]...)

	return b
}

func checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}

func padLeft(str string, charLen int, sign string) string {
	if sign == "" {
		sign = "0"
	}
	length := charLen - len(str)
	if length <= 0 {
		return str
	}
	return strings.Repeat(sign, length) + str
}

//put draft unhappy diary arctic sponsor alien awesome adjust bubble maid brave
//2ihA5K7sSsA78gekyhuh7gcnX4JkGVqJmSGnf8Kj1hZefR4sX5
//04c0f6abf0e3122f4a49646d67bacf85c80ad726ca781ccba572033a31162f22e55a4a106760cbf1306f26c25aea1e4bb71ee66cb3c5104245d6040cce64546cc7
