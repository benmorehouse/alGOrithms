package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAvltree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Avltree Suite")
}
