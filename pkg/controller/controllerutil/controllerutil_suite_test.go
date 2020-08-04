/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllerutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
)

func TestControllerutil(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteName := "Controllerutil Suite"
	RunSpecsWithDefaultAndCustomReporters(t, suiteName, []Reporter{printer.NewlineReporter{}, printer.NewProwReporter(suiteName)})
}

var t *envtest.Environment
var cfg *rest.Config
var c client.Client

var _ = BeforeSuite(func() {
	var err error

	t = &envtest.Environment{}

	cfg, err = t.Start()
	Expect(err).NotTo(HaveOccurred())

	c, err = client.New(cfg, client.Options{})
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(t.Stop()).To(Succeed())
})
