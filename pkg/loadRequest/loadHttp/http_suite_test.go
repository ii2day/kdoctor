// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0
package loadHttp_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLoadHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "load request Suite")
}

var _ = BeforeSuite(func() {
	// nothing to do
})
