package generateAccountKeys

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/onflow/flow-go-sdk/crypto"
)

/*
	Flow uses ECDSA 256 to control access to user accounts.
	https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm
	You need to authorize at least one public key to control a new account
*/
/*
	This generates a private & public key to register an account via a phrase which eventually the user creates?
	Then this is used to authorize access to their account on the flow blockchain?
	Function takes a Hashing Algorithim Type and returns the hex version of private + public key
*/

func CreateKeys(algoType string) (string, string) {
	seed := make([]byte, crypto.MinSeedLength)
	_, err := rand.Read(seed)
	Except(err)

	hashingAlgo := crypto.StringToSignatureAlgorithm(algoType)
	privateKey, err := crypto.GeneratePrivateKey(hashingAlgo, seed)
	Except(err)

	publicKey := privateKey.PublicKey()
	pubHex := hex.EncodeToString(publicKey.Encode())
	priHex := hex.EncodeToString(privateKey.Encode())

	return pubHex, priHex

}

func Except(err error) {
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}

func ReadFile(path string) string {
	contents, err := ioutil.ReadFile(path)
	Except(err)

	return string(contents)
}
