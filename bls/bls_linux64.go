// +build !android,linux,amd64,!musl

package bls

import (
  blsproxied "github.com/kobigurk/celo-bls-go-linux/bls"
)

func HashDirect(message []byte, usePoP bool) ([]byte, error) {
  return blsproxied.HashDirect(message, usePoP)
}
