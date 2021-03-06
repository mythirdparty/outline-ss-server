// Copyright 2018 Jigsaw Operations LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shadowsocks

import (
	"fmt"

	"github.com/shadowsocks/go-shadowsocks2/core"
	"github.com/shadowsocks/go-shadowsocks2/shadowaead"
)

func MakeTestCiphers(numCiphers int) (CipherList, error) {
	cipherList := NewCipherList()
	for i := 0; i < numCiphers; i++ {
		cipherID := fmt.Sprintf("id-%v", i)
		secret := fmt.Sprintf("secret-%v", i)
		cipher, err := core.PickCipher("chacha20-ietf-poly1305", nil, secret)
		if err != nil {
			return nil, fmt.Errorf("Failed to create cipher %v: %v", i, err)
		}
		cipherList.PushBack(cipherID, cipher.(shadowaead.Cipher))
	}
	return cipherList, nil
}

func MakeTestPayload(size int) []byte {
	payload := make([]byte, size)
	for i := 0; i < size; i++ {
		payload[i] = byte(i)
	}
	return payload
}
