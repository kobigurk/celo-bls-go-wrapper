// +build darwin,amd64,!ios

package bls

import (
  blsproxied "github.com/kobigurk/celo-bls-go-darwin/bls"
)

func HashDirect(message []byte, usePoP bool) ([]byte, error) {
  return blsproxied.HashDirect(message, usePoP)
}
