/*
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

package akeyless

import (
	"os"

	// nolint
	. "github.com/onsi/ginkgo"
	// nolint
	. "github.com/onsi/ginkgo/extensions/table"

	"github.com/external-secrets/external-secrets/e2e/framework"
	"github.com/external-secrets/external-secrets/e2e/suite/common"
)

var _ = Describe("[akeyless] ", func() {
	f := framework.New("eso-akeyless")
	accessID := os.Getenv("AKEYLESS_ACCESS_ID")
	accessType := os.Getenv("AKEYLESS_ACCESS_TYPE")
	accessTypeParam := os.Getenv("AKEYLESS_ACCESS_TYPE_PARAM")
	prov := newAkeylessProvider(f, accessID, accessType, accessTypeParam)

	DescribeTable("sync secrets", framework.TableFunc(f, prov),
		Entry(common.SimpleDataSync(f)),
		Entry(common.NestedJSONWithGJSON(f)),
		Entry(common.JSONDataFromSync(f)),
		Entry(common.JSONDataWithProperty(f)),
		Entry(common.JSONDataWithTemplate(f)),
		Entry(common.DockerJSONConfig(f)),
		Entry(common.DataPropertyDockerconfigJSON(f)),
		Entry(common.SSHKeySync(f)),
		Entry(common.SSHKeySyncDataProperty(f)),
		Entry(common.SyncWithoutTargetName(f)),
		Entry(common.JSONDataWithoutTargetName(f)),
	)
})
