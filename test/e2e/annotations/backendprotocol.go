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

package annotations

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/ingress-nginx/test/e2e/framework"
)

var _ = framework.IngressNginxDescribe("Annotations - Backendprotocol", func() {
	f := framework.NewDefaultFramework("backendprotocol")

	BeforeEach(func() {
		f.NewEchoDeploymentWithReplicas(2)
	})

	AfterEach(func() {
	})

	It("should set backend protocol to https:// and use proxy_pass", func() {
		host := "backendprotocol.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/backend-protocol": "HTTPS",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, "http-svc", 80, &annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring("proxy_pass https://upstream_balancer;"))
			})
	})

	It("should set backend protocol to grpc:// and use grpc_pass", func() {
		host := "backendprotocol.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/backend-protocol": "GRPC",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, "http-svc", 80, &annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring("grpc_pass grpc://upstream_balancer;"))
			})
	})

	It("should set backend protocol to grpcs:// and use grpc_pass", func() {
		host := "backendprotocol.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/backend-protocol": "GRPCS",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, "http-svc", 80, &annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring("grpc_pass grpcs://upstream_balancer;"))
			})
	})

	It("should set backend protocol to '' and use fastcgi_pass", func() {
		host := "backendprotocol.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/backend-protocol": "FCGI",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, "http-svc", 80, &annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring("fastcgi_pass upstream_balancer;"))
			})
	})

	It("should set backend protocol to '' and use ajp_pass", func() {
		host := "backendprotocol.foo.com"
		annotations := map[string]string{
			"nginx.ingress.kubernetes.io/backend-protocol": "AJP",
		}

		ing := framework.NewSingleIngress(host, "/", host, f.Namespace, "http-svc", 80, &annotations)
		f.EnsureIngress(ing)

		f.WaitForNginxServer(host,
			func(server string) bool {
				return Expect(server).Should(ContainSubstring("ajp_pass upstream_balancer;"))
			})
	})
})
